package models

type ErrorResponse struct {
	Error      error
	StatusCode int
}

type ValidationError struct {
	Field   string   `json:"field"`
	Message []string `json:"messages"`
}

type ValidationErrorResponse struct {
	Errors []ValidationError `json:"errors"`
}
