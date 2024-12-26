package services

import (
	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/entities"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/models"
	"github.com/google/uuid"

	"net/http"
)

type DeleteProductRequest struct {
	Id uuid.UUID `param:"id"`
}

type DeleteProductResponse struct {
	ProductId uuid.UUID `json:"productId"`
}

func (service ProductService) DeleteProduct(request DeleteProductRequest) (DeleteProductResponse, models.ErrorResponse) {
	var product entities.Product
	result := service.DB.Delete(&product, request.Id)
	if result.Error != nil {
		return DeleteProductResponse{}, models.ErrorResponse{StatusCode: http.StatusInternalServerError, Error: result.Error}
	}

	return DeleteProductResponse{ProductId: request.Id}, models.ErrorResponse{}
}
