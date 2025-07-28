package entity

import sharedVO "github.com/leninner/shared/domain/valueobject"

type CustomerBuilder struct {
	customer *Customer
}

func NewCustomerBuilder() *CustomerBuilder {
	return &CustomerBuilder{customer: &Customer{}}
}

func (b *CustomerBuilder) WithID(id *sharedVO.CustomerID) *CustomerBuilder {
	b.customer.ID = id
	return b
}

func (b *CustomerBuilder) Build() *Customer {
	return b.customer
}
