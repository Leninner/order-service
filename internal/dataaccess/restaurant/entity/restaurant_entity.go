package entity

import "github.com/google/uuid"

type RestaurantModel struct {
	RestaurantID uuid.UUID `json:"restaurant_id"`
	ProductID    uuid.UUID `json:"product_id"`
	RestaurantName string `json:"restaurant_name"`
	RestaurantActive bool `json:"restaurant_active"`
	ProductName string `json:"product_name"`
	ProductPrice float64 `json:"product_price"`
	ProductAvailable bool `json:"product_available"`
}