package usecase

import (
	"context"

	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/model"
)

type AiRepo interface {
	Next(ctx context.Context, collection string) (uint64, error)
}

type ClientRepo interface {
	Create(ctx context.Context, client model.Client) error
	Update(ctx context.Context, filter model.ClientFilter, update model.ClientUpdateData) error
	GetWithFilter(ctx context.Context, filter model.ClientFilter) (model.Client, error)
	GetListWithFilter(ctx context.Context, filter model.ClientFilter) ([]model.Client, error)
}
