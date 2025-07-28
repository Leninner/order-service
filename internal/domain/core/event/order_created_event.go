package event

import (
	"github.com/leninner/order-service/internal/domain/core/entity"
)

// OrderCreatedEvent represents an event that occurs when an order is created
type OrderCreatedEvent struct {
	OrderEvent
}

// NewOrderCreatedEvent creates a new OrderCreatedEvent
func NewOrderCreatedEvent(order *entity.Order) *OrderCreatedEvent {
	return &OrderCreatedEvent{
		OrderEvent: NewOrderEvent(order),
	}
}
