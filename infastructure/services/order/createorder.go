package services

import (
	"net/http"

	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/entities"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/enums"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/models"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/samber/lo"
)

type OrderService struct {
	DB *gorm.DB
}

type CreateOrderRequest struct {
	UserId     uuid.UUID        `json:"userId" validate:"required,uuid"`
	OrderItems []OrderItemModel `json:"orderitems" validate:"required,dive,min=1"`
}

type OrderItemModel struct {
	ProductId uuid.UUID `json:"productId" validate:"required,uuid"`
	Quantity  int       `json:"quantity" validate:"required,gt=0"`
}

type CreateOrderResponse struct {
	Id uuid.UUID `json:"orderId"`
}

func (service OrderService) CreateOrder(request CreateOrderRequest) (CreateOrderResponse, models.ErrorResponse) {
	var err error

	orderId, _ := uuid.NewV7()

	productIds := lo.Map(request.OrderItems, func(item OrderItemModel, index int) uuid.UUID {
		return item.ProductId
	})

	var ProductInfos []struct {
		Id    uuid.UUID
		Price float64
	}

	//check the stock for each product
	result := service.DB.Table("products").
		Where("Id IN ?", productIds).
		Where("IsDeprecated = ?", false).
		Where("Status = ?", enums.Active).
		Select("Id, Price").
		Find(&ProductInfos)

	if result.Error != nil {
		return CreateOrderResponse{}, models.ErrorResponse{StatusCode: http.StatusInternalServerError, Error: result.Error}
	}

	productMap := lo.KeyBy(request.OrderItems, func(item OrderItemModel) uuid.UUID {
		return item.ProductId
	})

	var totalAmount float64 = 0
	orderItems := make([]*entities.OrderItem, len(ProductInfos))
	for _, p := range ProductInfos {
		orderItemId, _ := uuid.NewV7()
		orderItem := &entities.OrderItem{
			ID:        orderItemId,
			OrderID:   orderId,
			ProductID: p.Id,
			Quantity:  productMap[p.Id].Quantity,
			UnitPrice: p.Price,
		}

		orderItem.CalculateTotal()
		totalAmount += orderItem.TotalPrice

		orderItems = append(orderItems, orderItem)
	}

	order := &entities.Order{
		ID:          orderId,
		UserID:      request.UserId,
		Status:      enums.Pending,
		TotalAmount: totalAmount,
		TaxAmount:   models.TaxPercentage * totalAmount,
	}

	err = service.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(orderItems).Error; err != nil {
			// return any error will rollback
			return err
		}

		if err := tx.Create(order).Error; err != nil {
			// return any error will rollback
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})

	if err != nil {
		return CreateOrderResponse{}, models.ErrorResponse{StatusCode: http.StatusInternalServerError, Error: err}
	}

	return CreateOrderResponse{Id: orderId}, models.ErrorResponse{}

}
