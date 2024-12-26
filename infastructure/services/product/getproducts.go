package services

import (
	"net/http"

	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/entities"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/enums"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/models"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

type GetProductsRequest struct {
	Page       int    `query:"page"`
	PageLength int    `query:"pageLength"`
	SortBy     string `query:"sortBy"`
	Order      string `query:"order"`
}

type GetProductResponse struct {
	Id                  uuid.UUID           `json:"id"`
	Name                string              `json:"name"`
	Description         string              `json:"description"`
	Stock               int                 `json:"stock"`
	Status              enums.ProductStatus `json:"status"`
	ImageUrl            string              `json:"imageUrl"`
	CategoryName        string              `json:"categoryName"`
	CategoryDescription string              `json:"categoryDescription"`
	CategoryType        enums.CategoryType  `json:"categoryType"`
}

type Result struct {
	Page                int                  `json:"page"`
	PerPage             int                  `json:"perPage"`
	TotalPage           int                  `json:"totalPage"`
	TotalResults        int64                `json:"totalResults"`
	RequestedPageLength int                  `json:"requestedPageLength"`
	Products            []GetProductResponse `json:"products"`
}

func (service ProductService) GetProducts(request GetProductsRequest) (Result, models.ErrorResponse) {
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
	if err := service.DB.Model(&entities.Product{}).Count(&totalItems).Error; err != nil {
		return Result{}, models.ErrorResponse{StatusCode: http.StatusInternalServerError, Error: err}
	}

	var products []GetProductResponse
	// Query products with pagination
	result := service.DB.Table("products").
		Joins("left join categories on categories.Id = products.CategoryId").
		Select("products.Id AS Id","products.Name", "products.Description", "products.Stock", "products.Status", "products.ImageUrl", "categories.Name AS CategoryName", "categories.Description AS CategoryDescription", "categories.Type AS CategoryType").
		Offset(offset).
		Limit(request.PageLength).
		Order(clause.OrderByColumn{Column: clause.Column{Name: request.SortBy}, Desc: isDescending}).
		Find(&products)

	if result.Error != nil {
		return Result{}, models.ErrorResponse{StatusCode: http.StatusInternalServerError, Error: result.Error}
	}

	// Calculate pagination metadata
	totalPages := int(totalItems+int64(request.Page)-1) / request.PageLength

	response := Result{
		Products: products,
	}

	response.Page = request.Page
	response.RequestedPageLength = request.PageLength
	response.TotalPage = totalPages
	response.TotalResults = totalItems
	response.PerPage = len(products)

	return response, models.ErrorResponse{}
}
