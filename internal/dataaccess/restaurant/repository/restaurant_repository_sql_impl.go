package repository

import (
	"database/sql"

	"github.com/google/uuid"
	restaurant "github.com/leninner/order-service/internal/dataaccess/restaurant/entity"
)

type RestaurantRepositorySqlImpl struct {
	db *sql.DB
}

func NewRestaurantRepositorySqlImpl(db *sql.DB) *RestaurantRepositorySqlImpl {
	return &RestaurantRepositorySqlImpl{
		db: db,
	}
}

func (r *RestaurantRepositorySqlImpl) FindInformation(restaurantID uuid.UUID, productIDs []uuid.UUID) (*restaurant.RestaurantModel, error) {
	query := `
		SELECT restaurant_id, product_id, restaurant_name, restaurant_active, product_name, product_price, product_available 
		FROM order_restaurant_m_view 
		WHERE restaurant_id = $1 AND product_id = ANY($2)
	`

	rows, err := r.db.Query(query, restaurantID, productIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var restaurantModel restaurant.RestaurantModel
	err = rows.Scan(&restaurantModel.RestaurantID, &restaurantModel.ProductID, &restaurantModel.RestaurantName, &restaurantModel.RestaurantActive, &restaurantModel.ProductName, &restaurantModel.ProductPrice, &restaurantModel.ProductAvailable)
	if err != nil {
		return nil, err
	}

	return &restaurantModel, nil
}
