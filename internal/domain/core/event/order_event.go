package event

import (
	"time"

	"github.com/leninner/order-service/internal/domain/core/entity"
)

// OrderEvent represents an event related to an order domain
type OrderEvent struct {
	Order     *entity.Order
	CreatedAt time.Time
}

func (e *OrderEvent) IsDomainEvent() {}

// NewOrderEvent creates a new OrderEvent
func NewOrderEvent(order *entity.Order) OrderEvent {
	return OrderEvent{
		Order:     order,
		CreatedAt: time.Now().UTC(),
	}
}

// GetOrder returns the order associated with this event
func (e *OrderEvent) GetOrder() *entity.Order {
	return e.Order
}

// SetOrder sets the order associated with this event
func (e *OrderEvent) SetOrder(order *entity.Order) {
	e.Order = order
}
