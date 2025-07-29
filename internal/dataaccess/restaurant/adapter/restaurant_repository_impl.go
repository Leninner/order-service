package adapter

import (
	"database/sql"

	"github.com/leninner/order-service/internal/dataaccess/restaurant/mapper"
	restaurant "github.com/leninner/order-service/internal/dataaccess/restaurant/repository"
	"github.com/leninner/order-service/internal/domain/core/entity"
)

type RestaurantRepositoryImpl struct {
	restaurantRepository *restaurant.RestaurantRepositorySqlImpl
	restaurantDataAccessMapper *mapper.RestaurantDataAccessMapper
}

func NewRestaurantRepositoryImpl(db *sql.DB) *RestaurantRepositoryImpl {
	return &RestaurantRepositoryImpl{
		restaurantRepository: restaurant.NewRestaurantRepositorySqlImpl(db),
		restaurantDataAccessMapper: mapper.NewRestaurantDataAccessMapper(),
	}
}

func (r *RestaurantRepositoryImpl) FindInformation(restaurant entity.Restaurant) (*entity.Restaurant, error) {
	restaurantProducts := r.restaurantDataAccessMapper.RestaurantToRestaurantProducts(&restaurant)

	restaurantModel, err := r.restaurantRepository.FindInformation(restaurant.GetID().GetValue(), restaurantProducts)
	if err != nil {
		return nil, err
	}

	if restaurantModel == nil {
		return nil, nil
	}

	restaurantDomain := r.restaurantDataAccessMapper.RestaurantModelToRestaurantDomain(restaurantModel)

	return restaurantDomain, nil
}