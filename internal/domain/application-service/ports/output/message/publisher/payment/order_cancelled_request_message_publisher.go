package payment

import (
	"github.com/leninner/order-service/internal/domain/core/event"
	"github.com/leninner/shared/domain/event/publisher"
)

type OrderCancelledPaymentRequestMessagePublisher interface {
	publisher.DomainEventPublisher[*event.OrderCancelledEvent]
}
