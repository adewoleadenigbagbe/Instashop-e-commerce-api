// package services

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/database/entities"
// 	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/enums"
// 	sequentialguid "github.com/adewoleadenigbagbe/instashop-e-commerce/shared/helpers"
// 	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/models"
// )

// type CreateRoleRequest struct {
// 	Name        string `json:"name"`
// 	Role        int    `json:"role"`
// 	Description string `json:"description"`
// }

// type CreateRoleResponse struct {
// 	UserRoleId string `json:"userRoleId"`
// }

// func (userService UserService) AddRole(request CreateRoleRequest) (CreateRoleResponse, models.ErrorResponse) {
// 	requiredFieldErrors := validateRole(request)
// 	if len(requiredFieldErrors) > 0 {
// 		return CreateRoleResponse{}, models.ErrorResponse{StatusCode: http.StatusBadRequest, Errors: requiredFieldErrors}
// 	}

// 	userRole := entities.UserRole{
// 		Id:          sequentialguid.New().String(),
// 		Name:        request.Name,
// 		Description: request.Description,
// 		Role:        enums.Role(request.Role),
// 	}

// 	rowsAffected := userService.DB.Where(entities.UserRole{Role: enums.Role(request.Role)}).FirstOrCreate(&userRole).RowsAffected

// 	if rowsAffected < 1 {
// 		err := fmt.Errorf("role already exist")
// 		return CreateRoleResponse{}, models.ErrorResponse{StatusCode: http.StatusBadRequest, Errors: []error{err}}
// 	}

// 	return CreateRoleResponse{UserRoleId: userRole.Id}, models.ErrorResponse{}
// }

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