package entity

import (
	sharedEntity "github.com/leninner/shared/domain/entity"
	sharedVO "github.com/leninner/shared/domain/valueobject"
)

type Customer struct {
	sharedEntity.AggregateRoot[*sharedVO.CustomerID]
}
