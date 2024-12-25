package routes

import (
	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/controllers"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/core"
	"github.com/labstack/echo/v4"
)

func movieRoutes(app *core.BaseApp, group *echo.Group) {
	movieController := controllers.MovieController{App: app}
	group.GET("movies", movieController.GetMoviesHandler)
	group.GET("movies/:id", movieController.GetMovieByIdHandler)
	group.GET("movies/search", movieController.SearchMovieHandler)
}
