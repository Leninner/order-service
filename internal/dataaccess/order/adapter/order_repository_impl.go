package adapter

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/leninner/order-service/internal/dataaccess/order/mapper"
	"github.com/leninner/order-service/internal/dataaccess/order/repository"
	"github.com/leninner/order-service/internal/domain/core/entity"
	"github.com/leninner/order-service/internal/domain/core/valueobject"
)

type OrderRepositoryImpl struct {
	orderSqlRepository    *repository.OrderRepositorySQLImpl
	orderDataAccessMapper *mapper.OrderDataAccessMapper
}

func NewOrderRepositoryImpl(db *sql.DB) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{
		orderSqlRepository:    repository.NewOrderRepositorySQLImpl(db),
		orderDataAccessMapper: mapper.NewOrderDataAccessMapper(),
	}
}

func (r *OrderRepositoryImpl) Save(order *entity.Order) (*entity.Order, error) {
	orderModel := r.orderDataAccessMapper.OrderDomainToOrderModel(order)
	orderModel, err := r.orderSqlRepository.Save(orderModel)
	if err != nil {
		return nil, err
	}

	return r.orderDataAccessMapper.OrderModelToOrderDomain(orderModel), nil
}

func (r *OrderRepositoryImpl) FindByTrackingID(trackingID uuid.UUID) (*entity.Order, error) {
	orderModel, err := r.orderSqlRepository.FindByTrackingID(valueobject.TrackingIDFromUUID(trackingID))
	if err != nil {
		return nil, err
	}

	return r.orderDataAccessMapper.OrderModelToOrderDomain(orderModel), nil
}