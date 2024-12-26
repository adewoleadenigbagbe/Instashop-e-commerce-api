package services

import (
	"errors"
	"net/http"

	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/entities"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/enums"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/models"
	"github.com/google/uuid"
)

type CancelOrderRequest struct {
	Id uuid.UUID `json:"orderId" validate:"required,uuid"`
}

type CancelOrderResponse struct {
	Id uuid.UUID `json:"orderId"`
}

func (service OrderService) CancelOrder(request CancelOrderRequest) (CancelOrderResponse, models.ErrorResponse) {

	var order entities.Order
	if err := service.DB.First(&order, request.Id).Error; err != nil {
		return CancelOrderResponse{}, models.ErrorResponse{StatusCode: http.StatusNotFound, Error: errors.New("order info not found")}
	}

	//get the product count back, this will be in a transaction

	order.Status = enums.Cancelled

	result := service.DB.Save(&order)
	if result.Error != nil {
		return CancelOrderResponse{}, models.ErrorResponse{StatusCode: http.StatusInternalServerError, Error: result.Error}
	}

	return CancelOrderResponse{Id: order.ID}, models.ErrorResponse{}

}
