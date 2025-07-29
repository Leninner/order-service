package rest

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/leninner/order-service/internal/domain/application-service/dto/create"
	"github.com/leninner/order-service/internal/domain/application-service/ports/input/service"
	"github.com/leninner/shared/exception"
	utils "github.com/leninner/shared/utils"
	"github.com/leninner/shared/utils/validator"
)

type OrderController struct {
	orderService service.OrderApplicationService
}

func NewOrderController(orderService service.OrderApplicationService) *OrderController {
	return &OrderController{orderService: orderService}
}

func (c *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var input struct {
		CustomerID   *uuid.UUID          `json:"customerId" validate:"required"`
		RestaurantID *uuid.UUID          `json:"restaurantId" validate:"required"`
		Price        *float64            `json:"price" validate:"required,gt=0"`
		Items        []create.OrderItem  `json:"items" validate:"required,min=1"`
		Address      create.OrderAddress `json:"address" validate:"required"`
	}

	err := utils.ReadJSON(w, r, &input)
	if err != nil {
		exception.BadRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	items := []create.OrderItem{}

	for _, item := range input.Items {
		itemToCreate := create.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
		}

		items = append(items, itemToCreate)
	}

	address := create.OrderAddress{
		City:       input.Address.City,
		State:      input.Address.State,
		PostalCode: input.Address.PostalCode,
		Country:    input.Address.Country,
		Street:     input.Address.Street,
	}

	command := &create.CreateOrderCommand{
		CustomerID:   input.CustomerID,
		RestaurantID: input.RestaurantID,
		Price:        input.Price,
		Items:        items,
		Address:      address,
	}

	if create.ValidateOrderCommand(v, command); !v.Valid() {
		exception.FailedValidationResponse(w, r, v.Errors)
		return
	}

	createOrderResponse, err := c.orderService.CreateOrder(*command)
	if err != nil {
		exception.ErrorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"order": createOrderResponse}, nil)
}
