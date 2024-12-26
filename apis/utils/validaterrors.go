package utils

import (
	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/models"
	"github.com/labstack/echo/v4"
)

func GetValidationErrors(err error) *models.ValidationErrorResponse {
	if he, ok := err.(*echo.HTTPError); ok {
		if v, isvalid := he.Message.(models.ValidationErrorResponse); isvalid {
			return &v
		}
	}

	return nil
}
