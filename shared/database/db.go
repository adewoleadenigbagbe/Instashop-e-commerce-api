package db

import (
	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/configs"
	"github.com/fatih/color"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDatabase(config configs.DatabaseConfig) (*gorm.DB, error) {
	var err error
	dsn := config.GetDsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	color.Blue("Connected to the Database")

	return db, nil
}
