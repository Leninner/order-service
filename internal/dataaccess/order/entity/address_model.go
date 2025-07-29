package order

import "github.com/google/uuid"

type AddressModel struct {
	ID         uuid.UUID `json:"id"`
	OrderID    uuid.UUID `json:"order_id"`
	Street     string    `json:"street"`
	PostalCode string    `json:"postal_code"`
	City       string    `json:"city"`
} 