package services

import (
	"errors"
	"net/http"

	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/entities"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/enums"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/models"
	"github.com/google/uuid"
)

type UpdateStatusRequest struct {
	Id     uuid.UUID         `param:"id" validate:"required,uuid"`
	Status enums.OrderStatus `json:"status" validate:"required"`
}

type UpdateStatusResponse struct {
	Id uuid.UUID `json:"orderId"`
}

func (service OrderService) UpdateOrderStatus(request UpdateStatusRequest) (UpdateStatusResponse, models.ErrorResponse) {

	var order entities.Order
	if err := service.DB.First(&order, request.Id).Error; err != nil {
		return UpdateStatusResponse{}, models.ErrorResponse{StatusCode: http.StatusNotFound, Error: errors.New("order info not found")}
	}

	order.Status = request.Status

	result := service.DB.Save(&order)
	if result.Error != nil {
		return UpdateStatusResponse{}, models.ErrorResponse{StatusCode: http.StatusInternalServerError, Error: result.Error}
	}

	return UpdateStatusResponse{Id: order.ID}, models.ErrorResponse{}

}
