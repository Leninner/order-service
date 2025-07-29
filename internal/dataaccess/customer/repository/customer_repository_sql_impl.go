package repository

import (
	"database/sql"

	"github.com/google/uuid"
	customer "github.com/leninner/order-service/internal/dataaccess/customer/entity"
)

type CustomerRepositorySQLImpl struct {
	db *sql.DB
}

func NewCustomerRepositorySQLImpl(db *sql.DB) *CustomerRepositorySQLImpl {
	return &CustomerRepositorySQLImpl{
		db: db,
	}
}

func (r *CustomerRepositorySQLImpl) FindByID(customerID *uuid.UUID) (*customer.CustomerModel, error) {
	query := `
		SELECT id FROM order_customer_m_view WHERE id = $1
	`

	row := r.db.QueryRow(query, customerID)

	var customer customer.CustomerModel
	err := row.Scan(&customer.ID)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}
