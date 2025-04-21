package handler

import (
	"context"
	"github.com/shynggys9219/ap2_microservices_project/statistics/internal/adapter/nats/handler/dto"
	"log"

	"github.com/nats-io/nats.go"
)

type Client struct {
	usecase ClientUsecase
}

func NewClient(usecase ClientUsecase) *Client {
	return &Client{usecase: usecase}
}

func (c *Client) Handler(ctx context.Context, msg *nats.Msg) error {
	client, err := dto.ToClient(msg)
	if err != nil {
		log.Println("failed to convert Client NATS msg", err)

		return err
	}

	err = c.usecase.Create(ctx, client)
	if err != nil {
		log.Println("failed to create many Bonus", err)

		return err
	}

	return nil
}
