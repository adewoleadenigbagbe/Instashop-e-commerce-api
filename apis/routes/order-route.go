package routes

import (
	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/controllers"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/core"
	middlewares "github.com/adewoleadenigbagbe/instashop-e-commerce/apis/middleware"
	"github.com/labstack/echo/v4"
)

func orderRoutes(app *core.BaseApp, group *echo.Group) {
	orderController := controllers.OrderController{App: app}

	group.POST("orders", orderController.CreateOrderHandler, middlewares.AuthorizeUser)
	group.GET("orders", orderController.GetOrdersHandler, middlewares.AuthorizeUser)
	group.PUT("orders/:id", orderController.UpdateOrderStatusHandler, middlewares.AuthorizeAdmin)
	group.PUT("orders/:id/cancel", orderController.CancelOrderHandler, middlewares.AuthorizeUser)
}
