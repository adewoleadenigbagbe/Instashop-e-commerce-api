package middlewares

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/adewoleadenigbagbe/instashop-e-commerce/shared/models"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewValidator() *CustomValidator {
	v := validator.New()

	// Register function to get json tag name instead of struct field name
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return &CustomValidator{validator: v}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Validation errors
		validationErrors := err.(validator.ValidationErrors)

		// Group errors by field
		fieldErrors := make(map[string][]string)

		for _, err := range validationErrors {
			field := err.Field()
			message := getErrorMessage(err)
			fieldErrors[field] = append(fieldErrors[field], message)
		}

		// Convert map to slice of ValidationError
		var errors []models.ValidationError
		for field, messages := range fieldErrors {
			errors = append(errors, models.ValidationError{
				Field:   field,
				Message: messages,
			})
		}

		// Return custom error response
		return echo.NewHTTPError(http.StatusBadRequest, models.ValidationErrorResponse{
			Errors: errors,
		})
	}

	return nil
}

func getErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Must be a valid email address"
	case "min":
		if err.Type().Kind() == reflect.String {
			return "Must be at least " + err.Param() + " characters long"
		}
		return "Must be at least " + err.Param()
	case "max":
		if err.Type().Kind() == reflect.String {
			return "Must not exceed " + err.Param() + " characters"
		}
		return "Must not exceed " + err.Param()
	case "oneof":
		return "Must be one of: " + strings.Replace(err.Param(), " ", ", ", -1)
	case "gt":
		return "Must be greater than " + err.Param()
	case "gte":
		return "Must be greater than or equal to " + err.Param()
	case "lt":
		return "Must be less than " + err.Param()
	case "lte":
		return "Must be less than or equal to " + err.Param()
	case "url":
		return "Must be a valid URL"
	case "uuid":
		return "Must be a valid UUID"
	case "e164":
		return "Must be a valid phone number"
	}
	return "Invalid value"
}

// Custom error handler for Echo framework
func CustomHTTPErrorHandler(err error, c echo.Context) error {
	// Check if it's an Echo HTTP error
	if he, ok := err.(*echo.HTTPError); ok {
		return c.JSON(he.Code, he.Message)
	}
	// Return generic error for other errors
	return c.JSON(http.StatusInternalServerError, map[string]string{
		"message": "Internal server error",
	})
}
