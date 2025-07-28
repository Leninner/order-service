package create

import (
	"github.com/google/uuid"
	sharedVO "github.com/leninner/shared/domain/valueobject"
)

type CreateOrderResponse struct {
	OrderTrackingID *uuid.UUID           `json:"orderTrackingId"`
	OrderStatus     sharedVO.OrderStatus `json:"orderStatus"`
	Message         *string              `json:"message,omitempty"`
}
