package core

import (
	authService "github.com/adewoleadenigbagbe/instashop-e-commerce/infastructure/services/auth"
	productService "github.com/adewoleadenigbagbe/instashop-e-commerce/infastructure/services/product"
	services "github.com/adewoleadenigbagbe/instashop-e-commerce/infastructure/services/product"
	userService "github.com/adewoleadenigbagbe/instashop-e-commerce/infastructure/services/user"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/configs"
	db "github.com/adewoleadenigbagbe/instashop-e-commerce/shared/database"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// BaseApp implements core.App and defines the structure of the whole application
type BaseApp struct {
	DB             *gorm.DB
	Echo           *echo.Echo
	ProductService productService.ProductService
	UserService    userService.UserService
	AuthService    authService.AuthService
}

func ConfigureApp() (*BaseApp, error) {
	var err error

	//create a database connection
	dbConfig, err := configs.CreateDbConfig()
	if err != nil {
		return nil, err
	}

	db, err := db.ConnectToDatabase(*dbConfig)
	if err != nil {
		return nil, err
	}

	app := &BaseApp{
		Echo:           echo.New(),
		DB:             db,
		ProductService: services.ProductService{DB: db},
	}

	return app, nil
}

func (app *BaseApp) Db() *gorm.DB {
	return app.DB
}

func (app *BaseApp) GetEcho() *echo.Echo {
	return app.Echo
}
