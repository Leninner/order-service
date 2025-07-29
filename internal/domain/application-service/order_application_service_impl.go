package applicationservice

import (
	"github.com/leninner/order-service/internal/domain/application-service/dto/create"
	"github.com/leninner/order-service/internal/domain/application-service/dto/track"
	service "github.com/leninner/order-service/internal/domain/application-service/ports/input/service"
)

type OrderApplicationServiceImpl struct {
	orderCreateCommandHandler OrderCreateCommandHandler
	orderTrackCommandHandler  OrderTrackCommandHandler
}

func NewOrderApplicationService(
	orderCreateCommandHandler OrderCreateCommandHandler,
	orderTrackCommandHandler OrderTrackCommandHandler,
) service.OrderApplicationService {
	return &OrderApplicationServiceImpl{
		orderCreateCommandHandler: orderCreateCommandHandler,
		orderTrackCommandHandler:  orderTrackCommandHandler,
	}
}

func (o *OrderApplicationServiceImpl) CreateOrder(
	command create.CreateOrderCommand,
) (*create.CreateOrderResponse, error) {
	return o.orderCreateCommandHandler.Handle(command)
}

func (o *OrderApplicationServiceImpl) TrackOrder(
	query track.TrackOrderQuery,
) (*track.TrackOrderResponse, error) {
	return o.orderTrackCommandHandler.Handle(query)
}
