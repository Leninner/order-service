package core

import (
	"errors"

	"github.com/leninner/order-service/internal/domain/core/entity"
	"github.com/leninner/order-service/internal/domain/core/event"
	sharedVO "github.com/leninner/shared/domain/valueobject"
	"github.com/leninner/shared/logger"
	"go.uber.org/zap"
)

// OrderDomainServiceImpl implements the OrderDomainService interface and provides the domain logic for the order service
type OrderDomainServiceImpl struct {
	logger *logger.Logger
}

func NewOrderDomainServiceImpl(logger *logger.Logger) *OrderDomainServiceImpl {
	return &OrderDomainServiceImpl{
		logger: logger,
	}
}

func (s *OrderDomainServiceImpl) ValidateAndInitiateOrder(
	order *entity.Order,
	restaurant *entity.Restaurant,
) (*event.OrderCreatedEvent, error) {
	err := s.ValidateRestaurant(restaurant)
	if err != nil {
		return nil, err
	}

	s.SetOrderProductInformation(order, restaurant)

	err = order.ValidateOrder()
	if err != nil {
		return nil, err
	}

	order.InitializeOrder()

	s.logger.Info(
		"Order validated and initiated successfully",
		zap.String("orderId", order.GetID().GetValue().String()),
	)

	return event.NewOrderCreatedEvent(order), nil
}

func (s *OrderDomainServiceImpl) PayOrder(order *entity.Order) (*event.OrderPaidEvent, error) {
	s.logger.Info(
		"Processing order payment",
		zap.String("orderId", order.GetID().GetValue().String()),
	)

	err := order.Pay()
	if err != nil {
		return nil, err
	}

	return event.NewOrderPaidEvent(order), nil
}

func (s *OrderDomainServiceImpl) ApproveOrder(order *entity.Order) error {
	s.logger.Info("Approving order",
		zap.String("orderId", order.GetID().GetValue().String()),
		zap.String("orderStatus", string(*order.OrderStatus)),
	)

	err := order.Approve()
	if err != nil {
		return err
	}

	return nil
}

func (s *OrderDomainServiceImpl) CancelOrderPayment(
	order *entity.Order,
	failureMessages []string,
) (*event.OrderCancelledEvent, error) {
	s.logger.Error(
		"Order payment cancelled",
		zap.String("orderId", order.GetID().GetValue().String()),
		zap.Any("failureMessages", failureMessages),
	)

	err := order.InitCancel(failureMessages)
	if err != nil {
		return nil, err
	}

	return event.NewOrderCancelledEvent(order), nil
}

func (s *OrderDomainServiceImpl) CancelOrder(order *entity.Order, failureMessages []string) error {
	s.logger.Error("Order cancelled",
		zap.String("orderId", order.GetID().GetValue().String()),
		zap.Any("failureMessages", failureMessages),
	)

	err := order.Cancel(failureMessages)
	if err != nil {
		return err
	}

	return nil
}

func (s *OrderDomainServiceImpl) ValidateRestaurant(restaurant *entity.Restaurant) error {
	if !restaurant.Active {
		s.logger.Error(
			"Restaurant is not active",
			zap.String("restaurantId", restaurant.GetID().GetValue().String()),
		)

		return errors.New("restaurant with id " + restaurant.GetID().GetValue().String() + " is not active")
	}

	return nil
}

func (s *OrderDomainServiceImpl) SetOrderProductInformation(
	order *entity.Order,
	restaurant *entity.Restaurant,
) {
	restaurantProductsMap := make(map[sharedVO.ProductID]*entity.Product)

	for i := range restaurant.Products {
		restaurantProductsMap[restaurant.Products[i].ID] = &restaurant.Products[i]
	}

	for i := range order.Items {
		orderItem := &order.Items[i]

		if restaurantProduct, exists := restaurantProductsMap[orderItem.Product.ID]; exists {
			orderItem.Product.UpdateWithConfirmedNameAndPrice(
				restaurantProduct.Name,
				restaurantProduct.Price,
			)
		}
	}
}
