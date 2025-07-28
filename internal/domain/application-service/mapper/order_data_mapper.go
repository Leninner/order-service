package mapper

import (
	"time"

	"github.com/google/uuid"
	"github.com/leninner/order-service/internal/domain/application-service/dto/create"
	"github.com/leninner/order-service/internal/domain/application-service/dto/message"
	"github.com/leninner/order-service/internal/domain/application-service/dto/track"
	"github.com/leninner/order-service/internal/domain/core/entity"
	"github.com/leninner/order-service/internal/domain/core/event"
	"github.com/leninner/order-service/internal/domain/core/valueobject"
	sharedVO "github.com/leninner/shared/domain/valueobject"
)

type OrderDataMapper struct{}

func NewOrderDataMapper() *OrderDataMapper {
	return &OrderDataMapper{}
}

func (m *OrderDataMapper) CreateOrderCommandToRestaurant(
	createOrderCommand *create.CreateOrderCommand,
) (*entity.Restaurant, error) {
	restaurantID := sharedVO.NewRestaurantIDFromUUID(createOrderCommand.RestaurantID)
	products := make([]entity.Product, 0, len(createOrderCommand.Items))

	for _, orderItem := range createOrderCommand.Items {
		productID := sharedVO.NewProductIDFromUUID(orderItem.ProductID)
		product := entity.NewProductBuilder().
			WithID(&productID).
			Build()

		products = append(products, *product)
	}

	return entity.NewRestaurantBuilder().
		WithID(&restaurantID).
		WithProducts(products).
		Build(), nil
}

func (m *OrderDataMapper) CreateOrderCommandToOrder(
	createOrderCommand *create.CreateOrderCommand,
) (*entity.Order, error) {
	customerID := sharedVO.NewCustomerIDFromUUID(createOrderCommand.CustomerID)
	restaurantID := sharedVO.NewRestaurantIDFromUUID(createOrderCommand.RestaurantID)

	deliveryAddress := m.orderAddressToStreetAddress(createOrderCommand.Address)
	price := sharedVO.NewMoney(*createOrderCommand.Price)
	items, err := m.orderItemsToOrderItemEntities(createOrderCommand.Items)
	if err != nil {
		return nil, err
	}

	return entity.NewOrderBuilder().
		WithCustomerID(customerID).
		WithRestaurantID(restaurantID).
		WithDeliveryAddress(deliveryAddress).
		WithPrice(*price).
		WithItems(items).
		Build(), nil
}

func (m *OrderDataMapper) OrderToCreateOrderResponse(order *entity.Order, message string) *create.CreateOrderResponse {
	trackingID := order.TrackingID.GetValue()
	orderStatus := *order.OrderStatus

	return &create.CreateOrderResponse{
		OrderTrackingID: &trackingID,
		OrderStatus:     orderStatus,
		Message:         &message,
	}
}

func (m *OrderDataMapper) OrderToTrackOrderResponse(order *entity.Order) *track.TrackOrderResponse {
	trackingID := order.TrackingID.GetValue()
	orderStatus := *order.OrderStatus

	return &track.TrackOrderResponse{
		OrderTrackingID: &trackingID,
		OrderStatus:     orderStatus,
		FailureMessages: order.FailureMessages,
	}
}

func (m *OrderDataMapper) OrderCreatedEventToOrderPaymentEventPayload(
	orderCreatedEvent *event.OrderCreatedEvent,
) *OrderPaymentEventPayload {
	order := orderCreatedEvent.GetOrder()
	customerID := order.CustomerID.GetValue().String()
	orderID := order.GetID().GetValue().String()
	price := order.Price.Amount

	return &OrderPaymentEventPayload{
		CustomerID:         customerID,
		OrderID:            orderID,
		Price:              price,
		CreatedAt:          orderCreatedEvent.CreatedAt,
		PaymentOrderStatus: string(sharedVO.PaymentStatusPending),
	}
}

