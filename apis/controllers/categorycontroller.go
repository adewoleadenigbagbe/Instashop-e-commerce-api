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
