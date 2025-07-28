package repository

import (
	"github.com/leninner/order-service/internal/domain/core/entity"
	"github.com/leninner/order-service/internal/domain/core/valueobject"
)

type OrderRepository interface {
	Save(order *entity.Order) (*entity.Order, error)

	FindByTrackingID(orderTrackingID valueobject.TrackingID) (*entity.Order, error)
}
