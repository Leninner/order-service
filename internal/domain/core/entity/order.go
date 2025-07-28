package entity

import (
	"github.com/leninner/order-service/internal/domain/core/exception"
	vo "github.com/leninner/order-service/internal/domain/core/valueobject"
	"github.com/leninner/shared/domain/entity"
	sharedVO "github.com/leninner/shared/domain/valueobject"
)

type Order struct {
	entity.AggregateRoot[*sharedVO.OrderID]

	CustomerID      sharedVO.CustomerID
	RestaurantID    sharedVO.RestaurantID
	DeliveryAddress vo.StreetAddress
	Price           sharedVO.Money
	Items           []OrderItem

	TrackingID      vo.TrackingID
	OrderStatus     *sharedVO.OrderStatus
	FailureMessages []string
}

func (o *Order) InitializeOrder() {
	orderID := sharedVO.NewOrderID()
	o.SetID(&orderID)

	o.TrackingID = vo.NewTrackingID()

	status := sharedVO.OrderStatusPending
	o.OrderStatus = &status

	o.initializeOrderItems()
}

func (o *Order) initializeOrderItems() {
	var itemId int64 = 1

	desiredItemId, err := vo.NewOrderItemID(itemId)
	if err != nil {
		panic(err)
	}

	for i := range o.Items {
		o.Items[i].initializeOrderItem(*o.GetID(), desiredItemId)
		itemId++
	}
}

func (o *Order) ValidateOrder() error {
	err := o.validateInitialOrder()
	if err != nil {
		return err
	}

	err = o.validateTotalPrice()
	if err != nil {
		return err
	}

	err = o.validateItemsPrice()
	if err != nil {
		return err
	}

	return nil
}

func (o *Order) Pay() error {
	if *o.OrderStatus != sharedVO.OrderStatusPending {
		return exception.NewOrderDomainException("order is not in pending state for payment operation")
	}

	status := sharedVO.OrderStatusPaid
	o.OrderStatus = &status

	return nil
}

func (o *Order) Approve() error {
	if *o.OrderStatus != sharedVO.OrderStatusPaid {
		return exception.NewOrderDomainException("order is not in paid state for approval operation")
	}

	status := sharedVO.OrderStatusApproved
	o.OrderStatus = &status

	return nil
}

func (o *Order) InitCancel(failureMessages []string) error {
	if *o.OrderStatus != sharedVO.OrderStatusPaid {
		return exception.NewOrderDomainException("order is not in paid state for init cancel operation")
	}

	status := sharedVO.OrderStatusCancelling
	o.OrderStatus = &status
	o.updateFailureMessages(failureMessages)

	return nil
}

func (o *Order) Cancel(failureMessages []string) error {
	if !(*o.OrderStatus == sharedVO.OrderStatusCancelling || *o.OrderStatus == sharedVO.OrderStatusPending) {
		return exception.NewOrderDomainException("order is not in cancelling or pending state for cancel operation")
	}

	status := sharedVO.OrderStatusCancelled
	o.OrderStatus = &status
	o.updateFailureMessages(failureMessages)

	return nil
}

func (o *Order) validateInitialOrder() error {
	if o.OrderStatus != nil || o.GetID() != nil {
		return exception.NewOrderDomainException("order is not in initial state")
	}

	return nil
}

func (o *Order) validateTotalPrice() error {
	if !o.Price.IsGreaterThanZero() {
		err := exception.NewOrderDomainException("total price must be greater than zero")
		return err
	}

	return nil
}

func (o *Order) validateItemsPrice() error {
	totalPrice := &sharedVO.Money{}

	for _, item := range o.Items {
		err := o.validateItemPrice(item)

		if err != nil {
			return err
		}

		totalPrice = totalPrice.Add(&item.SubTotal)
	}

	if !totalPrice.Equals(&o.Price) {
		return exception.NewOrderDomainException("total price is not equal to the price of the order")
	}

	return nil
}

func (o *Order) validateItemPrice(item OrderItem) error {
	if !item.IsPriceValid() {
		return exception.NewOrderDomainException("item price must be greater than zero")
	}

	return nil
}

func (o *Order) updateFailureMessages(failureMessages []string) {
	o.FailureMessages = append(o.FailureMessages, failureMessages...)
}
