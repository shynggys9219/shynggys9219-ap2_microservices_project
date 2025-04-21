package handler

import (
	"context"
	"github.com/shynggys9219/ap2_microservices_project/statistics/internal/model"
)

type ClientUsecase interface {
	Create(ctx context.Context, client model.Client) error
}

type OrderUsecase interface {
	Create(ctx context.Context, client model.Order) error
}
