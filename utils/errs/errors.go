package errs

import (
	"log"
	"net/http"
)

type AppError struct {
	Code             int                 `json:",omitempty"`
	Message          string              `json:"message,omitempty"`
	ValidationErrors map[string][]string `json:"errors,omitempty"`
}

func (e AppError) ToMessage() *AppError {
	log.Println(e.Message, e.ValidationErrors)
	return &AppError{
		Message:          e.Message,
		ValidationErrors: e.ValidationErrors,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func NewValidationError(errors map[string][]string) *AppError {
	return &AppError{
		ValidationErrors: errors,
		Code:             http.StatusBadRequest,
	}
}
