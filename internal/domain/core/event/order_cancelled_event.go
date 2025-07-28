package event

import (
	"github.com/leninner/order-service/internal/domain/core/entity"
)

// OrderCancelledEvent represents an event that occurs when an order is cancelled
type OrderCancelledEvent struct {
	OrderEvent
}

// NewOrderCancelledEvent creates a new OrderCancelledEvent
func NewOrderCancelledEvent(order *entity.Order) *OrderCancelledEvent {
	return &OrderCancelledEvent{
		OrderEvent: NewOrderEvent(order),
	}
}
