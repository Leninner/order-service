package track

import (
	"github.com/google/uuid"
	sharedVO "github.com/leninner/shared/domain/valueobject"
	"github.com/leninner/shared/utils/validator"
)

type TrackOrderQuery struct {
	OrderTrackingID uuid.UUID            `json:"orderTrackingId" validate:"required"`
	OrderStatus     sharedVO.OrderStatus `json:"orderStatus" validate:"required"`
	FailureMessages []string             `json:"failureMessages"`
}

func ValidateTrackOrderQuery(v *validator.Validator, query *TrackOrderQuery) {
	v.Check(query.OrderTrackingID != uuid.Nil, "orderTrackingId", "orderTrackingId is required")
	v.Check(query.OrderStatus != "", "orderStatus", "orderStatus is required")
	v.Check(len(query.FailureMessages) == 0, "failureMessages", "failureMessages is required")
}
