package mapper

import (
	"github.com/google/uuid"
	restaurant "github.com/leninner/order-service/internal/dataaccess/restaurant/entity"
	"github.com/leninner/order-service/internal/domain/core/entity"
	sharedVO "github.com/leninner/shared/domain/valueobject"
)

type RestaurantDataAccessMapper struct {
}

func NewRestaurantDataAccessMapper() *RestaurantDataAccessMapper {
	return &RestaurantDataAccessMapper{}
}

func (m *RestaurantDataAccessMapper) RestaurantToRestaurantProducts(restaurant *entity.Restaurant) []uuid.UUID {
	if restaurant == nil || len(restaurant.Products) == 0 {
		return []uuid.UUID{}
	}

	productIDs := make([]uuid.UUID, len(restaurant.Products))
	for i, product := range restaurant.Products {
		productID := product.GetID()
		productIDs[i] = productID.GetValue()
	}
	return productIDs
}

func (m *RestaurantDataAccessMapper) RestaurantModelToRestaurantDomain(restaurantModel *restaurant.RestaurantModel) *entity.Restaurant {
	if restaurantModel == nil {
		return nil
	}

	restaurantDomain := &entity.Restaurant{}
	restaurantDomain.SetID(&sharedVO.RestaurantID{WithID: sharedVO.WithID[uuid.UUID]{ID: restaurantModel.RestaurantID}})

	productBuilder := entity.NewProductBuilder()
	productBuilder.WithID(&sharedVO.ProductID{WithID: sharedVO.WithID[uuid.UUID]{ID: restaurantModel.ProductID}})
	productBuilder.WithName(restaurantModel.ProductName)
	
	productBuilder.WithPrice(sharedVO.NewMoney(restaurantModel.ProductPrice))
	product := productBuilder.Build()
	restaurantDomain.Products = append(restaurantDomain.Products, *product)

	return restaurantDomain
}

func (m *RestaurantDataAccessMapper) RestaurantDomainToRestaurantModel(restaurantDomain *entity.Restaurant) *restaurant.RestaurantModel {
	if restaurantDomain == nil {
		return nil
	}

	restaurantModel := &restaurant.RestaurantModel{
		RestaurantID: restaurantDomain.GetID().GetValue(),
	}

	return restaurantModel
}