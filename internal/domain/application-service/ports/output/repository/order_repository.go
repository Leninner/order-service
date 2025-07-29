package repository

import (
	"github.com/google/uuid"
	"github.com/leninner/order-service/internal/domain/core/entity"
)

type OrderRepository interface {
	Save(order *entity.Order) (*entity.Order, error)

	FindByTrackingID(orderTrackingID uuid.UUID) (*entity.Order, error)
}
