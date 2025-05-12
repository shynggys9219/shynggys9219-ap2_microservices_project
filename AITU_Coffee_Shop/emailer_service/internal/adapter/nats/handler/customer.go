package handler

import (
	"context"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/shynggys9219/ap2_microservices_project/emailer_service/internal/adapter/nats/handler/dto"
)

type Customer struct {
	usecase CustomerUsecase
}

func NewCustomer(usecase CustomerUsecase) *Customer {
	return &Customer{usecase: usecase}
}

func (c *Customer) Handler(ctx context.Context, msg *nats.Msg) error {
	client, err := dto.ToCustomer(msg)
	if err != nil {
		log.Println("failed to convert Client NATS msg", err)

		return err
	}

	err = c.usecase.Send(ctx, client)
	if err != nil {
		log.Println("failed to create many Bonus", err)

		return err
	}

	return nil
}
