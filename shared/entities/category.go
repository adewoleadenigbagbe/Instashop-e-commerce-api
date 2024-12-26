package entities

import (
	"time"

	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/enums"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID           uuid.UUID          `gorm:"column:Id"`
	Name         string             `gorm:"column:Name"`
	Description  string             `gorm:"column:Description"`
	Type         enums.CategoryType `gorm:"column:Type"`
	CreatedOn    time.Time          `gorm:"column:CreatedOn;autoCreateTime"`
	ModifiedOn   time.Time          `gorm:"column:ModifiedOn;autoUpdateTime"`
	IsDeprecated bool               `gorm:"column:IsDeprecated"`
}

func (Category) TableName() string {
	return "Categories"
}

func (category *Category) BeforeCreate(tx *gorm.DB) (err error) {
	if category.ID == uuid.Nil {
		id, err := uuid.NewV7()
		if err == nil {
			category.ID = id
		}
	}
	category.IsDeprecated = false
	return
}
