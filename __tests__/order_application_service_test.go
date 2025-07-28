package tests

import (
	"testing"

	"github.com/google/uuid"
	applicationservice "github.com/leninner/order-service/internal/domain/application-service"
	"github.com/leninner/order-service/internal/domain/application-service/dto/create"
	"github.com/leninner/order-service/internal/domain/application-service/mapper"
	service "github.com/leninner/order-service/internal/domain/application-service/ports/input/service"
	sharedVO "github.com/leninner/shared/domain/valueobject"
	"github.com/leninner/shared/logger"
)

type OrderApplicationServiceTest struct {
	orderApplicationService service.OrderApplicationService
	orderDataMapper         *mapper.OrderDataMapper
	config                  *OrderTestConfiguration
}

func setupOrderApplicationServiceTest(t *testing.T) *OrderApplicationServiceTest {
	logger, err := logger.NewTestLogger("order-service-test")
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}

	config := NewOrderTestConfiguration(logger)
	orderDataMapper := mapper.NewOrderDataMapper()

	orderCreateHelper := applicationservice.NewOrderCreateHelper(
		config.OrderRepository,
		config.CustomerRepository,
		config.RestaurantRepository,
		*orderDataMapper,
		config.OrderDomainService,
		logger,
	)

	orderCreateCommandHandler := applicationservice.NewOrderCreateCommandHandler(
		orderCreateHelper,
		*orderDataMapper,
		config.PaymentRequestMessagePublisher,
	)

	orderTrackCommandHandler := applicationservice.NewOrderTrackCommandHandler(
		*orderDataMapper,
		config.OrderRepository,
	)

	orderApplicationService := applicationservice.New(
		*orderCreateCommandHandler,
		*orderTrackCommandHandler,
	)

	return &OrderApplicationServiceTest{
		orderApplicationService: orderApplicationService,
		orderDataMapper:         orderDataMapper,
		config:                  config,
	}
}

func TestOrderApplicationService_CreateOrder(t *testing.T) {
	test := setupOrderApplicationServiceTest(t)

	customerID := uuid.MustParse("d215b5f8-0249-4dc5-89a3-51fd148cfb41")
	restaurantID := uuid.MustParse("d215b5f8-0249-4dc5-89a3-51fd148cfb45")
	productID := uuid.MustParse("d215b5f8-0249-4dc5-89a3-51fd148cfb48")

	price := 200.00
	street := "street_1"
	postalCode := "1000AB"
	city := "Paris"
	state := "State"
	country := "Country"

	createOrderCommand := create.NewCreateOrderCommand(
		create.WithCustomerID(customerID),
		create.WithRestaurantID(restaurantID),
		create.WithPrice(price),
		create.WithAddress(create.OrderAddress{
			Street:     &street,
			PostalCode: &postalCode,
			City:       &city,
			State:      &state,
			Country:    &country,
		}),
		create.WithItems([]create.OrderItem{
			{
				ProductID: &productID,
				Quantity:  intPtr(1),
				Price:     float64Ptr(50.00),
			},
			{
				ProductID: &productID,
				Quantity:  intPtr(3),
				Price:     float64Ptr(50.00),
			},
		}),
	)

	response, err := test.orderApplicationService.CreateOrder(*createOrderCommand)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if response == nil {
		t.Fatal("Expected response, got nil")
	}
	if response.OrderStatus != sharedVO.OrderStatusPending {
		t.Errorf("Expected order status %s, got %s", sharedVO.OrderStatusPending, response.OrderStatus)
	}
	if response.Message == nil || *response.Message != "Order created successfully" {
		t.Errorf("Expected message 'Order created successfully', got %v", response.Message)
	}
	if response.OrderTrackingID == nil {
		t.Error("Expected order tracking ID, got nil")
	}
}

func TestOrderApplicationService_CreateOrderWithWrongTotalPrice(t *testing.T) {
	test := setupOrderApplicationServiceTest(t)

	customerID := uuid.MustParse("d215b5f8-0249-4dc5-89a3-51fd148cfb41")
	restaurantID := uuid.MustParse("d215b5f8-0249-4dc5-89a3-51fd148cfb45")
	productID := uuid.MustParse("d215b5f8-0249-4dc5-89a3-51fd148cfb48")

	price := 250.00
	street := "street_1"
	postalCode := "1000AB"
	city := "Paris"
	state := "State"
	country := "Country"

	createOrderCommand := create.NewCreateOrderCommand(
		create.WithCustomerID(customerID),
		create.WithRestaurantID(restaurantID),
		create.WithPrice(price),
		create.WithAddress(create.OrderAddress{
			Street:     &street,
			PostalCode: &postalCode,
			City:       &city,
			State:      &state,
			Country:    &country,
		}),
		create.WithItems([]create.OrderItem{
			{
				ProductID: &productID,
				Quantity:  intPtr(1),
				Price:     float64Ptr(50.00),
			},
			{
				ProductID: &productID,
				Quantity:  intPtr(3),
				Price:     float64Ptr(50.00),
			},
		}),
	)

	response, err := test.orderApplicationService.CreateOrder(*createOrderCommand)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if response != nil {
		t.Errorf("Expected nil response, got %v", response)
	}
	if err.Error() != "total price is not equal to the price of the order" {
		t.Errorf("Expected error message 'total price is not equal to the price of the order', got %s", err.Error())
	}
}

func TestOrderApplicationService_CreateOrderWithPassiveRestaurant(t *testing.T) {
	test := setupOrderApplicationServiceTest(t)

	customerID := uuid.MustParse("d215b5f8-0249-4dc5-89a3-51fd148cfb41")
	restaurantID := uuid.MustParse("d215b5f8-0249-4dc5-89a3-51fd148cfb45")
	productID := uuid.MustParse("d215b5f8-0249-4dc5-89a3-51fd148cfb48")

	test.config.RestaurantRepository.(*MockRestaurantRepository).SetRestaurantActive(restaurantID.String(), false)

	price := 200.00
	street := "street_1"
	postalCode := "1000AB"
	city := "Paris"
	state := "State"
	country := "Country"

	createOrderCommand := create.NewCreateOrderCommand(
		create.WithCustomerID(customerID),
		create.WithRestaurantID(restaurantID),
		create.WithPrice(price),
		create.WithAddress(create.OrderAddress{
			Street:     &street,
			PostalCode: &postalCode,
			City:       &city,
			State:      &state,
			Country:    &country,
		}),
		create.WithItems([]create.OrderItem{
			{
				ProductID: &productID,
				Quantity:  intPtr(1),
				Price:     float64Ptr(50.00),
			},
			{
				ProductID: &productID,
				Quantity:  intPtr(3),
				Price:     float64Ptr(50.00),
			},
		}),
	)

	response, err := test.orderApplicationService.CreateOrder(*createOrderCommand)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if response != nil {
		t.Errorf("Expected nil response, got %v", response)
	}
	if err.Error() != "restaurant with id "+restaurantID.String()+" is not active" {
		t.Errorf("Expected error message about inactive restaurant, got %s", err.Error())
	}
}

func intPtr(i int) *int {
	return &i
}

func float64Ptr(f float64) *float64 {
	return &f
}
