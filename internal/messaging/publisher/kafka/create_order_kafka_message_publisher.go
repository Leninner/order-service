package kafka

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	kafkamodule "github.com/leninner/infrastructure/kafka"
	kafkamodel "github.com/leninner/infrastructure/kafka/kafka-model"
	"github.com/leninner/order-service/internal/domain/application-service/config"
	"github.com/leninner/order-service/internal/domain/core/event"
	"github.com/leninner/order-service/internal/messaging/mapper"
)

type CreateOrderKafkaMessagePublisher struct {
	kafkaModule *kafkamodule.KafkaModule
	mapper      *mapper.OrderMessagingDataMapper
	config      *config.OrderServiceConfigData
}

func NewCreateOrderKafkaMessagePublisher(
	kafkaModule *kafkamodule.KafkaModule,
	mapper *mapper.OrderMessagingDataMapper,
	config *config.OrderServiceConfigData,
) *CreateOrderKafkaMessagePublisher {
	return &CreateOrderKafkaMessagePublisher{
		kafkaModule: kafkaModule,
		mapper:      mapper,
		config:      config,
	}
}

func (p *CreateOrderKafkaMessagePublisher) Publish(event *event.OrderCreatedEvent) error {
	if p.kafkaModule == nil {
		return fmt.Errorf("kafka module not initialized")
	}

	orderID := event.GetOrder().GetID().GetValue().String()
	
	// For now, publish a simple JSON message
	// TODO: Implement proper protobuf mapping
	messageData := fmt.Sprintf(`{"order_id":"%s","event_type":"order_created"}`, orderID)
	
	err := p.kafkaModule.ProduceMessage(
		"payment-request",
		[]byte(orderID),
		[]byte(messageData),
		[]kafka.Header{
			{Key: "event_type", Value: []byte("order_created")},
			{Key: "content_type", Value: []byte("application/json")},
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish order created event: %w", err)
	}

	return nil
}

func (p *CreateOrderKafkaMessagePublisher) PublishWithHeaders(event *event.OrderCreatedEvent, headers []kafka.Header) error {
	if p.kafkaModule == nil {
		return fmt.Errorf("kafka module not initialized")
	}

	orderID := event.GetOrder().GetID().GetValue().String()
	
	// For now, publish a simple JSON message with custom headers
	messageData := fmt.Sprintf(`{"order_id":"%s","event_type":"order_created"}`, orderID)
	
	// Create Kafka message with custom headers
	kafkaMessage := kafkamodel.NewMessage("payment-request", []byte(orderID), []byte(messageData))
	for _, header := range headers {
		kafkaMessage.AddHeader(header.Key, string(header.Value))
	}

	return p.kafkaModule.GetProducer().ProduceMessage(kafkaMessage)
}