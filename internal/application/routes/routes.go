package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/leninner/order-service/internal/application/rest"
	"github.com/leninner/shared/config"
	"github.com/leninner/shared/exception"
	"github.com/leninner/shared/middleware"
)

func Routes(app *config.Application) http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(exception.NotFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(exception.MethodNotAllowedResponse)

	orderController := rest.NewOrderController(nil)

	router.HandlerFunc(http.MethodPost, "/v1/orders", orderController.CreateOrder)

	return middleware.RecoverPanic(router)
}