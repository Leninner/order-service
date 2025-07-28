package track

import (
	"github.com/google/uuid"
	sharedVO "github.com/leninner/shared/domain/valueobject"
)

type TrackOrderResponse struct {
	OrderTrackingID *uuid.UUID           `json:"orderTrackingId"`
	OrderStatus     sharedVO.OrderStatus `json:"orderStatus"`
	Message         *string              `json:"message,omitempty"`
	FailureMessages []string             `json:"failureMessages"`
}
