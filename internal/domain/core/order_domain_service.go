package core

import (
	"github.com/leninner/order-service/internal/domain/core/entity"
	"github.com/leninner/order-service/internal/domain/core/event"
)

// OrderDomainService defines the domain service for the order domain
type OrderDomainService interface {
	ValidateAndInitiateOrder(order *entity.Order, restaurant *entity.Restaurant) (*event.OrderCreatedEvent, error)
	PayOrder(order *entity.Order) (*event.OrderPaidEvent, error)
	ApproveOrder(order *entity.Order) error
	CancelOrderPayment(order *entity.Order, failureMessages []string) (*event.OrderCancelledEvent, error)
	CancelOrder(order *entity.Order, failureMessages []string) error
}
