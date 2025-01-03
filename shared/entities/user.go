package entities

import (
	"errors"
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID `gorm:"column:Id"`
	FirstName    string    `gorm:"column:FirstName"`
	LastName     string    `gorm:"column:LastName"`
	RoleId       string    `gorm:"column:RoleId"`
	Email        string    `gorm:"column:Email"`
	PhoneNumber  string    `gorm:"column:PhoneNumber"`
	Password     string    `gorm:"column:Password"`
	IsDeprecated bool      `gorm:"column:IsDeprecated"`
	CreatedOn    time.Time `gorm:"column:CreatedOn;autoCreateTime"`
	ModifiedOn   time.Time `gorm:"column:ModifiedOn;autoUpdateTime"`
	UserRole     UserRole  `gorm:"foreignKey:RoleId;references:Id;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

func (User) TableName() string {
	return "Users"
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	if user.ID == uuid.Nil {
		id, err := uuid.NewV7()
		if err != nil {
			user.ID = id
		}
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.FirstName = html.EscapeString(strings.TrimSpace(user.FirstName))
	user.LastName = html.EscapeString(strings.TrimSpace(user.LastName))
	user.IsDeprecated = false
	return
}

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	var result *gorm.DB
	result = tx.Where("Email = ?", user.Email).
		Where("IsDeprecated = ?", false).
		First(user)

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return fmt.Errorf("email already in use")
	}

	result = tx.Where("Password = ?", user.Password).
		Where("IsDeprecated = ?", false).
		First(user)

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return fmt.Errorf("password already taken")
	}

	return nil
}
