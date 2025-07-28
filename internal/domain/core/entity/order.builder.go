package entity

import (
	vo "github.com/leninner/order-service/internal/domain/core/valueobject"
	sharedVO "github.com/leninner/shared/domain/valueobject"
)

type OrderBuilder struct {
	order *Order
}

func NewOrderBuilder() *OrderBuilder {
	return &OrderBuilder{
		order: &Order{
			Items:           make([]OrderItem, 0),
			CustomerID:      sharedVO.NewCustomerID(),
			RestaurantID:    sharedVO.NewRestaurantID(),
			DeliveryAddress: vo.StreetAddress{},
			Price:           sharedVO.Money{},
			FailureMessages: make([]string, 0),
		},
	}
}

func (ob *OrderBuilder) WithCustomerID(customerID sharedVO.CustomerID) *OrderBuilder {
	ob.order.CustomerID = customerID
	return ob
}

func (ob *OrderBuilder) WithRestaurantID(restaurantID sharedVO.RestaurantID) *OrderBuilder {
	ob.order.RestaurantID = restaurantID
	return ob
}

func (ob *OrderBuilder) WithDeliveryAddress(address vo.StreetAddress) *OrderBuilder {
	ob.order.DeliveryAddress = address
	return ob
}

func (ob *OrderBuilder) WithPrice(price sharedVO.Money) *OrderBuilder {
	ob.order.Price = price
	return ob
}

func (ob *OrderBuilder) WithOrderStatus(status sharedVO.OrderStatus) *OrderBuilder {
	ob.order.OrderStatus = &status
	return ob
}

func (ob *OrderBuilder) WithTrackingID(trackingID vo.TrackingID) *OrderBuilder {
	ob.order.TrackingID = trackingID
	return ob
}

func (ob *OrderBuilder) WithFailureMessages(messages []string) *OrderBuilder {
	ob.order.FailureMessages = messages
	return ob
}

func (ob *OrderBuilder) WithItems(items []OrderItem) *OrderBuilder {
	ob.order.Items = items
	return ob
}

func (ob *OrderBuilder) Build() *Order {

	return ob.order
}
