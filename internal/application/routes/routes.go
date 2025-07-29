package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/leninner/order-service/internal/application/rest"
	customerAdapter "github.com/leninner/order-service/internal/dataaccess/customer/adapter"
	orderAdapter "github.com/leninner/order-service/internal/dataaccess/order/adapter"
	restaurantAdapter "github.com/leninner/order-service/internal/dataaccess/restaurant/adapter"
	applicationservice "github.com/leninner/order-service/internal/domain/application-service"
	"github.com/leninner/order-service/internal/domain/application-service/mapper"
	"github.com/leninner/order-service/internal/domain/core"
	"github.com/leninner/order-service/internal/messaging/publisher"
	"github.com/leninner/shared/config"
	"github.com/leninner/shared/exception"
	"github.com/leninner/shared/logger"
	"github.com/leninner/shared/middleware"
)

func Routes(app *config.Application) http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(exception.NotFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(exception.MethodNotAllowedResponse)

	loggerInstance, err := logger.NewDevelopmentLogger("order-service")
	if err != nil {
		panic(err)
	}

	orderRepository := orderAdapter.NewOrderRepositoryImpl(app.DataSource)
	customerRepository := customerAdapter.NewCustomerRepositoryImpl(app.DataSource)
	restaurantRepository := restaurantAdapter.NewRestaurantRepositoryImpl(app.DataSource)

	orderDataMapper := mapper.NewOrderDataMapper()
	orderDomainService := core.NewOrderDomainServiceImpl(loggerInstance)
	orderCreateHelper := applicationservice.NewOrderCreateHelper(
		orderRepository,
		customerRepository,
		restaurantRepository,
		*orderDataMapper,
		orderDomainService,
		loggerInstance,
	)
	orderCreatedPaymentRequestMessagePublisher := publisher.NewOrderCreatedPaymentRequestMessagePublisherImpl()

	orderCreateCommandHandler := applicationservice.NewOrderCreateCommandHandler(
		orderCreateHelper,
		orderDataMapper,
		orderCreatedPaymentRequestMessagePublisher,
	)
	orderTrackCommandHandler := applicationservice.NewOrderTrackCommandHandler(
		orderDataMapper,
		orderRepository,
	)
	orderApplicationService := applicationservice.NewOrderApplicationService(
		*orderCreateCommandHandler,
		*orderTrackCommandHandler,
	)

	orderController := rest.NewOrderController(orderApplicationService)

	router.HandlerFunc(http.MethodPost, "/v1/orders", orderController.CreateOrder)
	router.HandlerFunc(http.MethodGet, "/v1/orders/:orderTrackingId", orderController.TrackOrder)

	return middleware.RecoverPanic(router)
}
