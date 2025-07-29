package entity

import (
	vo "github.com/leninner/order-service/internal/domain/core/valueobject"
	"github.com/leninner/shared/domain/entity"
	sharedVO "github.com/leninner/shared/domain/valueobject"
)

type OrderItem struct {
	entity.Entity[vo.OrderItemID]

	OrderID  sharedVO.OrderID
	Product  Product
	Quantity int16
	Price    sharedVO.Money
	SubTotal sharedVO.Money
}

func (o *OrderItem) initializeOrderItem(orderID sharedVO.OrderID, orderItemID vo.OrderItemID) {
	o.SetID(orderItemID)
	o.OrderID = orderID
	o.Product = Product{}
	o.Quantity = 0
	o.Price = sharedVO.Money{}
	o.SubTotal = sharedVO.Money{}
}

func (o *OrderItem) IsPriceValid() bool {
	return o.Price.IsGreaterThanZero() && o.Product.Price.Equals(&o.Price) && o.Price.Multiply(int32(o.Quantity)).Equals(&o.SubTotal)
}

func (o *OrderItem) GetProductID() sharedVO.ProductID {
	return o.Product.GetID()
}

func (o *OrderItem) GetSubTotal() *sharedVO.Money {
	return &o.SubTotal
}
