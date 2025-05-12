package usecase

import (
	"context"
	"github.com/shynggys9219/ap2_microservices_project/emailer_service/internal/model"
)

type Customer struct {
	renderer EmailTemplateRenderer
	sender   EmailPresenter // Interface implemented by MailerSend adapter

}

func NewCustomer(renderer EmailTemplateRenderer, sender EmailPresenter) *Customer {
	return &Customer{
		renderer: renderer,
		sender:   sender,
	}
}

func (c *Customer) Send(ctx context.Context, customer model.Customer) error {
	return c.sender.Send(ctx, customer)
}
