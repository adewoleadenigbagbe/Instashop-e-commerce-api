package controllers

import (
	"net/http"
	"reflect"

	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/core"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/utils"
	services "github.com/adewoleadenigbagbe/instashop-e-commerce/infastructure/services/product"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	App *core.BaseApp
}

func (controller ProductController) CreateProductHandler(productContext echo.Context) error {
	request := new(services.CreateProductRequest)
	err := productContext.Bind(request)
	if err != nil {
		return productContext.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate request
	if err := productContext.Validate(request); err != nil {
		errors := utils.GetValidationErrors(err)
		return productContext.JSON(http.StatusBadRequest, errors)
	}

	dataResp, errResp := controller.App.ProductService.CreateProduct(*request)
	if !reflect.ValueOf(errResp).IsZero() {
		return productContext.JSON(errResp.StatusCode, errResp.Error.Error())
	}
	return productContext.JSON(http.StatusCreated, dataResp)
}

func (controller ProductController) GetProductByIdHandler(productContext echo.Context) error {
	request := new(services.GetProductRequestById)
	err := productContext.Bind(request)
	if err != nil {
		return productContext.JSON(http.StatusBadRequest, err.Error())
	}

	dataResp, errResp := controller.App.ProductService.GetProductByID(*request)
	if !reflect.ValueOf(errResp).IsZero() {
		return productContext.JSON(errResp.StatusCode, errResp.Error.Error())
	}
	return productContext.JSON(http.StatusOK, dataResp)
}

func (controller ProductController) GetProductsHandler(productContext echo.Context) error {
	request := new(services.GetProductsRequest)
	err := productContext.Bind(request)
	if err != nil {
		return productContext.JSON(http.StatusBadRequest, err.Error())
	}

	dataResp, errResp := controller.App.ProductService.GetProducts(*request)
	if !reflect.ValueOf(errResp).IsZero() {
		return productContext.JSON(errResp.StatusCode, errResp.Error.Error())
	}
	return productContext.JSON(http.StatusOK, dataResp)
}

func (controller ProductController) DeleteProductHandler(productContext echo.Context) error {
	request := new(services.DeleteProductRequest)
	err := productContext.Bind(request)
	if err != nil {
		return productContext.JSON(http.StatusBadRequest, err.Error())
	}

	dataResp, errResp := controller.App.ProductService.DeleteProduct(*request)
	if !reflect.ValueOf(errResp).IsZero() {
		return productContext.JSON(errResp.StatusCode, errResp.Error.Error())
	}
	return productContext.JSON(http.StatusNoContent, dataResp)
}

func (controller ProductController) UpdateProductHandler(productContext echo.Context) error {
	request := new(services.UpdateProductRequest)
	err := productContext.Bind(request)
	if err != nil {
		return productContext.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate request
	if err := productContext.Validate(request); err != nil {
		errors := utils.GetValidationErrors(err)
		return productContext.JSON(http.StatusBadRequest, errors)
	}

	dataResp, errResp := controller.App.ProductService.UpdateProduct(*request)
	if !reflect.ValueOf(errResp).IsZero() {
		return productContext.JSON(errResp.StatusCode, errResp.Error.Error())
	}
	return productContext.JSON(http.StatusOK, dataResp)
}
