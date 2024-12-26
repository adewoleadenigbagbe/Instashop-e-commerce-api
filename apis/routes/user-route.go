package routes

import (
	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/controllers"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/core"
	middlewares "github.com/adewoleadenigbagbe/instashop-e-commerce/apis/middleware"
	"github.com/labstack/echo/v4"
)

func userRoutes(app *core.BaseApp, group *echo.Group) {
	userController := controllers.UserController{App: app}
	group.POST("user/add-role", userController.AddRoleHandler, middlewares.AuthorizeAdmin)
}
