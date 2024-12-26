package services

import (
	"net/http"

	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/models"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"gorm.io/gorm/clause"
)

type GetOrderRequest struct {
	UserId     string `json:"userId" validate:"required,uuid"`
	Page       int    `query:"page"`
	PageLength int    `query:"pageLength"`
	SortBy     string `query:"sortBy"`
	Order      string `query:"order"`
}

type GetOrderResponse struct {
	Orders []OrderModel `json:"orders"`
}

type OrderModel struct {
	OrderId    uuid.UUID        `json:"orderId"`
	TotalPrice float64          `json:"totalPrice"`
	TaxAmount  float64          `json:"taxAmount"`
	OrderItems []OrderItemModel `json:"orderitems"`
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

	countResult := service.DB.Table("orders").
		Where("orders.UserId = ?", request.UserId).
		Where("orders.IsDeprecated = ?", false).
		Joins("join orderitems on orders.Id = orderitems.OrderId").
		Where("orderitems.IsDeprecated = ?", false).
		Joins("join products on orderitems.ProductId = product.Id").
		Joins("join categories on products.CategoryId = categories.Id").
		Select("products.Id AS Id,products.Name", "products.Description", "products.Stock", "products.Status", "products.ImageUrl", "categories.Name AS CategoryName", "categories.Description AS CategoryDescription", "categories.Type AS CategoryType").
		Offset(offset).
		Limit(request.PageLength).
		Order(clause.OrderByColumn{Column: clause.Column{Name: request.SortBy}, Desc: isDescending}).
		Count(&totalItems)

	if countResult.Error != nil {
		return GetOrderResponse{}, models.ErrorResponse{StatusCode: http.StatusInternalServerError, Error: err}
	}

	type OrderDTO struct {
		ProductId          uuid.UUID
		ProductName        string
		ProductDescription string
		OrderId            uuid.UUID
		TotalPrice         float64
		TaxAmount          float64
		SubTotal           float64
		CategoryName       string
		ProductPrice       float64
		Quantity           int
		OrderItemId        uuid.UUID
	}

	var orders []OrderDTO
	result := service.DB.Table("orders").
		Where("orders.UserId = ?", request.UserId).
		Where("orders.IsDeprecated = ?", false).
		Joins("join orderitems on orders.Id = orderitems.OrderId").
		Where("orderitems.IsDeprecated = ?", false).
		Joins("join products on orderitems.ProductId = product.Id").
		Joins("join categories on products.CategoryId = categories.Id").
		Select("products.Id AS ProductId", "products.Name AS ProductName", "products.Price AS ProductPrice", "products.Description AS ProductDescription", "orders.Id AS OrderId", "orders.TotalPrice AS TotalPrice", "order.TaxAmount AS TaxAmount", "orderItems.TotalPrice AS SubTotal", "orderItems.Quantity AS Quantity", "orderItems.Id AS OrderItemId", "categories.Name AS CategoryName").
		Offset(offset).
		Limit(request.PageLength).
		Order(clause.OrderByColumn{Column: clause.Column{Name: request.SortBy}, Desc: isDescending}).
		Find(&orders)


	    // First, group by OrderId
		ordersByOrderId := lo.GroupBy(orders, func(order OrderDTO) uuid.UUID {
			return order.OrderId
		})
	
		// Then, for each OrderId group, group by OrderItemId
		ordersByOrderIdAndItemId := lo.MapValues(ordersByOrderId, func(orderItems []OrderDTO, _ uuid.UUID) map[uuid.UUID]OrderDTO {
			return lo.KeyBy(orderItems, func(order OrderDTO) uuid.UUID {
				return order.OrderItemId
			})
		})
}
