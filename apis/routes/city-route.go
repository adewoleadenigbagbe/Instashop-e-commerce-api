package routes

import (
	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/controllers"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/core"
	middlewares "github.com/adewoleadenigbagbe/instashop-e-commerce/apis/middleware"
	"github.com/labstack/echo/v4"
)

func cityRoutes(app *core.BaseApp, group *echo.Group) {
	cityController := controllers.CityController{App: app}
	group.GET("city/:id", cityController.GetCityByIdHandler, middlewares.AuthorizeAdmin)
}
