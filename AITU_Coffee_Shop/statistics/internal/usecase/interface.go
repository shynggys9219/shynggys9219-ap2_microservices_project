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

type ClientCache interface {
	Get(clientID uint64) (model.Client, bool)
	Set(client model.Client)
	SetMany(clients []model.Client)
	Delete(clientID uint64)
}

type RedisCache interface {
	Get(ctx context.Context, clientID uint64) (model.Client, error)
	Set(ctx context.Context, client model.Client) error
	SetMany(ctx context.Context, clients []model.Client) error
	Delete(ctx context.Context, clientID uint64) error
}
