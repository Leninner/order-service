package order

import "github.com/google/uuid"

type OrderModel struct {
	ID              uuid.UUID        `json:"id"`
	CustomerID      uuid.UUID        `json:"customer_id"`
	RestaurantID    uuid.UUID        `json:"restaurant_id"`
	TrackingID      uuid.UUID        `json:"tracking_id"`
	Price           float64          `json:"price"`
	OrderStatus     string           `json:"order_status"`
	FailureMessages []string         `json:"failure_messages"`
	Address         *AddressModel    `json:"address"`
	Items           []OrderItemModel `json:"items"`
}

