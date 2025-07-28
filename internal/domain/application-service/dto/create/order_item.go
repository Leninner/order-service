package create

import (
	"github.com/google/uuid"
	"github.com/leninner/shared/utils/validator"
)

type OrderItem struct {
	ProductID *uuid.UUID `json:"productId" validate:"required"`
	Quantity  *int       `json:"quantity" validate:"required,gt=0"`
	Price     *float64   `json:"price" validate:"required,gt=0"`
}

func ValidateOrderItem(v *validator.Validator, item *OrderItem) {
	v.Check(item.ProductID != nil, "productId", "productId is required")
	v.Check(item.Quantity != nil && *item.Quantity > 0, "quantity", "quantity must be greater than 0")
	v.Check(item.Price != nil && *item.Price > 0, "price", "price must be greater than 0")
}
