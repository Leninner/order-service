package repository

import (
	"github.com/google/uuid"
	"github.com/leninner/order-service/internal/domain/core/entity"
)

type CustomerRepository interface {
	FindByID(customerID *uuid.UUID) (*entity.Customer, error)
}
