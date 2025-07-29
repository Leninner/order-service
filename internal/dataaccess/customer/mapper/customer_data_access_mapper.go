package mapper

import (
	"github.com/google/uuid"
	customer "github.com/leninner/order-service/internal/dataaccess/customer/entity"
	"github.com/leninner/order-service/internal/domain/core/entity"
	sharedVO "github.com/leninner/shared/domain/valueobject"
)

type CustomerDataAccessMapper struct {
}

func NewCustomerDataAccessMapper() *CustomerDataAccessMapper {
	return &CustomerDataAccessMapper{}
}

func (m *CustomerDataAccessMapper) CustomerModelToCustomerDomain(customerModel *customer.CustomerModel) *entity.Customer {
	if customerModel == nil {
		return nil
	}

	customerDomain := &entity.Customer{}
	customerDomain.SetID(&sharedVO.CustomerID{WithID: sharedVO.WithID[uuid.UUID]{ID: customerModel.ID}})

	return customerDomain
}

func (m *CustomerDataAccessMapper) CustomerDomainToCustomerModel(customerDomain *entity.Customer) *customer.CustomerModel {
	if customerDomain == nil {
		return nil
	}

	customerModel := &customer.CustomerModel{
		ID: customerDomain.GetID().GetValue(),
	}

	return customerModel
}