package applicationservice

import (
	"github.com/leninner/order-service/internal/domain/application-service/dto/message"
	service "github.com/leninner/order-service/internal/domain/application-service/ports/input/service"
)

type PaymentResponseMessageListenerImpl struct {
	orderApplicationService service.OrderApplicationService
}

func NewPaymentResponseMessageListenerImpl(orderApplicationService service.OrderApplicationService) *PaymentResponseMessageListenerImpl {
	return &PaymentResponseMessageListenerImpl{orderApplicationService: orderApplicationService}
}

func (p *PaymentResponseMessageListenerImpl) PaymentCompleted(paymentResponse *message.PaymentResponse) error {
	return nil
}

func (p *PaymentResponseMessageListenerImpl) PaymentCancelled(paymentResponse *message.PaymentResponse) error {
	return nil
}
