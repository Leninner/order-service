package service

import (
	"github.com/leninner/order-service/internal/domain/application-service/dto/create"
	"github.com/leninner/order-service/internal/domain/application-service/dto/track"
)

type OrderApplicationService interface {
	CreateOrder(create.CreateOrderCommand) (*create.CreateOrderResponse, error)

	TrackOrder(track.TrackOrderQuery) (*track.TrackOrderResponse, error)
}
