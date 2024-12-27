package routes

import (
	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/controllers"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/core"
	middlewares "github.com/adewoleadenigbagbe/instashop-e-commerce/apis/middleware"
	"github.com/labstack/echo/v4"
)

func productRoutes(app *core.BaseApp, group *echo.Group) {
	productController := controllers.ProductController{App: app}

	group.POST("products", productController.CreateProductHandler, middlewares.AuthorizeAdmin)
	group.GET("products/:id", productController.GetProductByIdHandler, middlewares.AuthorizeAdmin)
	group.GET("products", productController.GetProductsHandler, middlewares.AuthorizeAdmin)
	group.PUT("products/:id", productController.UpdateProductHandler, middlewares.AuthorizeAdmin)
	group.DELETE("products/:id", productController.DeleteProductHandler, middlewares.AuthorizeAdmin)

}
