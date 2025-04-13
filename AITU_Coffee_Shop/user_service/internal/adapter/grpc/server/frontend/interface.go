package frontend

import (
	"context"

	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/model"
)

type ClientUsecase interface {
	Create(ctx context.Context, client model.Client) (model.Client, error)
}
