package mapper

import (
	"github.com/google/uuid"
	"github.com/leninner/infrastructure/kafka/kafka-model/generated"
	"github.com/leninner/order-service/internal/domain/core/event"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OrderMessagingDataMapper struct{}

func NewOrderMessagingDataMapper() *OrderMessagingDataMapper {
	return &OrderMessagingDataMapper{}
}

func (m *OrderMessagingDataMapper) OrderCreatedEventToPaymentRequestMessage(
	orderEvent *event.OrderCreatedEvent,
) *generated.PaymentRequestProtoModel {
	order := orderEvent.Order
	return &generated.PaymentRequestProtoModel{
		Id: uuid.New().String(),
		SagaId: "",
		CustomerId: order.CustomerID.GetValue().String(),
		OrderId: order.ID.GetValue().String(),
		Price: order.Price.Amount,
		CreatedAt: timestamppb.New(orderEvent.CreatedAt),
		PaymentOrderStatus: generated.PaymentOrderStatus_PENDING,
	}
}


func (m *OrderMessagingDataMapper) OrderCancelledEventToPaymentRequestMessage(
	orderEvent *event.OrderCancelledEvent,
) *generated.PaymentRequestProtoModel {
	order := orderEvent.Order
	return &generated.PaymentRequestProtoModel{
		Id: uuid.New().String(),
		SagaId: "",
		CustomerId: order.CustomerID.GetValue().String(),
		OrderId: order.ID.GetValue().String(),
		Price: order.Price.Amount,
		CreatedAt: timestamppb.New(orderEvent.CreatedAt),
		PaymentOrderStatus: generated.PaymentOrderStatus_PAYMENT_ORDER_CANCELLED,
	}
}