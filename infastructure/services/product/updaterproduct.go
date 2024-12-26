package services

import (
	"errors"

	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/entities"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/enums"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/models"
	"github.com/google/uuid"

	"net/http"
)

type UpdateProductRequest struct {
	ID          uuid.UUID           `param:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Price       float64             `json:"price"`
	CategoryId  uuid.UUID           `json:"categoryId"`
	Stock       int                 `json:"stock"`
	Status      enums.ProductStatus `json:"status"`
	ImageURL    string              `json:"imageUrl"`
}

type UpdateProductResponse struct {
	ProductId uuid.UUID `json:"productId"`
}

func (service ProductService) UpdateProduct(request UpdateProductRequest) (UpdateProductResponse, models.ErrorResponse) {
	var product entities.Product

	if err := service.DB.First(&product, request.ID).Error; err != nil {
		return UpdateProductResponse{}, models.ErrorResponse{StatusCode: http.StatusNotFound, Error: errors.New("product info not found")}
	}

	product.CategoryID = request.CategoryId
	product.Name = request.Name
	product.Description = request.Description
	product.Price = request.Price
	product.Stock = request.Stock
	product.ImageURL = request.ImageURL

	result := service.DB.Save(&product)
	if result.Error != nil {
		return UpdateProductResponse{}, models.ErrorResponse{StatusCode: http.StatusInternalServerError, Error: result.Error}
	}

	return UpdateProductResponse{ProductId: product.ID}, models.ErrorResponse{}
}
