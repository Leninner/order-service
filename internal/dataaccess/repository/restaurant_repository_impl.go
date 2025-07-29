package repository

import (
	"github.com/leninner/order-service/internal/domain/core/entity"
	"github.com/leninner/order-service/internal/domain/core/exception"
	sharedVO "github.com/leninner/shared/domain/valueobject"
)

type RestaurantRepositoryImpl struct {
	restaurants map[string]*entity.Restaurant
}

func NewRestaurantRepositoryImpl() *RestaurantRepositoryImpl {
	return &RestaurantRepositoryImpl{
		restaurants: make(map[string]*entity.Restaurant),
	}
}

func (r *RestaurantRepositoryImpl) FindInformation(restaurantID sharedVO.RestaurantID) (*entity.Restaurant, error) {
	if restaurant, exists := r.restaurants[restaurantID.GetValue().String()]; exists {
		return restaurant, nil
	}
	return nil, exception.NewOrderDomainException("restaurant with id " + restaurantID.GetValue().String() + " not found")
}
