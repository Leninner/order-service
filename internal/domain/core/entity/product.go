package entity

import (
	"github.com/leninner/shared/domain/entity"
	sharedVO "github.com/leninner/shared/domain/valueobject"
)

type Product struct {
	entity.Entity[sharedVO.ProductID]

	Name  string
	Price sharedVO.Money
}

func (p *Product) UpdateWithConfirmedNameAndPrice(name string, price sharedVO.Money) {
	p.Name = name
	p.Price = price
}
