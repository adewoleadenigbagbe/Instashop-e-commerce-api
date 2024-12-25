package routes

import (
	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/controllers"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/core"
	middlewares "github.com/adewoleadenigbagbe/instashop-e-commerce/apis/middleware"
	"github.com/labstack/echo/v4"
)

func cinemaRoutes(app *core.BaseApp, group *echo.Group) {
	cinemaController := controllers.CinemaController{App: app}
	group.POST("cinemas", cinemaController.CreateCinemaHandler, middlewares.AuthorizeAdmin)
	group.POST("cinemas/:id/cinemahall", cinemaController.CreateCinemaHallHandler, middlewares.AuthorizeAdmin)
	group.POST("cinemas/:id/cinemahall/:cinemahallId/seat", cinemaController.CreateCinemaSeatHandler, middlewares.AuthorizeAdmin)
	group.POST("cinemas/:id/rate", cinemaController.AddCinemaRateHandler, middlewares.AuthorizeAdmin)
}
