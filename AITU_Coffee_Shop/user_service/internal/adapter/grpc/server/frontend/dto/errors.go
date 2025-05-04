package dto

import (
	"errors"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrEmailAlreadyRegistered = status.Error(codes.InvalidArgument, "this email is already registered")
)

func FromError(err error) error {
	switch {
	case errors.Is(err, model.ErrEmailAlreadyRegistered):
		return ErrEmailAlreadyRegistered
	default:
		return status.Error(codes.Internal, "something went wrong")
	}
}
