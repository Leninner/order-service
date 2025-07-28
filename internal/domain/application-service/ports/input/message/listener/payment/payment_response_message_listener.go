package payment

import "github.com/leninner/order-service/internal/domain/application-service/dto/message"

type PaymentResponseMessageListener interface {
	PaymentCompleted(paymentResponse *message.PaymentResponse) error

	PaymentCancelled(paymentResponse *message.PaymentResponse) error
}
