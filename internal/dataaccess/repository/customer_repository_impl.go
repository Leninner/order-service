package repository

import (
	"github.com/google/uuid"
	"github.com/leninner/order-service/internal/domain/core/entity"
	"github.com/leninner/order-service/internal/domain/core/exception"
)

type CustomerRepositoryImpl struct {
	customers map[string]*entity.Customer
}

func NewCustomerRepositoryImpl() *CustomerRepositoryImpl {
	return &CustomerRepositoryImpl{
		customers: make(map[string]*entity.Customer),
	}
}

func (r *CustomerRepositoryImpl) FindByID(customerID *uuid.UUID) (*entity.Customer, error) {
	if customer, exists := r.customers[customerID.String()]; exists {
		return customer, nil
	}

	return nil, exception.NewOrderDomainException("customer with id " + customerID.String() + " not found")
}
