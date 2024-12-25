package routes

import (
	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/controllers"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/core"
	middlewares "github.com/adewoleadenigbagbe/instashop-e-commerce/apis/middleware"
	"github.com/labstack/echo/v4"
)

func showRoutes(app *core.BaseApp, group *echo.Group) {
	showController := controllers.ShowController{App: app}
	group.POST("shows", showController.CreateShowHandler, middlewares.AuthorizeAdmin)
	group.POST("shows", showController.CreateShowHandler)
	group.GET("shows/user-location", showController.GetShowsByUserLocationHandler, middlewares.AuthorizeUser)
	group.GET("shows/:id/available-seat", showController.GetAvailableShowSeatHandler, middlewares.AuthorizeUser)
}
