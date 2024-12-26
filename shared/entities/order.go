package entities

import (
	"time"

	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/enums"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID           uuid.UUID         `gorm:"column:Id"`
	UserID       uuid.UUID         `gorm:"column:UserId"`
	Status       enums.OrderStatus `gorm:"column:Status"`
	TotalAmount  float64           `gorm:"column:TotalAmount"`
	TaxAmount    float64           `gorm:"column:TaxAmount"`
	CreatedOn    time.Time         `gorm:"column:CreatedOn;autoCreateTime"`
	ModifiedOn   time.Time         `gorm:"column:ModifiedOn;autoUpdateTime"`
	IsDeprecated bool              `gorm:"column:IsDeprecated"`
}

func (Order) TableName() string {
	return "Orders"
}

func (order *Order) BeforeCreate(tx *gorm.DB) (err error) {
	if order.ID == uuid.Nil {
		id, err := uuid.NewV7()
		if err == nil {
			order.ID = id
		}
	}
	order.IsDeprecated = false
	return
}
