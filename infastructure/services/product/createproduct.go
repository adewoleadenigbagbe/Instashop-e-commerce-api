package services

import (
	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/entities"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/enums"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/models"
	"github.com/google/uuid"

	"gorm.io/gorm"

	"net/http"
)

type ProductService struct {
	DB *gorm.DB
}

type CreateProductRequest struct {
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Price       float64             `json:"price"`
	CategoryId  uuid.UUID           `json:"categoryId"`
	Stock       int                 `json:"stock"`
	Status      enums.ProductStatus `json:"status"`
	ImageURL    string              `json:"imageUrl"`
}

type CreateProductResponse struct {
	ProductId uuid.UUID `json:"productId"`
}

func (service ProductService) CreateProduct(request CreateProductRequest) (CreateProductResponse, models.ErrorResponse) {
	var err error
	productId, _ := uuid.NewV7()
	product := entities.Product{
		ID:          productId,
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		CategoryID:  request.CategoryId,
		Stock:       request.Stock,
		Status:      request.Status,
		ImageURL:    request.ImageURL,
	}

	if err = service.DB.Create(&product).Error; err != nil {
		return CreateProductResponse{}, models.ErrorResponse{StatusCode: http.StatusInternalServerError, Error: err}
	}

	return CreateProductResponse{ProductId: productId}, models.ErrorResponse{}
}
