package usecase

import (
	"context"
	"github.com/shynggys9219/ap2_microservices_project/statistics/internal/model"
)

type ClientRepo interface {
	Upsert(ctx context.Context, client model.Client, filter model.ClientFilter) error
	GetWithFilter(ctx context.Context, filter model.ClientFilter) (model.Client, error)
	GetListWithFilter(ctx context.Context, filter model.ClientFilter) ([]model.Client, error)
}
