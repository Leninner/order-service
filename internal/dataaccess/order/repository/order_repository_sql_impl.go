package repository

import (
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"
	order "github.com/leninner/order-service/internal/dataaccess/order/entity"
	"github.com/leninner/order-service/internal/domain/core/exception"
	"github.com/leninner/order-service/internal/domain/core/valueobject"
)

type OrderRepositorySQLImpl struct {
	db *sql.DB
}

func NewOrderRepositorySQLImpl(db *sql.DB) *OrderRepositorySQLImpl {
	return &OrderRepositorySQLImpl{
		db: db,
	}
}

func (r *OrderRepositorySQLImpl) Save(orderModel *order.OrderModel) (*order.OrderModel, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := `
		INSERT INTO orders (id, customer_id, restaurant_id, tracking_id, price, order_status, failure_messages)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (id) DO UPDATE SET
			customer_id = EXCLUDED.customer_id,
			restaurant_id = EXCLUDED.restaurant_id,
			tracking_id = EXCLUDED.tracking_id,
			price = EXCLUDED.price,
			order_status = EXCLUDED.order_status,
			failure_messages = EXCLUDED.failure_messages,
			updated_at = NOW()
	`

	failureMessagesJSON, _ := json.Marshal(orderModel.FailureMessages)
	
	_, err = tx.Exec(query, 
		orderModel.ID, 
		orderModel.CustomerID, 
		orderModel.RestaurantID, 
		orderModel.TrackingID, 
		orderModel.Price, 
		orderModel.OrderStatus, 
		failureMessagesJSON,
	)
	if err != nil {
		return nil, err
	}

	if orderModel.Address != nil {
		addressQuery := `
			INSERT INTO order_addresses (id, order_id, street, postal_code, city)
			VALUES ($1, $2, $3, $4, $5)
			ON CONFLICT (order_id) DO UPDATE SET
				street = EXCLUDED.street,
				postal_code = EXCLUDED.postal_code,
				city = EXCLUDED.city
		`
		_, err = tx.Exec(addressQuery,
			orderModel.Address.ID,
			orderModel.Address.OrderID,
			orderModel.Address.Street,
			orderModel.Address.PostalCode,
			orderModel.Address.City,
		)
		if err != nil {
			return nil, err
		}
	}

	if len(orderModel.Items) > 0 {
		deleteItemsQuery := `DELETE FROM order_items WHERE order_id = $1`
		_, err = tx.Exec(deleteItemsQuery, orderModel.ID)
		if err != nil {
			return nil, err
		}

		itemQuery := `
			INSERT INTO order_items (id, order_id, product_id, quantity, price, sub_total)
			VALUES ($1, $2, $3, $4, $5, $6)
		`
		for _, item := range orderModel.Items {
			_, err = tx.Exec(itemQuery,
				item.ID,
				item.OrderID,
				item.ProductID,
				item.Quantity,
				item.Price,
				item.SubTotal,
			)
			if err != nil {
				return nil, err
			}
		}
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return orderModel, nil
}

func (r *OrderRepositorySQLImpl) FindByTrackingID(orderTrackingID valueobject.TrackingID) (*order.OrderModel, error) {
	query := `
		SELECT o.id, o.customer_id, o.restaurant_id, o.tracking_id, o.price, o.order_status, o.failure_messages,
		       a.id, a.street, a.postal_code, a.city
		FROM orders o
		LEFT JOIN order_addresses a ON o.id = a.order_id
		WHERE o.tracking_id = $1
	`

	row := r.db.QueryRow(query, orderTrackingID.GetValue())
	
	var orderModel order.OrderModel
	var failureMessagesJSON []byte
	var street, postalCode, city sql.NullString
	var addressID uuid.UUID
	
	err := row.Scan(
		&orderModel.ID,
		&orderModel.CustomerID,
		&orderModel.RestaurantID,
		&orderModel.TrackingID,
		&orderModel.Price,
		&orderModel.OrderStatus,
		&failureMessagesJSON,
		&addressID,
		&street,
		&postalCode,
		&city,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, exception.NewOrderDomainException("order with tracking id " + orderTrackingID.GetValue().String() + " not found")
		}
		return nil, err
	}

	json.Unmarshal(failureMessagesJSON, &orderModel.FailureMessages)

	if addressID != uuid.Nil {
		orderModel.Address = &order.AddressModel{
			ID:         addressID,
			OrderID:    orderModel.ID,
			Street:     street.String,
			PostalCode: postalCode.String,
			City:       city.String,
		}
	}

	itemsQuery := `
		SELECT id, product_id, quantity, price, sub_total
		FROM order_items
		WHERE order_id = $1
	`
	
	rows, err := r.db.Query(itemsQuery, orderModel.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item order.OrderItemModel
		err := rows.Scan(&item.ID, &item.ProductID, &item.Quantity, &item.Price, &item.SubTotal)
		if err != nil {
			return nil, err
		}
		item.OrderID = orderModel.ID
		orderModel.Items = append(orderModel.Items, item)
	}

	return &orderModel, nil
}
