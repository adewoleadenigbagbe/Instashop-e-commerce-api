package services

import (
	"net/http"

	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/entities"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/enums"
	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryService struct {
	DB *gorm.DB
}

type CreateCategoryRequest struct {
	Description string             `json:"description" validate:"required"`
	Type        enums.CategoryType `json:"type"`
}

type CreateCategoryResponse struct {
	Id uuid.UUID `json:"id"`
}

func (service CategoryService) CreateProduct(request CreateCategoryRequest) (CreateCategoryResponse, models.ErrorResponse) {
	var err error
	categoryId, _ := uuid.NewV7()
	categoryName, err := request.Type.GetCategoryName()
	if err != nil{
		return CreateCategoryResponse{}, models.ErrorResponse{StatusCode: http.StatusInternalServerError, Error: err}
	}
	category := entities.Category{
		ID:          categoryId,
		Name:        categoryName,
		Description: request.Description,
		Type:        request.Type,
	}

	if err = service.DB.Create(&category).Error; err != nil {
		return CreateCategoryResponse{}, models.ErrorResponse{StatusCode: http.StatusInternalServerError, Error: err}
	}

	return CreateCategoryResponse{Id: categoryId}, models.ErrorResponse{}
}
