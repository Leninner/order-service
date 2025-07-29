package publisher

import (
	"github.com/leninner/order-service/internal/domain/core/event"
)

type OrderCreatedPaymentRequestMessagePublisherImpl struct{}

func NewOrderCreatedPaymentRequestMessagePublisherImpl() *OrderCreatedPaymentRequestMessagePublisherImpl {
	return &OrderCreatedPaymentRequestMessagePublisherImpl{}
}

func (p *OrderCreatedPaymentRequestMessagePublisherImpl) Publish(event *event.OrderCreatedEvent) error {
	return nil
}
