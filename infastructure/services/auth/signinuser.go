package services

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/entities"
	jwtauth "github.com/adewoleadenigbagbe/instashop-e-commerce/shared/helpers/utilities/auth"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SignInResponse struct {
	Token string `json:"access_token"`
}

func (authService AuthService) SignIn(request SignInRequest) (SignInResponse, models.ErrorResponse) {
	var err error

	var user entities.User
	err = authService.DB.Where("Email = ?", request.Email).
		Where("IsDeprecated = ?", false).
		Preload("UserRole").
		First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return SignInResponse{}, models.ErrorResponse{StatusCode: http.StatusBadRequest, Error: fmt.Errorf("email or password not found")}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return SignInResponse{}, models.ErrorResponse{StatusCode: http.StatusBadRequest, Error: fmt.Errorf("email or password not found")}
	}

	token, err := jwtauth.GenerateJWT(user)
	if err != nil {
		return SignInResponse{}, models.ErrorResponse{StatusCode: http.StatusBadRequest, Error: err}
	}

	return SignInResponse{Token: token}, models.ErrorResponse{}
}
