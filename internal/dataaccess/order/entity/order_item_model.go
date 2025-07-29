package order

import "github.com/google/uuid"

type OrderItemModel struct {
	ID        int64 `json:"id"`
	OrderID   uuid.UUID `json:"order_id"`
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int16     `json:"quantity"`
	Price     float64   `json:"price"`
	SubTotal  float64   `json:"sub_total"`
} 