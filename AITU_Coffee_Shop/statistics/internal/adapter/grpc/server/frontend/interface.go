package frontend

import (
	"context"
	"github.com/shynggys9219/ap2_microservices_project/statistics/internal/model"
)

type ClientUsecase interface {
	Get(ctx context.Context, id uint64) (model.Client, error)
	// List TODO: add pagination for list
	List(ctx context.Context) ([]model.Client, error)
}
