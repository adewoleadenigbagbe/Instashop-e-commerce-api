package routes

import (
	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/controllers"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/core"
	"github.com/labstack/echo/v4"
)

func orderRoutes(app *core.BaseApp, group *echo.Group) {
	orderController := controllers.OrderController{App: app}
	//group.POST("orders", orderController.CreateOrderHandler, middlewares.AuthorizeUser)
	//group.POST("orders", orderController.CreateOrderHandler,)

	group.POST("orders", orderController.CreateOrderHandler)
	group.GET("orders", orderController.GetOrdersHandler)
	group.PUT("orders/:id", orderController.UpdateOrderStatusHandler)
	group.PUT("orders/:id/cancel", orderController.CancelOrderHandler)
}
