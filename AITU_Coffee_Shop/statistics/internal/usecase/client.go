package usecase

import (
	"context"
	"github.com/shynggys9219/ap2_microservices_project/statistics/internal/model"
	"github.com/shynggys9219/ap2_microservices_project/statistics/pkg/def"
	"time"
)

type Client struct {
	repo ClientRepo
}

func NewClient(repo ClientRepo) *Client {
	return &Client{
		repo: repo,
	}
}

func (c *Client) Create(ctx context.Context, client model.Client) error {
	client.UpdatedAt = time.Now().UTC()

	return c.repo.Upsert(ctx, client, model.ClientFilter{ID: def.Pointer(client.ID)})
}

func (c *Client) Get(ctx context.Context, id uint64) (model.Client, error) {
	return c.repo.GetWithFilter(ctx, model.ClientFilter{ID: def.Pointer(id)})
}

func (c *Client) List(ctx context.Context) ([]model.Client, error) {
	return c.repo.GetListWithFilter(ctx, model.ClientFilter{})
}
