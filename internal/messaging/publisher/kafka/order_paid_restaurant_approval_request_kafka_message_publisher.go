package kafka

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	kafkamodule "github.com/leninner/infrastructure/kafka"
	"github.com/leninner/order-service/internal/domain/application-service/config"
	"github.com/leninner/order-service/internal/domain/core/event"
	"github.com/leninner/order-service/internal/messaging/mapper"
)

type OrderPaidRestaurantApprovalRequestKafkaMessagePublisher struct {
	kafkaModule *kafkamodule.KafkaModule
	mapper      *mapper.OrderMessagingDataMapper
	config      *config.OrderServiceConfigData
}

func NewOrderPaidRestaurantApprovalRequestKafkaMessagePublisher(
	kafkaModule *kafkamodule.KafkaModule,
	mapper *mapper.OrderMessagingDataMapper,
	config *config.OrderServiceConfigData,
) *OrderPaidRestaurantApprovalRequestKafkaMessagePublisher {
	return &OrderPaidRestaurantApprovalRequestKafkaMessagePublisher{
		kafkaModule: kafkaModule,
		mapper:      mapper,
		config:      config,
	}
}

func (p *OrderPaidRestaurantApprovalRequestKafkaMessagePublisher) Publish(event *event.OrderPaidEvent) error {
	if p.kafkaModule == nil {
		return fmt.Errorf("kafka module not initialized")
	}

	orderID := event.GetOrder().GetID().GetValue().String()
	
	messageData := fmt.Sprintf(`{"order_id":"%s","event_type":"order_paid"}`, orderID)
	
	err := p.kafkaModule.ProduceMessage(
		"restaurant-approval-request",
		[]byte(orderID),
		[]byte(messageData),
		[]kafka.Header{
			{Key: "event_type", Value: []byte("order_paid")},
			{Key: "content_type", Value: []byte("application/json")},
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish order paid event: %w", err)
	}

	return nil
} 