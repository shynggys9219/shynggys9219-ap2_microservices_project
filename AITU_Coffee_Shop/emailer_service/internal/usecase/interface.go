package usecase

import (
	"context"

	"github.com/shynggys9219/ap2_microservices_project/emailer_service/internal/model"
)

type EmailTemplateRenderer interface {
	Render(templateName string, data map[string]any) (string, error)
}

type EmailPresenter interface {
	Send(ctx context.Context, customer model.Customer) error
}
