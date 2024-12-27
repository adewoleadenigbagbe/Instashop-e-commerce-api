package routes

import (
	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/controllers"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/core"
	middlewares "github.com/adewoleadenigbagbe/instashop-e-commerce/apis/middleware"
	"github.com/labstack/echo/v4"
)

func categoryRoutes(app *core.BaseApp, group *echo.Group) {
	categoryController := controllers.CategoryController{App: app}

	group.POST("categories", categoryController.CreateCategoryHandler, middlewares.AuthorizeAdmin)
}
