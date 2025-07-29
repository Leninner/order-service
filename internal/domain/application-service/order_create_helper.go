package applicationservice

import (
	"github.com/google/uuid"
	"github.com/leninner/order-service/internal/domain/application-service/dto/create"
	"github.com/leninner/order-service/internal/domain/application-service/mapper"
	repository "github.com/leninner/order-service/internal/domain/application-service/ports/output/repository"
	domain "github.com/leninner/order-service/internal/domain/core"
	"github.com/leninner/order-service/internal/domain/core/entity"
	"github.com/leninner/order-service/internal/domain/core/event"
	"github.com/leninner/order-service/internal/domain/core/exception"
	sharedVO "github.com/leninner/shared/domain/valueobject"
	"github.com/leninner/shared/logger"
	"go.uber.org/zap"
)

type OrderCreateHelper struct {
	orderRepository      repository.OrderRepository
	customerRepository   repository.CustomerRepository
	restaurantRepository repository.RestaurantRepository

	orderDataMapper    mapper.OrderDataMapper
	orderDomainService domain.OrderDomainService

	logger *logger.Logger
}

func NewOrderCreateHelper(
	orderRepository repository.OrderRepository,
	customerRepository repository.CustomerRepository,
	restaurantRepository repository.RestaurantRepository,
	orderDataMapper mapper.OrderDataMapper,
	orderDomainService domain.OrderDomainService,
	logger *logger.Logger,
) *OrderCreateHelper {
	return &OrderCreateHelper{
		orderRepository:      orderRepository,
		customerRepository:   customerRepository,
		restaurantRepository: restaurantRepository,
		orderDataMapper:      orderDataMapper,
		orderDomainService:   orderDomainService,
		logger:               logger,
	}
}

func (h *OrderCreateHelper) PersistOrder(command create.CreateOrderCommand) (*event.OrderCreatedEvent, error) {
	err := h.checkCustomer(command.CustomerID)
	if err != nil {
		return nil, err
	}

	restaurant, err := h.checkRestaurant(&command)
	if err != nil {
		return nil, err
	}

	order, err := h.orderDataMapper.CreateOrderCommandToOrder(&command)
	if err != nil {
		return nil, err
	}

	orderCreatedEvent, err := h.orderDomainService.ValidateAndInitiateOrder(order, restaurant)
	if err != nil {
		return nil, err
	}

	order, err = h.saveOrder(orderCreatedEvent.GetOrder())
	if err != nil {
		return nil, err
	}

	h.logger.Info(
		"Order persisted successfully",
		zap.String("orderId", order.GetID().GetValue().String()),
	)

	return orderCreatedEvent, nil
}

func (h *OrderCreateHelper) checkCustomer(customerID *uuid.UUID) error {
	_, err := h.customerRepository.FindByID(customerID)
	if err != nil {
		return exception.NewOrderDomainException("customer with id " + customerID.String() + " not found")
	}

	return nil
}

func (h *OrderCreateHelper) checkRestaurant(command *create.CreateOrderCommand) (*entity.Restaurant, error) {
	restaurantID := sharedVO.NewRestaurantIDFromUUID(command.RestaurantID)

	restaurantInformation, err := h.restaurantRepository.FindInformation(restaurantID)
	if err != nil {
		return nil, exception.NewOrderDomainException("restaurant with id " + command.RestaurantID.String() + " not found")
	}

	return restaurantInformation, nil
}

func (h *OrderCreateHelper) saveOrder(order *entity.Order) (*entity.Order, error) {
	order, err := h.orderRepository.Save(order)
	if err != nil {
		return nil, exception.NewOrderDomainException("could not save order")
	}

	return order, nil
}
