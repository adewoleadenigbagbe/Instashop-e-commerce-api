package entities

import (
	"time"

	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/enums"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type UserRole struct {
	ID          uuid.UUID  `gorm:"column:Id"`
	Name        string     `gorm:"column:Name"`
	Role        enums.Role `gorm:"column:Role"`
	Description string     `gorm:"column:Description"`
	CreatedOn   time.Time  `gorm:"column:CreatedOn;autoCreateTime"`
	ModifiedOn  time.Time  `gorm:"column:ModifiedOn;autoUpdateTime"`
}

func (userRole *UserRole) BeforeCreate(tx *gorm.DB) (err error) {
	if userRole.ID == uuid.Nil {
		id, err := uuid.NewV7()
		if err == nil {
			userRole.ID = id
		}
	}
	return
}

func (UserRole) TableName() string {
	return "UserRoles"
}
