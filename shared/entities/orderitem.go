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
	CreatedOn    time.Time `gorm:"column:CreatedOn;autoCreateTime"`
	ModifiedOn   time.Time `gorm:"column:ModifiedOn;autoUpdateTime"`
	IsDeprecated bool      `gorm:"column:IsDeprecated"`
}

func (OrderItem) TableName() string {
	return "OrderItems"
}

func (orderItem *OrderItem) CalculateTotal() {
	orderItem.TotalPrice = orderItem.UnitPrice * float64(orderItem.Quantity)
}

func (orderItem *OrderItem) BeforeCreate(tx *gorm.DB) (err error) {
	if orderItem.ID == uuid.Nil {
		id, err := uuid.NewV7()
		if err == nil {
			orderItem.ID = id
		}
	}
	orderItem.IsDeprecated = false
	return
}
