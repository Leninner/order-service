package restaurantapproval

import "github.com/leninner/order-service/internal/domain/application-service/dto/message"

type RestaurantApprovalResponseMessageListener interface {
	OrderApproved(restaurantApprovalResponse *message.RestaurantApprovalResponse) error

	OrderRejected(restaurantApprovalResponse *message.RestaurantApprovalResponse) error
}
