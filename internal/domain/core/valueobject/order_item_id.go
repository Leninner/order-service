package valueobject

import (
	"errors"

	sharedVO "github.com/leninner/shared/domain/valueobject"
)

type OrderItemID struct {
	sharedVO.WithID[int64]
}

func NewOrderItemID(id int64) (OrderItemID, error) {
	if id <= 0 {
		return OrderItemID{}, errors.New("order item ID must be positive")
	}

	return OrderItemID{WithID: sharedVO.WithID[int64]{ID: id}}, nil
}

func MustNewOrderItemID(id int64) OrderItemID {
	orderItemID, err := NewOrderItemID(id)
	if err != nil {
		panic(err)
	}
	return orderItemID
}
