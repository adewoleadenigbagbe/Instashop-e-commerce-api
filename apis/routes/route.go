package routes

import (
	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/core"
)

// RegisterRoutes - Configuration for all incoming routes
func RegisterAllRoutes(app *core.BaseApp) {
	group := app.Echo.Group("/api/v1/")
	authRoutes(app, group)
	productRoutes(app, group)
	userRoutes(app, group)
	orderRoutes(app, group)
	categoryRoutes(app, group)
}
