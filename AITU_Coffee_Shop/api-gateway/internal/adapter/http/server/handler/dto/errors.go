package dto

import (
	"errors"
	"github.com/shynggys9219/ap2_microservices_project/api-gateway/internal/model"
	"net/http"
)

type HTTPError struct {
	Code    int
	Message string
}

var (
	ErrInvalidEmail = &HTTPError{
		Code:    http.StatusBadRequest,
		Message: "only @gmail.com and @astanait.edu.kz are allowed",
	}

	ErrInvalidPassword = &HTTPError{
		Code:    http.StatusBadRequest,
		Message: "password must contain at least 8 symbols, 1 capital letter and 1 special symbol",
	}

	ErrInvalidID = &HTTPError{
		Code:    http.StatusBadRequest,
		Message: "you provided wrong id",
	}
)

func FromError(err error) *HTTPError {
	switch {
	case errors.Is(err, model.ErrInvalidEmail):
		return ErrInvalidEmail
	case errors.Is(err, model.ErrInvalidPassword):
		return ErrInvalidPassword
	case errors.Is(err, model.ErrInvalidID):
		return ErrInvalidID
	default:
		return &HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "something went wrong",
		}
	}
}
