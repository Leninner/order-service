package create

import (
	"github.com/google/uuid"
	"github.com/leninner/shared/utils/validator"
)

type CreateOrderCommand struct {
	CustomerID   *uuid.UUID   `json:"customerId" validate:"required"`
	RestaurantID *uuid.UUID   `json:"restaurantId" validate:"required"`
	Price        *float64     `json:"price" validate:"required,gt=0"`
	Items        []OrderItem  `json:"items" validate:"required,min=1"`
	Address      OrderAddress `json:"address" validate:"required"`
}

type CreateOrderCommandOption func(*CreateOrderCommand)

func NewCreateOrderCommand(opts ...CreateOrderCommandOption) *CreateOrderCommand {
	cmd := &CreateOrderCommand{}
	for _, opt := range opts {
		opt(cmd)
	}
	return cmd
}

func WithCustomerID(customerID uuid.UUID) CreateOrderCommandOption {
	return func(cmd *CreateOrderCommand) {
		cmd.CustomerID = &customerID
	}
}

func WithRestaurantID(restaurantID uuid.UUID) CreateOrderCommandOption {
	return func(cmd *CreateOrderCommand) {
		cmd.RestaurantID = &restaurantID
	}
}

func WithPrice(price float64) CreateOrderCommandOption {
	return func(cmd *CreateOrderCommand) {
		cmd.Price = &price
	}
}

func WithItems(items []OrderItem) CreateOrderCommandOption {
	return func(cmd *CreateOrderCommand) {
		cmd.Items = items
	}
}

func WithAddress(address OrderAddress) CreateOrderCommandOption {
	return func(cmd *CreateOrderCommand) {
		cmd.Address = address
	}
}

func ValidateOrderCommand(v *validator.Validator, command *CreateOrderCommand) {
	v.Check(command.CustomerID != nil, "customerId", "customerId is required")
	v.Check(command.RestaurantID != nil, "restaurantId", "restaurantId is required")
	v.Check(command.Price != nil && *command.Price > 0, "price", "price must be greater than 0")
	v.Check(len(command.Items) > 0, "items", "items list cannot be empty")

	for _, item := range command.Items {
		ValidateOrderItem(v, &item)
	}

	ValidateOrderAddress(v, &command.Address)
}
