package mapper

import (
	"github.com/google/uuid"
	order "github.com/leninner/order-service/internal/dataaccess/order/entity"
	"github.com/leninner/order-service/internal/domain/core/entity"
	"github.com/leninner/order-service/internal/domain/core/valueobject"
	sharedVO "github.com/leninner/shared/domain/valueobject"
)

type OrderDataAccessMapper struct {
}

func NewOrderDataAccessMapper() *OrderDataAccessMapper {
	return &OrderDataAccessMapper{}
}

func (m *OrderDataAccessMapper) OrderModelToOrderDomain(orderModel *order.OrderModel) *entity.Order {
	if orderModel == nil {
		return nil
	}

	orderDomain := &entity.Order{}
	orderDomain.SetID(&sharedVO.OrderID{WithID: sharedVO.WithID[uuid.UUID]{ID: orderModel.ID}})

	orderDomain.CustomerID = sharedVO.CustomerID{WithID: sharedVO.WithID[uuid.UUID]{ID: orderModel.CustomerID}}
	orderDomain.RestaurantID = sharedVO.RestaurantID{WithID: sharedVO.WithID[uuid.UUID]{ID: orderModel.RestaurantID}}
	orderDomain.TrackingID = valueobject.TrackingIDFromUUID(orderModel.TrackingID)
	orderDomain.Price = sharedVO.Money{Amount: orderModel.Price}

	orderStatus := sharedVO.OrderStatus(orderModel.OrderStatus)
	orderDomain.OrderStatus = &orderStatus

	orderDomain.FailureMessages = orderModel.FailureMessages

	if orderModel.Address != nil {
		orderDomain.DeliveryAddress = valueobject.StreetAddress{
			ID:         orderModel.Address.ID,
			Street:     orderModel.Address.Street,
			PostalCode: orderModel.Address.PostalCode,
			City:       orderModel.Address.City,
		}
	}

	orderDomain.Items = make([]entity.OrderItem, len(orderModel.Items))
	for i, itemModel := range orderModel.Items {
		orderDomain.Items[i] = entity.OrderItem{}
		orderDomain.Items[i].SetID(valueobject.MustNewOrderItemID(int64(i + 1)))
		orderDomain.Items[i].OrderID = sharedVO.OrderID{WithID: sharedVO.WithID[uuid.UUID]{ID: orderModel.ID}}
		orderDomain.Items[i].Product = entity.Product{}
		orderDomain.Items[i].Product.SetID(sharedVO.ProductID{WithID: sharedVO.WithID[uuid.UUID]{ID: itemModel.ProductID}})
		orderDomain.Items[i].Quantity = itemModel.Quantity
		orderDomain.Items[i].Price = sharedVO.Money{Amount: itemModel.Price}
		orderDomain.Items[i].SubTotal = sharedVO.Money{Amount: itemModel.SubTotal}
	}

	return orderDomain
}

func (m *OrderDataAccessMapper) OrderDomainToOrderModel(orderDomain *entity.Order) *order.OrderModel {
	if orderDomain == nil {
		return nil
	}

	orderModel := &order.OrderModel{
		ID:              orderDomain.GetID().GetValue(),
		CustomerID:      orderDomain.CustomerID.GetValue(),
		RestaurantID:    orderDomain.RestaurantID.GetValue(),
		TrackingID:      orderDomain.TrackingID.GetValue(),
		Price:           orderDomain.Price.Amount,
		OrderStatus:     string(*orderDomain.OrderStatus),
		FailureMessages: orderDomain.FailureMessages,
	}

	if orderDomain.DeliveryAddress.ID != uuid.Nil {
		orderModel.Address = &order.AddressModel{
			ID:         orderDomain.DeliveryAddress.ID,
			OrderID:    orderDomain.GetID().GetValue(),
			Street:     orderDomain.DeliveryAddress.Street,
			PostalCode: orderDomain.DeliveryAddress.PostalCode,
			City:       orderDomain.DeliveryAddress.City,
		}
	}

	orderModel.Items = make([]order.OrderItemModel, len(orderDomain.Items))
	for i, item := range orderDomain.Items {
		itemID := item.GetID()
		productID := item.Product.GetID()
		
		orderModel.Items[i] = order.OrderItemModel{
			ID:        itemID.GetValue(),
			OrderID:   orderDomain.GetID().GetValue(),
			ProductID: productID.GetValue(),
			Quantity:  item.Quantity,
			Price:     item.Price.Amount,
			SubTotal:  item.SubTotal.Amount,
		}
	}

	return orderModel
}
