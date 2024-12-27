package services

import (
	"errors"
	"net/http"

	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/entities"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/enums"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CancelOrderRequest struct {
	Id uuid.UUID `param:"id" validate:"required,uuid"`
}

type CancelOrderResponse struct {
	Id uuid.UUID `json:"orderId"`
}

func (service OrderService) CancelOrder(request CancelOrderRequest) (CancelOrderResponse, models.ErrorResponse) {

	var order entities.Order
	if err := service.DB.
		Where("Id = ?", request.Id).
		Where("Status != ?", enums.Cancelled).
		First(&order).Error; err != nil {
		return CancelOrderResponse{}, models.ErrorResponse{StatusCode: http.StatusNotFound, Error: errors.New("order info not found")}
	}

	var orderProductInfos []struct {
		Id        uuid.UUID
		ProductId uuid.UUID
		Quantity  int
	}

	result := service.DB.Table("orders").
		Where("orders.Id = ?", request.Id).
		Where("orders.IsDeprecated = ?", false).
		Joins("join orderitems on orders.Id = orderitems.OrderId").
		Where("orderitems.IsDeprecated = ?", false).
		Joins("join products on orderitems.ProductId = products.Id").
		Select("products.Id AS ProductId", "orders.Id AS Id", "orderItems.Quantity AS Quantity").
		Find(&orderProductInfos)

	if result.Error != nil {
		return CancelOrderResponse{}, models.ErrorResponse{StatusCode: http.StatusInternalServerError, Error: result.Error}
	}

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		//update order
		tx.Model(&entities.Order{}).Where("Id = ?", request.Id).Update("Status", enums.Cancelled)
		if tx.Error != nil {
			return tx.Error
		}

		//increment the product back
		for _, orderProductInfo := range orderProductInfos {
			tx.Model(&entities.Product{}).Where("Id = ?", orderProductInfo.ProductId).UpdateColumn("Stock", gorm.Expr("Stock + ?", orderProductInfo.Quantity))
			if tx.Error != nil {
				return tx.Error
			}
		}

		// return nil will commit the whole transaction
		return nil
	})

	if err != nil {
		return CancelOrderResponse{}, models.ErrorResponse{StatusCode: http.StatusInternalServerError, Error: err}
	}

	return CancelOrderResponse{Id: order.ID}, models.ErrorResponse{}
}