func (m *OrderDataMapper) OrderCancelledEventToOrderPaymentEventPayload(
	orderCancelledEvent *event.OrderCancelledEvent,
) *OrderPaymentEventPayload {
	order := orderCancelledEvent.GetOrder()
	customerID := order.CustomerID.GetValue().String()
	orderID := order.GetID().GetValue().String()
	price := order.Price.Amount

	return &OrderPaymentEventPayload{
		CustomerID:         customerID,
		OrderID:            orderID,
		Price:              price,
		CreatedAt:          orderCancelledEvent.CreatedAt,
		PaymentOrderStatus: string(sharedVO.PaymentStatusCancelled),
	}
}

func (m *OrderDataMapper) OrderPaidEventToOrderApprovalEventPayload(
	orderPaidEvent *event.OrderPaidEvent,
) *OrderApprovalEventPayload {
	order := orderPaidEvent.GetOrder()
	orderID := order.GetID().GetValue().String()
	restaurantID := order.RestaurantID.GetValue().String()
	price := order.Price.Amount
	products := make([]OrderApprovalEventProduct, 0, len(order.Items))

	for _, orderItem := range order.Items {
		productID := orderItem.Product.GetID().ID.String()
		product := OrderApprovalEventProduct{
			ID:       productID,
			Quantity: orderItem.Quantity,
		}
		products = append(products, product)
	}

	return &OrderApprovalEventPayload{
		OrderID:               orderID,
		RestaurantID:          restaurantID,
		RestaurantOrderStatus: string(sharedVO.OrderStatusPaid),
		Products:              products,
		Price:                 price,
		CreatedAt:             orderPaidEvent.CreatedAt,
	}
}

func (m *OrderDataMapper) CustomerModelToCustomer(customerModel *message.CustomerModel) *entity.Customer {
	customerID := sharedVO.NewCustomerIDFromUUID(&customerModel.ID)

	return entity.NewCustomerBuilder().
		WithID(&customerID).
		Build()
}

func (m *OrderDataMapper) orderItemsToOrderItemEntities(
	orderItems []create.OrderItem,
) ([]entity.OrderItem, error) {
	items := make([]entity.OrderItem, 0, len(orderItems))

	for _, orderItem := range orderItems {
		productID := sharedVO.NewProductIDFromUUID(orderItem.ProductID)
		price := sharedVO.NewMoney(*orderItem.Price)
		product := entity.NewProductBuilder().
			WithID(&productID).
			WithPrice(price).
			Build()

		subTotal := price.Multiply(int32(*orderItem.Quantity))

		item := entity.OrderItem{
			Product:  *product,
			Price:    *price,
			Quantity: int16(*orderItem.Quantity),
			SubTotal: *subTotal,
		}
		items = append(items, item)
	}

	return items, nil
}

func (m *OrderDataMapper) orderAddressToStreetAddress(orderAddress create.OrderAddress) valueobject.StreetAddress {
	return valueobject.StreetAddress{
		ID:         uuid.New(),
		Street:     *orderAddress.Street,
		PostalCode: *orderAddress.PostalCode,
		City:       *orderAddress.City,
	}
}

type OrderPaymentEventPayload struct {
	CustomerID         string    `json:"customerId"`
	OrderID            string    `json:"orderId"`
	Price              float64   `json:"price"`
	CreatedAt          time.Time `json:"createdAt"`
	PaymentOrderStatus string    `json:"paymentOrderStatus"`
}

type OrderApprovalEventPayload struct {
	OrderID               string                      `json:"orderId"`
	RestaurantID          string                      `json:"restaurantId"`
	RestaurantOrderStatus string                      `json:"restaurantOrderStatus"`
	Products              []OrderApprovalEventProduct `json:"products"`
	Price                 float64                     `json:"price"`
	CreatedAt             time.Time                   `json:"createdAt"`
}

type OrderApprovalEventProduct struct {
	ID       string `json:"id"`
	Quantity int16  `json:"quantity"`
}
