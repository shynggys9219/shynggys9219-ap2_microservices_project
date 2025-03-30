package handler

import (
	"context"

	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/model"
)

type ClientUsecase interface {
	Create(ctx context.Context, request model.Client) (model.Client, error)
	Update(ctx context.Context, request model.Client) (model.Client, error)
	Get(ctx context.Context, id uint64) (model.Client, error)
	Delete(ctx context.Context, id uint64) error
}
