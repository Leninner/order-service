package message

import (
	"time"

	sharedVO "github.com/leninner/shared/domain/valueobject"
)

type PaymentResponse struct {
	ID              string                 `json:"id"`
	SagaID          string                 `json:"sagaId"`
	OrderID         string                 `json:"orderId"`
	PaymentID       string                 `json:"paymentId"`
	CustomerID      string                 `json:"customerId"`
	Price           float64                `json:"price"`
	CreatedAt       time.Time              `json:"createdAt"`
	PaymentStatus   sharedVO.PaymentStatus `json:"paymentStatus"`
	FailureMessages []string               `json:"failureMessages"`
}
