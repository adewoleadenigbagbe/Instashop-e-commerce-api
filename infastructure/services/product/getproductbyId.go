package services

import (
	"errors"

	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/enums"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/models"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"net/http"
)

type GetProductRequestById struct {
	Id uuid.UUID `param:"id"`
}

type GetProductByIdResponse struct {
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

func (service ProductService) GetProductByID(request GetProductRequestById) (GetProductByIdResponse, models.ErrorResponse) {
	var productResponse GetProductByIdResponse
	productQuery := service.DB.Table("products").
		Where("products.Id = ?", request.Id).
		Joins("left join categories on products.CategoryId = categories.Id").
		Select("products.Id,products.Name", "products.Description", "products.Stock", "products.Status", "products.ImageUrl", "categories.Name AS CategoryName", "categories.Description AS CategoryDescription", "categories.Type AS CategoryType").
		First(&productResponse)

	if productQuery.Error != nil {
		if errors.Is(productQuery.Error, gorm.ErrRecordNotFound) {
			return GetProductByIdResponse{}, models.ErrorResponse{StatusCode: http.StatusNotFound, Error: errors.New("product info not found")}
		}
		return GetProductByIdResponse{}, models.ErrorResponse{StatusCode: http.StatusInternalServerError, Error: productQuery.Error}
	}

	return productResponse, models.ErrorResponse{}
}
