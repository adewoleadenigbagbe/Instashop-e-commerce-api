package entities

import (
	"time"

	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/enums"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID           uuid.UUID           `gorm:"column:Id"`
	Name         string              `gorm:"column:Name"`
	Description  string              `gorm:"column:Description"`
	Price        float64             `gorm:"column:Price"`
	Stock        int                 `gorm:"column:Stock"`
	CategoryID   uuid.UUID           `gorm:"column:CategoryId"`
	ImageURL     string              `gorm:"column:ImageUrl"`
	Status       enums.ProductStatus `gorm:"column:Status"`
	CreatedOn    time.Time           `gorm:"column:CreatedOn;autoCreateTime"`
	ModifiedOn   time.Time           `gorm:"column:ModifiedOn;autoUpdateTime"`
	IsDeprecated bool                `gorm:"column:IsDeprecated"`
}

func (Product) TableName() string {
	return "Products"
}

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	if product.ID == uuid.Nil {
		id, err := uuid.NewV7()
		if err == nil {
			product.ID = id
		}
	}
	product.IsDeprecated = false
	return
}
