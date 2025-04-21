package usecase

import (
	"context"
	"github.com/shynggys9219/ap2_microservices_project/api-gateway/internal/model"
)

type ClientPresenter interface {
	Create(ctx context.Context, request model.Client) (model.Client, error)
	Update(ctx context.Context, request model.Client) (model.Client, error)
	Get(ctx context.Context, id uint64) (model.Client, error)
	Delete(ctx context.Context, id uint64) error
}

type ClientStatisticPresenter interface {
	Get(ctx context.Context, id uint64) (model.ClientStatistic, error)
	List(ctx context.Context) ([]model.ClientStatistic, error)
}
