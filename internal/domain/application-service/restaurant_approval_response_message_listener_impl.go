package applicationservice

import (
	"github.com/leninner/order-service/internal/domain/application-service/dto/message"
	service "github.com/leninner/order-service/internal/domain/application-service/ports/input/service"
)

type RestaurantApprovalResponseMessageListenerImpl struct {
	orderApplicationService service.OrderApplicationService
}

func NewRestaurantApprovalResponseMessageListenerImpl(
	orderApplicationService service.OrderApplicationService,
) *RestaurantApprovalResponseMessageListenerImpl {
	return &RestaurantApprovalResponseMessageListenerImpl{orderApplicationService: orderApplicationService}
}

func (r *RestaurantApprovalResponseMessageListenerImpl) OrderApproved(
	restaurantApprovalResponse *message.RestaurantApprovalResponse,
) error {
	return nil
}

func (r *RestaurantApprovalResponseMessageListenerImpl) OrderRejected(
	restaurantApprovalResponse *message.RestaurantApprovalResponse,
) error {
	return nil
}
