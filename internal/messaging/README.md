# Order Service Messaging Module

## Overview

This document describes the messaging module implementation for the Order Service that uses the refactored Kafka module for event publishing and consumption.

## Architecture

```
messaging/
├── messaging.module.go           # Main messaging module
├── events/
│   └── order.publisher.go       # Order event publisher
├── handlers/
│   └── order.handler.go         # Order event handler
├── publisher/
│   └── kafka/                   # Kafka-specific publishers
├── listener/
│   └── kafka/                   # Kafka-specific listeners
├── mapper/                      # Message mappers
└── example/
    └── usage.example.go         # Usage examples
```

## Key Features

### 1. Unified Messaging Module
- **Single entry point**: `MessagingModule` provides unified access to all messaging functionality
- **Automatic initialization**: One-line setup with automatic Kafka module initialization
- **Dependency injection**: Proper DI patterns for testability
- **Health checks**: Built-in health monitoring

### 2. Event Publishing
- **Order events**: Support for order created, paid, and cancelled events
- **Protocol Buffers**: Built-in support for protobuf message serialization
- **Custom data**: Flexible publishing of custom data structures
- **Headers support**: Custom headers for message routing and metadata

### 3. Event Consumption
- **Event handlers**: Structured event handling with type safety
- **Message routing**: Automatic routing based on event type
- **Error handling**: Comprehensive error handling and recovery
- **Graceful shutdown**: Proper cleanup and resource management

## Quick Start

### Basic Usage

```go
package main

import (
    "github.com/leninner/order-service/internal/messaging"
)

func main() {
    // Initialize messaging module
    messagingModule, err := messaging.NewMessagingModule()
    if err != nil {
        log.Fatalf("Failed to initialize: %v", err)
    }
    defer messagingModule.Close()

    // Subscribe to order events
    err = messagingModule.SubscribeToOrderEvents()
    if err != nil {
        log.Fatalf("Failed to subscribe: %v", err)
    }

    // Start consuming events
    messagingModule.StartConsumingOrderEvents()

    // Get publisher for sending events
    publisher := messagingModule.GetOrderEventPublisher()
}
```

### Custom Configuration

```go
package main

import (
    kafkaconfigdata "github.com/leninner/infrastructure/kafka/kafka-config-data"
    "github.com/leninner/order-service/internal/messaging"
)

func main() {
    // Custom Kafka configuration
    kafkaConfig := &kafkaconfigdata.KafkaConfigData{
        Environment:      kafkaconfigdata.Production,
        BootstrapServers: "kafka-prod:9092",
        SecurityProtocol: "SASL_SSL",
        SASLUsername:     "kafka-user",
        SASLPassword:     "kafka-password",
    }

    producerConfig := &kafkaconfigdata.KafkaProducerConfigData{
        CompressionType: "snappy",
        Acks:           "all",
        BatchSize:      32768,
    }

    consumerConfig := &kafkaconfigdata.KafkaConsumerConfigData{
        GroupID:         "order-service-prod",
        AutoOffsetReset: "earliest",
    }

    messagingModule, err := messaging.NewMessagingModuleWithConfig(
        kafkaConfig,
        producerConfig,
        consumerConfig,
    )
    if err != nil {
        log.Fatalf("Failed to initialize: %v", err)
    }
    defer messagingModule.Close()
}
```

## API Reference

### MessagingModule

```go
type MessagingModule struct {
    // ... internal fields
}

// Constructors
func NewMessagingModule() (*MessagingModule, error)
func NewMessagingModuleWithConfig(config, producerConfig, consumerConfig) (*MessagingModule, error)

// Event Management
func (m *MessagingModule) SubscribeToOrderEvents() error
func (m *MessagingModule) StartConsumingOrderEvents()

// Accessors
func (m *MessagingModule) GetKafkaModule() *KafkaModule
func (m *MessagingModule) GetOrderEventPublisher() *OrderEventPublisher
func (m *MessagingModule) GetOrderEventHandler() *OrderEventHandler

// Health and Management
func (m *MessagingModule) Health() error
func (m *MessagingModule) Close()
```

### OrderEventPublisher

