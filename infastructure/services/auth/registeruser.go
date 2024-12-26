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
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
	RoleId      string `json:"roleId"`
}

type CreateUserResponse struct {
	UserId uuid.UUID `json:"userId"`
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

// func validateUser(request CreateUserRequest) []error {
// 	var validationErrors []error
// 	if request.FirstName == "" {
// 		validationErrors = append(validationErrors, fmt.Errorf(ErrRequiredField, "firstName"))
// 	}

// 	if request.LastName == "" {
// 		validationErrors = append(validationErrors, fmt.Errorf(ErrRequiredField, "lastName"))
// 	}

// 	if request.Password == "" {
// 		validationErrors = append(validationErrors, fmt.Errorf(ErrRequiredField, "password"))
// 	}

// 	if request.Address == "" {
// 		validationErrors = append(validationErrors, fmt.Errorf(ErrRequiredField, "address"))
// 	}

// 	if len(request.CityId) == 0 || len(request.CityId) < 36 {
// 		validationErrors = append(validationErrors, fmt.Errorf(ErrRequiredUUIDField, "cityId"))
// 	}

// 	if request.CityId == utilities.DEFAULT_UUID {
// 		validationErrors = append(validationErrors, fmt.Errorf(ErrInvalidUUID, "cityId"))
// 	}

// 	isEmailValid, _ := regexp.MatchString(EmailRegex, request.Email)
// 	if !isEmailValid {
// 		validationErrors = append(validationErrors, fmt.Errorf(ErrInValidField, "email"))
// 	}

// 	isPhoneValid, _ := regexp.MatchString(PhoneNumberRegex, request.PhoneNumber)

// 	if !isPhoneValid {
// 		validationErrors = append(validationErrors, fmt.Errorf(ErrInValidField, "phone number"))
// 	}

// 	return validationErrors
// }
