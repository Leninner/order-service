package entity

import sharedVO "github.com/leninner/shared/domain/valueobject"

type RestaurantBuilder struct {
	restaurant *Restaurant
}

func NewRestaurantBuilder() *RestaurantBuilder {
	return &RestaurantBuilder{
		restaurant: &Restaurant{
			Products: make([]Product, 0),
			Active:   false,
		},
	}
}

func (rb *RestaurantBuilder) WithID(id *sharedVO.RestaurantID) *RestaurantBuilder {
	rb.restaurant.ID = id
	return rb
}

func (rb *RestaurantBuilder) WithProducts(products []Product) *RestaurantBuilder {
	rb.restaurant.Products = products
	return rb
}

func (rb *RestaurantBuilder) AddProduct(product Product) *RestaurantBuilder {
	rb.restaurant.Products = append(rb.restaurant.Products, product)
	return rb
}

func (rb *RestaurantBuilder) WithActiveStatus(active bool) *RestaurantBuilder {
	rb.restaurant.Active = active
	return rb
}

func (rb *RestaurantBuilder) Build() *Restaurant {
	return rb.restaurant
}