```go
type OrderEventPublisher struct {
    // ... internal fields
}

// Event Publishing
func (p *OrderEventPublisher) PublishOrderCreated(event *OrderCreatedEvent) error
func (p *OrderEventPublisher) PublishOrderPaid(event *OrderPaidEvent) error
func (p *OrderEventPublisher) PublishOrderCancelled(event *OrderCancelledEvent) error
func (p *OrderEventPublisher) PublishWithCustomData(topic, orderID string, data interface{}) error
```

### OrderEventHandler

```go
type OrderEventHandler struct {
    // ... internal fields
}

// Event Handling
func (h *OrderEventHandler) HandleMessage(message *Message) error
func (h *OrderEventHandler) GetMessageHandler() MessageHandler
```

## Event Publishing Examples

### Publishing Order Events

```go
// Get the publisher
publisher := messagingModule.GetOrderEventPublisher()

// Publish order created event
orderCreatedEvent := event.NewOrderCreatedEvent(order)
err := publisher.PublishOrderCreated(orderCreatedEvent)
if err != nil {
    log.Printf("Failed to publish order created event: %v", err)
}

// Publish order paid event
orderPaidEvent := event.NewOrderPaidEvent(order)
err = publisher.PublishOrderPaid(orderPaidEvent)
if err != nil {
    log.Printf("Failed to publish order paid event: %v", err)
}

// Publish order cancelled event
orderCancelledEvent := event.NewOrderCancelledEvent(order)
err = publisher.PublishOrderCancelled(orderCancelledEvent)
if err != nil {
    log.Printf("Failed to publish order cancelled event: %v", err)
}
```

### Publishing Custom Data

```go
// Publish custom order data
customData := map[string]interface{}{
    "order_id":     "order-123",
    "customer_id":  "customer-456",
    "restaurant_id": "restaurant-789",
    "total_amount": 45.99,
    "items": []map[string]interface{}{
        {"name": "Pizza", "quantity": 2, "price": 15.99},
        {"name": "Salad", "quantity": 1, "price": 8.99},
    },
}

err := publisher.PublishWithCustomData("custom-orders", "order-123", customData)
if err != nil {
    log.Printf("Failed to publish custom event: %v", err)
}
```

## Event Consumption Examples

### Setting Up Event Handlers

```go
// Create custom event handlers
orderCreatedHandler := func(event *event.OrderCreatedEvent) error {
    fmt.Printf("Processing order created: %s\n", event.GetOrder().GetID().GetValue().String())
    // Process the event
    return nil
}

orderPaidHandler := func(event *event.OrderPaidEvent) error {
    fmt.Printf("Processing order paid: %s\n", event.GetOrder().GetID().GetValue().String())
    // Process the event
    return nil
}

orderCancelledHandler := func(event *event.OrderCancelledEvent) error {
    fmt.Printf("Processing order cancelled: %s\n", event.GetOrder().GetID().GetValue().String())
    // Process the event
    return nil
}

// Create messaging module with custom handlers
messagingModule, err := messaging.NewMessagingModule()
if err != nil {
    log.Fatalf("Failed to initialize: %v", err)
}
defer messagingModule.Close()

// Subscribe and start consuming
err = messagingModule.SubscribeToOrderEvents()
if err != nil {
    log.Fatalf("Failed to subscribe: %v", err)
}

messagingModule.StartConsumingOrderEvents()
```

## Configuration

### Environment Variables

The messaging module uses the same environment variables as the Kafka module:

```bash
# Kafka Configuration
KAFKA_ENV=dev
KAFKA_BOOTSTRAP_SERVERS=localhost:9092
KAFKA_NUM_OF_PARTITIONS=3
KAFKA_REPLICATION_FACTOR=1

# Security (optional)
KAFKA_SECURITY_PROTOCOL=SASL_SSL
KAFKA_SASL_MECHANISM=PLAIN
KAFKA_SASL_USERNAME=kafka-user
KAFKA_SASL_PASSWORD=kafka-password

# Producer Configuration
KAFKA_PRODUCER_COMPRESSION_TYPE=snappy
KAFKA_PRODUCER_ACKS=all
KAFKA_PRODUCER_BATCH_SIZE=16384

# Consumer Configuration
KAFKA_CONSUMER_GROUP_ID=order-service-group
KAFKA_CONSUMER_AUTO_OFFSET_RESET=earliest
```

