package application

import (
	customerAdapter "github.com/leninner/order-service/internal/dataaccess/customer/adapter"
	orderAdapter "github.com/leninner/order-service/internal/dataaccess/order/adapter"
	restaurantAdapter "github.com/leninner/order-service/internal/dataaccess/restaurant/adapter"
	applicationservice "github.com/leninner/order-service/internal/domain/application-service"
	appConfig "github.com/leninner/order-service/internal/domain/application-service/config"
	appMapper "github.com/leninner/order-service/internal/domain/application-service/mapper"
	"github.com/leninner/order-service/internal/domain/application-service/ports/input/service"
	paymentPublisher "github.com/leninner/order-service/internal/domain/application-service/ports/output/message/publisher/payment"
	restaurantApprovalPublisher "github.com/leninner/order-service/internal/domain/application-service/ports/output/message/publisher/restaurantapproval"
	"github.com/leninner/order-service/internal/domain/application-service/ports/output/repository"
	"github.com/leninner/order-service/internal/domain/core"
	messagingMapper "github.com/leninner/order-service/internal/messaging/mapper"
	messagingKafka "github.com/leninner/order-service/internal/messaging/publisher/kafka"
	sharedDI "github.com/leninner/shared/di"
)

type OrderServiceContainer struct {
	*sharedDI.SharedContainer
	Repositories *Repositories
	Services     *Services
	Publishers   *Publishers
	Configs      *Configs
	Mappers      *Mappers
}

type Repositories struct {
	Order      repository.OrderRepository
	Customer   repository.CustomerRepository
	Restaurant repository.RestaurantRepository
}

type Services struct {
	OrderApplication service.OrderApplicationService
	OrderDomain      core.OrderDomainService
}

type Publishers struct {
	OrderCreatedPaymentRequest    paymentPublisher.OrderCreatedPaymentRequestMessagePublisher
	OrderCancelledPaymentRequest  paymentPublisher.OrderCancelledPaymentRequestMessagePublisher
	OrderPaidRestaurantApproval   restaurantApprovalPublisher.OrderPaidRestaurantApprovalRequestMessagePublisher
}

type Configs struct {
	OrderService *appConfig.OrderServiceConfigData
}

type Mappers struct {
	OrderData        *appMapper.OrderDataMapper
	OrderMessaging   *messagingMapper.OrderMessagingDataMapper
}

func NewOrderServiceContainer(sharedContainer *sharedDI.SharedContainer) *OrderServiceContainer {
	dataSource := sharedContainer.GetDatabase()
	slogLogger := sharedContainer.GetLogger()
	sharedConfig := sharedContainer.GetConfig()
	
	// Create logger adapter for domain services
	loggerInstance := CreateLoggerAdapter(slogLogger)
	
	configs := &Configs{
		OrderService: appConfig.NewOrderServiceConfigDataFromShared(sharedConfig),
	}

	mappers := &Mappers{
		OrderData:      appMapper.NewOrderDataMapper(),
		OrderMessaging: messagingMapper.NewOrderMessagingDataMapper(),
	}

	repositories := &Repositories{
		Order:      orderAdapter.NewOrderRepositoryImpl(dataSource),
		Customer:   customerAdapter.NewCustomerRepositoryImpl(dataSource),
		Restaurant: restaurantAdapter.NewRestaurantRepositoryImpl(dataSource),
	}

	services := &Services{
		OrderDomain: core.NewOrderDomainServiceImpl(loggerInstance),
	}

	kafkaModule := sharedContainer.GetKafka()

	publishers := &Publishers{
		OrderCreatedPaymentRequest: messagingKafka.NewOrderCreatedPaymentRequestKafkaMessagePublisher(
			kafkaModule,
			mappers.OrderMessaging,
			configs.OrderService,
		),
		OrderCancelledPaymentRequest: messagingKafka.NewOrderCancelledPaymentRequestKafkaMessagePublisher(
			kafkaModule,
			mappers.OrderMessaging,
			configs.OrderService,
		),
		OrderPaidRestaurantApproval: messagingKafka.NewOrderPaidRestaurantApprovalRequestKafkaMessagePublisher(
			kafkaModule,
			mappers.OrderMessaging,
			configs.OrderService,
		),
	}

	orderCreateHelper := applicationservice.NewOrderCreateHelper(
		repositories.Order,
		repositories.Customer,
		repositories.Restaurant,
		*mappers.OrderData,
		services.OrderDomain,
		loggerInstance,
	)

	orderCreateCommandHandler := applicationservice.NewOrderCreateCommandHandler(
		orderCreateHelper,
		mappers.OrderData,
		publishers.OrderCreatedPaymentRequest,
	)

	orderTrackCommandHandler := applicationservice.NewOrderTrackCommandHandler(
		mappers.OrderData,
		repositories.Order,
	)

	services.OrderApplication = applicationservice.NewOrderApplicationService(
		*orderCreateCommandHandler,
		*orderTrackCommandHandler,
	)

	return &OrderServiceContainer{
		SharedContainer: sharedContainer,
		Repositories:    repositories,
		Services:        services,
		Publishers:      publishers,
		Configs:         configs,
		Mappers:         mappers,
	}
} 