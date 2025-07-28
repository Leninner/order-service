package restaurantapproval

import (
	"github.com/leninner/order-service/internal/domain/core/event"
	"github.com/leninner/shared/domain/event/publisher"
)

type OrderPaidRestaurantApprovalRequestMessagePublisher interface {
	publisher.DomainEventPublisher[*event.OrderPaidEvent]
}
