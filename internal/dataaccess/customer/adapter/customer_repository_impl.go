package adapter

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/leninner/order-service/internal/dataaccess/customer/mapper"
	"github.com/leninner/order-service/internal/dataaccess/customer/repository"
	"github.com/leninner/order-service/internal/domain/core/entity"
)

type CustomerRepositoryImpl struct {
	customerSqlRepository *repository.CustomerRepositorySQLImpl
	customerDataAccessMapper *mapper.CustomerDataAccessMapper
}

func NewCustomerRepositoryImpl(db *sql.DB) *CustomerRepositoryImpl {
	return &CustomerRepositoryImpl{
		customerSqlRepository: repository.NewCustomerRepositorySQLImpl(db),
		customerDataAccessMapper: mapper.NewCustomerDataAccessMapper(),
	}
}

func (r *CustomerRepositoryImpl) FindByID(customerID *uuid.UUID) (*entity.Customer, error) {
	customerModel, err := r.customerSqlRepository.FindByID(customerID)
	if err != nil {
		return nil, err
	}

	customerDomain := r.customerDataAccessMapper.CustomerModelToCustomerDomain(customerModel)

	return customerDomain, nil
}