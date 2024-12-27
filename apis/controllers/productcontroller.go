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

// Product godoc
// @Summary     Create a new product
// @Description    Create a new product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        CreateProductRequest  body  services.CreateProductRequest  true  "CreateProductRequest"
// @Success      201  {object}  services.CreateProductResponse
// @Failure      400  {object}  string
// @Failure      404  {object}  []string
// @Router       /api/v1/products [post]
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

// Product godoc
// @Summary     Get a product by the ID
// @Description    Get a product by the ID
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        GetProductRequestById  body  services.GetProductRequestById  true  "GetProductRequestById"
// @Success      200  {object}  services.GetProductByIdResponse
// @Failure      400  {object}  string
// @Failure      404  {object}  []string
// @Router       /api/v1/products/{id} [get]
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

// Product godoc
// @Summary     Get a products
// @Description    Get a products
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        GetProductsRequest  body  services.GetProductsRequest  true  "GetProductsRequest"
// @Success      200  {object}  services.Result
// @Failure      400  {object}  string
// @Failure      404  {object}  []string
// @Router       /api/v1/products [get]
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

// Product godoc
// @Summary     Delete a product by id
// @Description    Delete a product by id
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        DeleteProductRequest  body  services.DeleteProductRequest  true  "DeleteProductRequest"
// @Success      204  {object}  services.Result
// @Failure      400  {object}  string
// @Failure      404  {object}  []string
// @Router       /api/v1/products/{id} [delete]
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

// Product godoc
// @Summary     Update product with product info
// @Description    Update product with product info
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        UpdateProductRequest  body  services.UpdateProductRequest  true  "UpdateProductRequest"
// @Success      201  {object}  services.UpdateProductResponse
// @Failure      400  {object}  string
// @Failure      404  {object}  []string
// @Router       /api/v1/products/{id} [put]
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
	return productContext.JSON(http.StatusCreated, dataResp)
}
