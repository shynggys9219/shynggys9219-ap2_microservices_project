package frontend

import (
	"context"

	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/model"
)

type CustomerUsecase interface {
	Register(ctx context.Context, request model.Customer) (uint64, error)
	Update(ctx context.Context, token string, request model.Customer) (model.Customer, error)
	Get(ctx context.Context, token string, id uint64) (model.Customer, error)
	Delete(ctx context.Context, id uint64) error
	Login(ctx context.Context, email, password string) (model.Token, error)
	RefreshToken(ctx context.Context, refreshToken string) (model.Token, error)
}
