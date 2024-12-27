package controllers

import (
	"net/http"
	"reflect"

	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/core"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/utils"
	services "github.com/adewoleadenigbagbe/instashop-e-commerce/infastructure/services/category"
	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	App *core.BaseApp
}

// Category godoc
// @Summary     Create a category
// @Description    Create a new category for product
// @Tags         category
// @Accept       json
// @Produce      json
// @Param        CreateCategoryRequest  body  services.CreateCategoryRequest  true  "CreateCategoryRequest"
// @Success      200  {object}  services.CreateCategoryResponse
// @Failure      400  {object}  string
// @Failure      404  {object}  []string
// @Router       /api/v1/categories [post]

func (controller CategoryController) CreateCategoryHandler(orderContext echo.Context) error {
	request := new(services.CreateCategoryRequest)
	err := orderContext.Bind(request)
	if err != nil {
		return orderContext.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate request
	if err := orderContext.Validate(request); err != nil {
		errors := utils.GetValidationErrors(err)
		return orderContext.JSON(http.StatusBadRequest, errors)
	}

	dataResp, errResp := controller.App.CategoryService.CreateCategory(*request)
	if !reflect.ValueOf(errResp).IsZero() {
		return orderContext.JSON(errResp.StatusCode, errResp.Error.Error())
	}
	return orderContext.JSON(http.StatusCreated, dataResp)
}
