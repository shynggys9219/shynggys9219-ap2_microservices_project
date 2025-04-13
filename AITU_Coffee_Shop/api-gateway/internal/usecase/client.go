package usecase

import (
	"context"
	"github.com/shynggys9219/ap2_microservices_project/api-gateway/internal/model"
)

type Client struct {
	presenter ClientPresenter
}

func NewClient(presenter ClientPresenter) *Client {
	return &Client{
		presenter: presenter,
	}
}

func (c *Client) Create(ctx context.Context, request model.Client) (model.Client, error) {
	client, err := c.presenter.Create(ctx, request)
	if err != nil {
		return model.Client{}, err
	}

	return client, nil
}

func (c *Client) Update(ctx context.Context, request model.Client) (model.Client, error) {
	client, err := c.presenter.Update(ctx, request)
	if err != nil {
		return model.Client{}, err
	}

	return client, nil
}

func (c *Client) Get(ctx context.Context, id uint64) (model.Client, error) {
	return c.presenter.Get(ctx, id)
}

func (c *Client) Delete(ctx context.Context, id uint64) error {
	return c.presenter.Delete(ctx, id)
}
