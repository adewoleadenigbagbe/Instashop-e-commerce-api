package controllers

import (
	"net/http"
	"reflect"

	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/core"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/apis/utils"
	services "github.com/adewoleadenigbagbe/instashop-e-commerce/infastructure/services/order"
	"github.com/labstack/echo/v4"
)

type OrderController struct {
	App *core.BaseApp
}

func (controller OrderController) CreateOrderHandler(orderContext echo.Context) error {
	request := new(services.CreateOrderRequest)
	err := orderContext.Bind(request)
	if err != nil {
		return orderContext.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate request
	if err := orderContext.Validate(request); err != nil {
		errors := utils.GetValidationErrors(err)
		return orderContext.JSON(http.StatusBadRequest, errors)
	}

	dataResp, errResp := controller.App.OrderService.CreateOrder(*request)
	if !reflect.ValueOf(errResp).IsZero() {
		return orderContext.JSON(errResp.StatusCode, errResp.Error.Error())
	}
	return orderContext.JSON(http.StatusCreated, dataResp)
}

func (controller OrderController) GetOrdersHandler(orderContext echo.Context) error {
	request := new(services.GetOrderRequest)
	err := orderContext.Bind(request)
	if err != nil {
		return orderContext.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate request
	if err := orderContext.Validate(request); err != nil {
		errors := utils.GetValidationErrors(err)
		return orderContext.JSON(http.StatusBadRequest, errors)
	}

	dataResp, errResp := controller.App.OrderService.GetOrders(*request)
	if !reflect.ValueOf(errResp).IsZero() {
		return orderContext.JSON(errResp.StatusCode, errResp.Error.Error())
	}
	return orderContext.JSON(http.StatusOK, dataResp)
}

func (controller OrderController) UpdateOrderStatusHandler(orderContext echo.Context) error {
	request := new(services.UpdateStatusRequest)
	err := orderContext.Bind(request)
	if err != nil {
		return orderContext.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate request
	if err := orderContext.Validate(request); err != nil {
		errors := utils.GetValidationErrors(err)
		return orderContext.JSON(http.StatusBadRequest, errors)
	}

	dataResp, errResp := controller.App.OrderService.UpdateOrderStatus(*request)
	if !reflect.ValueOf(errResp).IsZero() {
		return orderContext.JSON(errResp.StatusCode, errResp.Error.Error())
	}
	return orderContext.JSON(http.StatusCreated, dataResp)
}

func (controller OrderController) CancelOrderHandler(orderContext echo.Context) error {
	request := new(services.CancelOrderRequest)
	err := orderContext.Bind(request)
	if err != nil {
		return orderContext.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate request
	if err := orderContext.Validate(request); err != nil {
		errors := utils.GetValidationErrors(err)
		return orderContext.JSON(http.StatusBadRequest, errors)
	}

	dataResp, errResp := controller.App.OrderService.CancelOrder(*request)
	if !reflect.ValueOf(errResp).IsZero() {
		return orderContext.JSON(errResp.StatusCode, errResp.Error.Error())
	}
	return orderContext.JSON(http.StatusCreated, dataResp)
}
