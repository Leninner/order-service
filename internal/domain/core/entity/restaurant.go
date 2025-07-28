package entity

import (
	sharedEntity "github.com/leninner/shared/domain/entity"
	sharedVO "github.com/leninner/shared/domain/valueobject"
)

type Restaurant struct {
	sharedEntity.AggregateRoot[*sharedVO.RestaurantID]

	Products []Product
	Active   bool
}
