package message

import (
	"time"

	sharedVO "github.com/leninner/shared/domain/valueobject"
)

type RestaurantApprovalResponse struct {
	ID                  string                       `json:"id"`
	SagaID              string                       `json:"sagaId"`
	OrderID             string                       `json:"orderId"`
	RestaurantID        string                       `json:"restaurantId"`
	CreatedAt           time.Time                    `json:"createdAt"`
	OrderApprovalStatus sharedVO.OrderApprovalStatus `json:"orderApprovalStatus"`
	FailureMessages     []string                     `json:"failureMessages"`
}
