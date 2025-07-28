package applicationservice

import (
	"github.com/leninner/order-service/internal/domain/application-service/dto/track"
	"github.com/leninner/order-service/internal/domain/application-service/mapper"
	repository "github.com/leninner/order-service/internal/domain/application-service/ports/output/repository"
	"github.com/leninner/order-service/internal/domain/core/exception"
	"github.com/leninner/order-service/internal/domain/core/valueobject"
)

type OrderTrackCommandHandler struct {
	orderDataMapper mapper.OrderDataMapper
	orderRepository repository.OrderRepository
}

func NewOrderTrackCommandHandler(
	orderDataMapper mapper.OrderDataMapper,
	orderRepository repository.OrderRepository,
) *OrderTrackCommandHandler {
	return &OrderTrackCommandHandler{orderDataMapper: orderDataMapper, orderRepository: orderRepository}
}

func (h *OrderTrackCommandHandler) Handle(command track.TrackOrderQuery) (*track.TrackOrderResponse, error) {
	order, err := h.orderRepository.FindByTrackingID(valueobject.TrackingIDFromUUID(command.OrderTrackingID))
	if err != nil {
		return nil, exception.NewOrderDomainException("order with tracking id " + command.OrderTrackingID.String() + " not found")
	}

	return h.orderDataMapper.OrderToTrackOrderResponse(order), nil
}
