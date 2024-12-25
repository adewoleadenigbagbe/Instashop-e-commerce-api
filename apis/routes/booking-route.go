package routes

import (
	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/controllers"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/core"
	middlewares "github.com/adewoleadenigbagbe/instashop-e-commerce/apis/middleware"
	"github.com/labstack/echo/v4"
)

func bookingRoutes(app *core.BaseApp, group *echo.Group) {
	bookingController := controllers.BookingController{App: app}
	group.POST("booking/book-show", bookingController.BookShowHandler)
	group.POST("booking/payment", bookingController.ChargeBookingHandler)
	group.GET("booking/generate-pdf", bookingController.GenerateInvoiceHandler, middlewares.AuthorizeUser)
}
