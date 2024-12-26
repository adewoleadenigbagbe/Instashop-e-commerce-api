package services

import (
	"fmt"
	"net/http"

	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/entities"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/enums"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

type CreateRoleRequest struct {
	Name        string `json:"name"`
	Role        int    `json:"role"`
	Description string `json:"description"`
}

type CreateRoleResponse struct {
	UserRoleId uuid.UUID `json:"userRoleId"`
}

func (userService UserService) AddRole(request CreateRoleRequest) (CreateRoleResponse, models.ErrorResponse) {

	roleId, _ := uuid.NewV7()
	userRole := entities.UserRole{
		ID:          roleId,
		Name:        request.Name,
		Description: request.Description,
		Role:        enums.Role(request.Role),
	}

	rowsAffected := userService.DB.Where(entities.UserRole{Role: enums.Role(request.Role)}).FirstOrCreate(&userRole).RowsAffected
	if rowsAffected < 1 {
		err := fmt.Errorf("role already exist")
		return CreateRoleResponse{}, models.ErrorResponse{StatusCode: http.StatusBadRequest, Error: err}
	}

	return CreateRoleResponse{UserRoleId: userRole.ID}, models.ErrorResponse{}
}

// func validateRole(request CreateRoleRequest) []error {
// 	vErrors := []error{}
// 	if request.Name == "" {
// 		vErrors = append(vErrors, fmt.Errorf(ErrRequiredField, "name"))
// 	}

// 	if request.Description == "" {
// 		vErrors = append(vErrors, fmt.Errorf(ErrRequiredField, "description"))
// 	}

// 	if request.Role <= 0 {
// 		vErrors = append(vErrors, fmt.Errorf(ErrRequiredField, "role"))
// 	}

// 	return vErrors
// }
