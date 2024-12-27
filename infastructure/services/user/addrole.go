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
	Role        enums.Role `json:"role" validate:"required,gt=0"`
	Description string     `json:"description" validate:"required"`
}

type CreateRoleResponse struct {
	RoleId uuid.UUID `json:"id"`
}

func (userService UserService) AddRole(request CreateRoleRequest) (CreateRoleResponse, models.ErrorResponse) {
	roleId, _ := uuid.NewV7()
	rolename, err := request.Role.GetRoleName()
	if err != nil {
		return CreateRoleResponse{}, models.ErrorResponse{StatusCode: http.StatusBadRequest, Error: err}
	}
	userRole := entities.UserRole{
		ID:          roleId,
		Name:        rolename,
		Description: request.Description,
		Role:        request.Role,
	}

	rowsAffected := userService.DB.Where(entities.UserRole{Role: enums.Role(request.Role)}).FirstOrCreate(&userRole).RowsAffected
	if rowsAffected < 1 {
		err := fmt.Errorf("role already exist")
		return CreateRoleResponse{}, models.ErrorResponse{StatusCode: http.StatusBadRequest, Error: err}
	}

	return CreateRoleResponse{RoleId: userRole.ID}, models.ErrorResponse{}
}
