package entity

import sharedVO "github.com/leninner/shared/domain/valueobject"

type ProductBuilder struct {
	product *Product
}

func NewProductBuilder() *ProductBuilder {
	return &ProductBuilder{
		product: &Product{},
	}
}

func (b *ProductBuilder) WithID(id *sharedVO.ProductID) *ProductBuilder {
	b.product.ID = *id
	return b
}

func (b *ProductBuilder) WithName(name string) *ProductBuilder {
	b.product.Name = name
	return b
}

func (b *ProductBuilder) WithPrice(price *sharedVO.Money) *ProductBuilder {
	b.product.Price = *price
	return b
}

func (b *ProductBuilder) Build() *Product {
	return b.product
}
