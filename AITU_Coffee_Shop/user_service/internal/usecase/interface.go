package usecase

import (
	"context"

	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/model"
)

type AiRepo interface {
	Next(ctx context.Context, collection string) (uint64, error)
}

type CustomerRepo interface {
	Create(ctx context.Context, customer model.Customer) error
	Update(ctx context.Context, filter model.CustomerFilter, update model.CustomerUpdateData) error
	GetWithFilter(ctx context.Context, filter model.CustomerFilter) (model.Customer, error)
	GetListWithFilter(ctx context.Context, filter model.CustomerFilter) ([]model.Customer, error)
}

type RefreshTokenRepo interface {
	Create(ctx context.Context, session model.Session) error
	GetByToken(ctx context.Context, token string) (model.Session, error)
	DeleteByToken(ctx context.Context, token string) error
}

type CustomerEventStorage interface {
	Push(ctx context.Context, client model.Customer) error
}
