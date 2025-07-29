package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/leninner/order-service/internal/application"
	"github.com/leninner/order-service/internal/application/rest"
	"github.com/leninner/shared/exception"
	"github.com/leninner/shared/middleware"
)

func Routes(container *application.OrderServiceContainer) http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(exception.NotFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(exception.MethodNotAllowedResponse)

	orderController := rest.NewOrderController(container.Services.OrderApplication)

	router.HandlerFunc(http.MethodPost, "/v1/orders", orderController.CreateOrder)
	router.HandlerFunc(http.MethodGet, "/v1/orders/:orderTrackingId", orderController.TrackOrder)

	return middleware.RecoverPanic(router)
} 