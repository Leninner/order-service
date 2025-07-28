package applicationservice

import (
	"github.com/leninner/order-service/internal/domain/application-service/dto/create"
	"github.com/leninner/order-service/internal/domain/application-service/mapper"
	"github.com/leninner/order-service/internal/domain/application-service/ports/output/message/publisher/payment"
	"github.com/leninner/order-service/internal/domain/core/exception"
	"github.com/leninner/shared/utils/validator"
)

type OrderCreateCommandHandler struct {
	orderCreateHelper                          *OrderCreateHelper
	orderDataMapper                            mapper.OrderDataMapper
	orderCreatedPaymentRequestMessagePublisher payment.OrderCreatedPaymentRequestMessagePublisher
}

func NewOrderCreateCommandHandler(
	orderCreateHelper *OrderCreateHelper,
	orderDataMapper mapper.OrderDataMapper,
	orderCreatedPaymentRequestMessagePublisher payment.OrderCreatedPaymentRequestMessagePublisher,
) *OrderCreateCommandHandler {
	return &OrderCreateCommandHandler{
		orderCreateHelper: orderCreateHelper,
		orderDataMapper:   orderDataMapper,
		orderCreatedPaymentRequestMessagePublisher: orderCreatedPaymentRequestMessagePublisher,
	}
}

func (h *OrderCreateCommandHandler) Handle(command create.CreateOrderCommand) (*create.CreateOrderResponse, error) {
	v := validator.New()
	if create.ValidateOrderCommand(v, &command); !v.Valid() {
		for _, message := range v.Errors {
			return nil, exception.NewOrderDomainException(message)
		}
		return nil, exception.NewOrderDomainException("validation failed")
	}

	orderCreatedEvent, err := h.orderCreateHelper.PersistOrder(command)
	if err != nil {
		return nil, err
	}

	h.orderCreatedPaymentRequestMessagePublisher.Publish(orderCreatedEvent)
	return h.orderDataMapper.OrderToCreateOrderResponse(orderCreatedEvent.GetOrder(), "Order created successfully"), nil
}
