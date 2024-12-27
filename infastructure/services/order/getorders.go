package services

import (
	"net/http"

	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/entities"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/models"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"gorm.io/gorm/clause"
)

type GetOrderRequest struct {
	UserId     string `query:"userId" validate:"required,uuid"`
	Page       int    `query:"page"`
	PageLength int    `query:"pageLength"`
	SortBy     string `query:"sortBy"`
	Order      string `query:"order"`
}

type GetOrderResponse struct {
	Orders []OrderModel `json:"orders"`
}

type OrderModel struct {
	OrderId    uuid.UUID           `json:"orderId"`
	TotalPrice float64             `json:"totalPrice"`
	TaxAmount  float64             `json:"taxAmount"`
	OrderItems []GetOrderItemModel `json:"orderitems"`
}

type GetOrderItemModel struct {
	OrderItemId        uuid.UUID `json:"orderItemId"`
	ProductName        string    `json:"productName"`
	ProductDescription string    `json:"productDescription"`
	CategoryName       string    `json:"categoryName"`
	ProductPrice       float64   `json:"productPrice"`
	Quantity           int       `json:"quantity"`
	SubTotal           float64   `json:"subTotal"`
}

func (service OrderService) GetOrders(request GetOrderRequest) (GetOrderResponse, models.ErrorResponse) {

	if request.Page < 1 {
		request.Page = 1
	}

	if request.PageLength > 25 {
		request.PageLength = 25
	}

	if request.SortBy == "" {
		request.SortBy = "Id"
	}

	if request.Order == "" {
		request.Order = "asc"
	}

	isDescending := !(request.Order == "asc" || request.Order == "ASC")

	// Calculate offset
	offset := (request.Page - 1) * request.PageLength

	// Get total count of products
	var totalItems int64

	countResult := service.DB.Model(&entities.Order{}).
		Where("UserId = ?", request.UserId).
		Count(&totalItems)

	if countResult.Error != nil {
		return GetOrderResponse{}, models.ErrorResponse{StatusCode: http.StatusInternalServerError, Error: countResult.Error}
	}

	type OrderDTO struct {
		ProductId          uuid.UUID
		ProductName        string
		ProductDescription string
		Id                 uuid.UUID
		TotalAmount        float64
		TaxAmount          float64
		SubTotal           float64
		CategoryName       string
		ProductPrice       float64
		Quantity           int
		OrderItemId        uuid.UUID
	}

	var ordersDTO []OrderDTO
	result := service.DB.Table("orders").
		Where("orders.UserId = ?", request.UserId).
		Where("orders.IsDeprecated = ?", false).
		Joins("join orderitems on orders.Id = orderitems.OrderId").
		Where("orderitems.IsDeprecated = ?", false).
		Joins("join products on orderitems.ProductId = products.Id").
		Joins("left join categories on products.CategoryId = categories.Id").
		Select("products.Id AS ProductId", "products.Name AS ProductName", "products.Price AS ProductPrice", "products.Description AS ProductDescription", "orders.Id AS Id", "orders.TotalAmount AS TotalAmount", "orders.TaxAmount AS TaxAmount", "orderItems.TotalPrice AS SubTotal", "orderItems.Quantity AS Quantity", "orderItems.Id AS OrderItemId", "categories.Name AS CategoryName").
		Offset(offset).
		Limit(request.PageLength).
		Order(clause.OrderByColumn{Column: clause.Column{Name: request.SortBy}, Desc: isDescending}).
		Find(&ordersDTO)

	if result.Error != nil {
		return GetOrderResponse{}, models.ErrorResponse{StatusCode: http.StatusInternalServerError, Error: result.Error}
	}

	ordersByOrderId := lo.GroupBy(ordersDTO, func(order OrderDTO) uuid.UUID {
		return order.Id
	})

	var orders []OrderModel
	for orderID, orderItemsModel := range ordersByOrderId {
		var order OrderModel
		order.OrderId = orderID
		var orderItems []GetOrderItemModel

		for _, orderItemModel := range orderItemsModel {
			var orderItem GetOrderItemModel
			orderItem.ProductName = orderItemModel.ProductName
			orderItem.OrderItemId = orderItemModel.OrderItemId
			orderItem.ProductDescription = orderItemModel.ProductDescription
			orderItem.CategoryName = orderItemModel.CategoryName
			orderItem.ProductPrice = orderItemModel.ProductPrice
			orderItem.Quantity = orderItemModel.Quantity
			orderItem.SubTotal = orderItemModel.SubTotal

			orderItems = append(orderItems, orderItem)
		}

		order.OrderItems = orderItems
		order.TaxAmount = orderItemsModel[0].TaxAmount
		order.TotalPrice = orderItemsModel[0].TotalAmount
		orders = append(orders, order)
	}

	return GetOrderResponse{Orders: orders}, models.ErrorResponse{}
}