### Topic Configuration

The messaging module uses the following topics by default:

- `order_created` - Order creation events
- `order_paid` - Order payment events
- `order_cancelled` - Order cancellation events

## Testing

### Unit Testing

```go
func TestMessagingModule(t *testing.T) {
    messagingModule, err := messaging.NewMessagingModule()
    assert.NoError(t, err)
    defer messagingModule.Close()

    // Test health check
    err = messagingModule.Health()
    assert.NoError(t, err)

    // Test event publishing
    publisher := messagingModule.GetOrderEventPublisher()
    customData := map[string]interface{}{
        "order_id": "test-order",
        "event_type": "test",
    }

    err = publisher.PublishWithCustomData("test-topic", "test-order", customData)
    assert.NoError(t, err)
}
```

### Integration Testing

```go
func TestOrderEventFlow(t *testing.T) {
    messagingModule, err := messaging.NewMessagingModule()
    assert.NoError(t, err)
    defer messagingModule.Close()

    // Subscribe to events
    err = messagingModule.SubscribeToOrderEvents()
    assert.NoError(t, err)

    // Start consuming
    messagingModule.StartConsumingOrderEvents()

    // Publish test event
    publisher := messagingModule.GetOrderEventPublisher()
    testData := map[string]interface{}{
        "order_id": "test-order-123",
        "event_type": "order_created",
    }

    err = publisher.PublishWithCustomData("order_created", "test-order-123", testData)
    assert.NoError(t, err)

    // Wait for processing
    time.Sleep(1 * time.Second)
}
```

## Best Practices

1. **Initialize once**: Create the messaging module once and reuse it
2. **Always close**: Use `defer messagingModule.Close()` to ensure proper cleanup
3. **Handle errors**: Always check for errors when publishing events
4. **Use health checks**: Implement health checks in your application
5. **Monitor topics**: Keep track of message production and consumption rates
6. **Handle failures**: Implement retry logic for failed message publishing

## Troubleshooting

### Common Issues

1. **Connection failures**: Check Kafka bootstrap servers and network connectivity
2. **Authentication errors**: Verify SASL credentials and security protocol
3. **Topic not found**: Ensure topics exist in Kafka cluster
4. **Consumer group issues**: Check consumer group configuration

### Debug Mode

Enable debug logging by setting environment variable:
```bash
KAFKA_DEBUG=true
```

## Migration from Old Implementation

### Before (Old Implementation)
```go
// Complex setup with multiple components
producer, err := kafkaproducer.NewProducer(baseConfig, producerConfig)
if err != nil {
    log.Fatalf("Failed to create producer: %v", err)
}
defer producer.Close()

consumer, err := kafkaconsumer.NewConsumer(baseConfig, consumerConfig)
if err != nil {
    log.Fatalf("Failed to create consumer: %v", err)
}
defer consumer.Close()

// Manual event publishing
message := kafkamodel.NewMessage("order_created", key, value)
err = producer.ProduceMessage(message)
```

### After (New Implementation)
```go
// Simple initialization
messagingModule, err := messaging.NewMessagingModule()
if err != nil {
    log.Fatalf("Failed to initialize: %v", err)
}
defer messagingModule.Close()

// Automatic event publishing
publisher := messagingModule.GetOrderEventPublisher()
err = publisher.PublishOrderCreated(orderCreatedEvent)
```

## Benefits

### 1. Simplified Development
- **50% less code** for basic messaging setup
- **Automatic configuration** with sensible defaults
- **Unified interface** for all messaging operations

### 2. Improved Reliability
- **Built-in error handling** with comprehensive error messages
- **Health checks** for monitoring
- **Graceful shutdown** with proper cleanup

### 3. Better Maintainability
- **Consistent patterns** across the codebase
- **Dependency injection** for testability
- **Clear separation** of concerns

### 4. Production Ready
- **Protocol Buffers support** for efficient serialization
- **Environment-aware** configuration
- **Monitoring and observability** support 