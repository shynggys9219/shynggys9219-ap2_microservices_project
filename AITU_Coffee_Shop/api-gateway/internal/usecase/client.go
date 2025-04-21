package usecase

import (
	"context"
	"github.com/shynggys9219/ap2_microservices_project/api-gateway/internal/model"
)

type Client struct {
	clientPresenter ClientPresenter
}

func NewClient(presenter ClientPresenter) *Client {
	return &Client{
		clientPresenter: presenter,
	}
}

func (c *Client) Create(ctx context.Context, request model.Client) (model.Client, error) {
	client, err := c.clientPresenter.Create(ctx, request)
	if err != nil {
		return model.Client{}, err
	}

	return client, nil
}

func (c *Client) Update(ctx context.Context, request model.Client) (model.Client, error) {
	client, err := c.clientPresenter.Update(ctx, request)
	if err != nil {
		return model.Client{}, err
	}

	return client, nil
}

func (c *Client) Get(ctx context.Context, id uint64) (model.Client, error) {
	return c.clientPresenter.Get(ctx, id)
}

func (c *Client) Delete(ctx context.Context, id uint64) error {
	return c.clientPresenter.Delete(ctx, id)
}
