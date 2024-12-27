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

// Order godoc
// @Summary     Create a order
// @Description    Create a a new customer order
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        CreateOrderRequest  body  services.CreateOrderRequest  true  "CreateOrderRequest"
// @Success      201  {object}  services.CreateOrderResponse
// @Failure      400  {object}  string
// @Failure      404  {object}  []string
// @Router       /api/v1/orders [post]
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

// Order godoc
// @Summary     Get all user order
// @Description    Get all user order
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        GetOrderRequest  body  services.GetOrderRequest  true  "GetOrderRequest"
// @Success      200  {object}  services.GetOrderResponse
// @Failure      400  {object}  string
// @Failure      404  {object}  []string
// @Router       /api/v1/orders [get]
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

// Order godoc
// @Summary     Update user order status
// @Description    Update user order status
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        UpdateStatusRequest  body  services.UpdateStatusRequest  true  "UpdateStatusRequest"
// @Success      201  {object}  services.UpdateStatusResponse
// @Failure      400  {object}  string
// @Failure      404  {object}  []string
// @Router       /api/v1/orders/{id} [put]
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

// Order godoc
// @Summary     Cancel user order
// @Description    Cancel user order
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        CancelOrderRequest  body  services.CancelOrderRequest  true  "CancelOrderRequest"
// @Success      201  {object}  services.CancelOrderResponse
// @Failure      400  {object}  string
// @Failure      404  {object}  []string
// @Router       /api/v1/orders/{id}/cancel [put]
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
