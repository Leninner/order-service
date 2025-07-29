package repository

import (
	"github.com/leninner/order-service/internal/domain/core/entity"
	"github.com/leninner/order-service/internal/domain/core/exception"
	"github.com/leninner/order-service/internal/domain/core/valueobject"
)

type OrderRepositoryImpl struct {
	orders map[string]*entity.Order
}

func NewOrderRepositoryImpl() *OrderRepositoryImpl {
	return &OrderRepositoryImpl{
		orders: make(map[string]*entity.Order),
	}
}

func (r *OrderRepositoryImpl) Save(order *entity.Order) (*entity.Order, error) {
	r.orders[order.GetID().GetValue().String()] = order
	return order, nil
}

func (r *OrderRepositoryImpl) FindByTrackingID(orderTrackingID valueobject.TrackingID) (*entity.Order, error) {
	for _, order := range r.orders {
		if order.TrackingID.GetValue() == orderTrackingID.GetValue() {
			return order, nil
		}
	}
	return nil, exception.NewOrderDomainException("order with tracking id " + orderTrackingID.GetValue().String() + " not found")
}
