package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderItem struct {
	ID           uuid.UUID `gorm:"column:Id"`
	OrderID      uuid.UUID `gorm:"column:OrderId"`
	ProductID    uuid.UUID `gorm:"column:ProductId"`
	Quantity     int       `gorm:"column:Quantity"`
	UnitPrice    float64   `gorm:"column:UnitPrice"`
	TotalPrice   float64   `gorm:"column:TotalPrice"`
	CreatedOn    time.Time `gorm:"column:CreatedOn"`
	ModifiedOn   time.Time `gorm:"column:ModifiedOn"`
	IsDeprecated bool      `gorm:"column:IsDeprecated"`
}

func (OrderItem) TableName() string {
	return "OrderItems"
}

func (orderItem *OrderItem) BeforeCreate(tx *gorm.DB) (err error) {
	if orderItem.ID == uuid.Nil {
		id, err := uuid.NewV7()
		if err != nil {
			orderItem.ID = id
		}
	}
	orderItem.IsDeprecated = false
	return
}