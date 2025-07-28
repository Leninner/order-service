package event

import (
	"github.com/leninner/order-service/internal/domain/core/entity"
)

// OrderPaidEvent represents an event that occurs when an order is paid
type OrderPaidEvent struct {
	OrderEvent
}

// NewOrderPaidEvent creates a new OrderPaidEvent
func NewOrderPaidEvent(order *entity.Order) *OrderPaidEvent {
	return &OrderPaidEvent{
		OrderEvent: NewOrderEvent(order),
	}
}
