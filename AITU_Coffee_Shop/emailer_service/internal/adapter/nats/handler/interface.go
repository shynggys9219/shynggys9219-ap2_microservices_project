package handler

import (
	"context"

	"github.com/shynggys9219/ap2_microservices_project/emailer_service/internal/model"
)

type CustomerUsecase interface {
	Send(ctx context.Context, customer model.Customer) error
}
