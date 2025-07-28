package tests

import (
	"github.com/google/uuid"
	"github.com/leninner/order-service/internal/domain/application-service/ports/output/message/publisher/payment"
	"github.com/leninner/order-service/internal/domain/application-service/ports/output/message/publisher/restaurantapproval"
	"github.com/leninner/order-service/internal/domain/application-service/ports/output/repository"
	"github.com/leninner/order-service/internal/domain/core"
	"github.com/leninner/order-service/internal/domain/core/entity"
	"github.com/leninner/order-service/internal/domain/core/event"
	"github.com/leninner/order-service/internal/domain/core/valueobject"
	sharedVO "github.com/leninner/shared/domain/valueobject"
	"github.com/leninner/shared/logger"
)

type OrderTestConfiguration struct {
	PaymentRequestMessagePublisher           payment.OrderCreatedPaymentRequestMessagePublisher
	RestaurantApprovalRequestMessagePublisher restaurantapproval.OrderPaidRestaurantApprovalRequestMessagePublisher
	OrderRepository                          repository.OrderRepository
	CustomerRepository                       repository.CustomerRepository
	RestaurantRepository                     repository.RestaurantRepository
	OrderDomainService                       core.OrderDomainService
}

func NewOrderTestConfiguration(logger *logger.Logger) *OrderTestConfiguration {
	return &OrderTestConfiguration{
		PaymentRequestMessagePublisher:           &MockPaymentRequestMessagePublisher{},
		RestaurantApprovalRequestMessagePublisher: &MockRestaurantApprovalRequestMessagePublisher{},
		OrderRepository:                          &MockOrderRepository{},
		CustomerRepository:                       &MockCustomerRepository{},
		RestaurantRepository:                     NewMockRestaurantRepository(),
		OrderDomainService:                       core.NewOrderDomainServiceImpl(logger),
	}
}

type MockPaymentRequestMessagePublisher struct{}

func (m *MockPaymentRequestMessagePublisher) Publish(event *event.OrderCreatedEvent) error {
	return nil
}

type MockRestaurantApprovalRequestMessagePublisher struct{}

func (m *MockRestaurantApprovalRequestMessagePublisher) Publish(event *event.OrderPaidEvent) error {
	return nil
}

type MockOrderRepository struct{}

func (m *MockOrderRepository) Save(order *entity.Order) (*entity.Order, error) {
	return order, nil
}

func (m *MockOrderRepository) FindByTrackingID(orderTrackingID valueobject.TrackingID) (*entity.Order, error) {
	return nil, nil
}

type MockCustomerRepository struct{}

func (m *MockCustomerRepository) FindByID(customerID *uuid.UUID) (*entity.Customer, error) {
	customerIDValue := sharedVO.NewCustomerIDFromUUID(customerID)
	return entity.NewCustomerBuilder().WithID(&customerIDValue).Build(), nil
}

type MockRestaurantRepository struct {
	activeRestaurants map[string]bool
}

func NewMockRestaurantRepository() *MockRestaurantRepository {
	return &MockRestaurantRepository{
		activeRestaurants: make(map[string]bool),
	}
}

func (m *MockRestaurantRepository) SetRestaurantActive(restaurantID string, active bool) {
	m.activeRestaurants[restaurantID] = active
}

func (m *MockRestaurantRepository) FindInformation(restaurant entity.Restaurant) (*entity.Restaurant, error) {
	restaurantID := restaurant.GetID()
	restaurantIDStr := restaurantID.GetValue().String()
	
	productID := sharedVO.NewProductID()
	price := sharedVO.NewMoney(50.00)
	product := entity.NewProductBuilder().
		WithID(&productID).
		WithName("Test Product").
		WithPrice(price).
		Build()
	
	active := true
	if isActive, exists := m.activeRestaurants[restaurantIDStr]; exists {
		active = isActive
	}
	
	return entity.NewRestaurantBuilder().
		WithID(restaurantID).
		WithProducts([]entity.Product{*product}).
		WithActiveStatus(active).
		Build(), nil
}
