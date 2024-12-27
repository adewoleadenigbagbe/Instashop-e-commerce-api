package services

import (
	"net/http"

	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/entities"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	EmailRegex       = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$"
	PhoneNumberRegex = "\\+[1-9]{1}[0-9]{0,2}-[2-9]{1}[0-9]{2}-[2-9]{1}[0-9]{2}-[0-9]{4}$"
)

type CreateUserRequest struct {
	FirstName   string `json:"firstName" validate:"required"`
	LastName    string `json:"lastName" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required,regex=\\+[1-9]{1}[0-9]{0,2}-[2-9]{1}[0-9]{2}-[2-9]{1}[0-9]{2}-[0-9]{4}$"`
	RoleId      string `json:"roleId" validate:"required,uuid"`
}

type CreateUserResponse struct {
	UserId uuid.UUID `json:"id"`
}

func (authService AuthService) RegisterUser(request CreateUserRequest) (CreateUserResponse, models.ErrorResponse) {
	var err error
	userId, _ := uuid.NewV7()

	user := entities.User{
		ID:           userId,
		FirstName:    request.FirstName,
		LastName:     request.LastName,
		RoleId:       request.RoleId,
		Email:        request.Email,
		PhoneNumber:  request.PhoneNumber,
		Password:     request.Password,
		IsDeprecated: false,
	}

	err = authService.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			// return any error will rollback
			return err
		}
		return nil
	})

	if err != nil {
		return CreateUserResponse{}, models.ErrorResponse{Error: err, StatusCode: http.StatusBadRequest}
	}

	return CreateUserResponse{UserId: user.ID}, models.ErrorResponse{}
}
