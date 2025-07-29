package create

import (
	"github.com/leninner/shared/utils/validator"
)

type OrderAddress struct {
	Street     *string `json:"street" validate:"required"`
	City       *string `json:"city" validate:"required"`
	State      *string `json:"state" validate:"required"`
	PostalCode *string `json:"postalCode" validate:"required"`
	Country    *string `json:"country" validate:"required"`
}

func ValidateOrderAddress(env *validator.ValidationEnvelope, addr *OrderAddress) {
	env.Check(addr.Street != nil, "street", "street is required")
	env.Check(addr.City != nil, "city", "city is required")
	env.Check(addr.State != nil, "state", "state is required")
	env.Check(addr.PostalCode != nil, "postalCode", "postalCode is required")
	env.Check(addr.Country != nil, "country", "country is required")
}
