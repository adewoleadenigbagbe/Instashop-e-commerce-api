package models

type ErrorResponse struct {
	Error      error
	StatusCode int
}
