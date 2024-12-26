package routes

import (
	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/controllers"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/core"
	"github.com/labstack/echo/v4"
)

func productRoutes(app *core.BaseApp, group *echo.Group) {
	productController := controllers.ProductController{App: app}

	//group.POST("products", productController.CreateProductHandler, middlewares.AuthorizeAdmin)
	// group.GET("products/:id", productController.GetProductByIdHandler, middlewares.AuthorizeUser)
	// group.GET("products", productController.GetProductsHandler, middlewares.AuthorizeUser)
	// group.PUT("products/:id", productController.UpdateProductHandler, middlewares.AuthorizeAdmin)
	// group.DELETE("products/:id", productController.DeleteProductHandler, middlewares.AuthorizeAdmin)

	group.POST("products", productController.CreateProductHandler)
	group.GET("products/:id", productController.GetProductByIdHandler)
	group.GET("products", productController.GetProductsHandler)
	group.PUT("products/:id", productController.UpdateProductHandler)
	group.DELETE("products/:id", productController.DeleteProductHandler)

}
