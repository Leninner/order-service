package repository

import (
	"github.com/leninner/order-service/internal/domain/core/entity"
	sharedVO "github.com/leninner/shared/domain/valueobject"
)

type RestaurantRepository interface {
	FindInformation(restaurantID sharedVO.RestaurantID) (*entity.Restaurant, error)
}
