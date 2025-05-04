package frontend

import (
	"context"

	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/model"
)

type CustomerUsecase interface {
	Register(ctx context.Context, request model.Customer) (uint64, error)
	Update(ctx context.Context, request model.Customer) (model.Customer, error)
	Get(ctx context.Context, id uint64) (model.Customer, error)
	Delete(ctx context.Context, id uint64) error
	Login(ctx context.Context, request model.Token) (model.Token, error)
	RefreshToken(ctx context.Context, request model.Token) (model.Token, error)
}
