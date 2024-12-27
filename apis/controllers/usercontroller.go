package controllers

import (
	"net/http"
	"reflect"

	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/core"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/utils"
	services "github.com/adewoleadenigbagbe/instashop-e-commerce/infastructure/services/user"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	App *core.BaseApp
}

// CreateUserRole godoc
// @Summary      Create a new role
// @Description   Create a new role
// @Tags         userole
// @Accept       json
// @Produce      json
// @Param        CreateRoleRequest  body  services.CreateRoleRequest  true  "CreateRoleRequest"
// @Success      200  {object}  services.CreateRoleResponse
// @Failure      400  {object}  string
// @Failure      404  {object}  []string
// @Router       /api/v1/user/add-role [post]
func (userController UserController) AddRoleHandler(userContext echo.Context) error {
	var err error
	request := new(services.CreateRoleRequest)
	err = userContext.Bind(request)
	if err != nil {
		return userContext.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate request
	if err := userContext.Validate(request); err != nil {
		errors := utils.GetValidationErrors(err)
		return userContext.JSON(http.StatusBadRequest, errors)
	}

	dataResp, errResp := userController.App.UserService.AddRole(*request)
	if !reflect.ValueOf(errResp).IsZero() {
		return userContext.JSON(errResp.StatusCode, errResp.Error.Error())
	}

	return userContext.JSON(http.StatusCreated, dataResp)
}
