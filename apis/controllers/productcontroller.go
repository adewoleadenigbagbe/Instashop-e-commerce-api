package controllers

import (
	"net/http"
	"reflect"

	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/core"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/infastructure/services"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

type ProductController struct {
	App *core.BaseApp
}

func (controller ProductController) CreateProductHandler(productContext echo.Context) error {
	var err error
	request := new(services.CreateProductRequest)
	err = productContext.Bind(request)
	if err != nil {
		return productContext.JSON(http.StatusBadRequest, err.Error())
	}

	dataResp, errResp := controller.App.ProductService.CreateProduct(*request)
	if !reflect.ValueOf(errResp).IsZero() {
		errs := lo.Map(errResp.Errors, func(er error, index int) string {
			return er.Error()
		})
		return productContext.JSON(errResp.StatusCode, errs)
	}
	return productContext.JSON(http.StatusOK, dataResp)
}
