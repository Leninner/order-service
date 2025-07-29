package repository

import (
	"github.com/leninner/order-service/internal/domain/core/entity"
)

type RestaurantRepository interface {
	FindInformation(restaurant entity.Restaurant) (*entity.Restaurant, error)
}
