package usecase

import (
	"context"
	"fmt"
	"github.com/shynggys9219/ap2_microservices_project/statistics/internal/model"
	"github.com/shynggys9219/ap2_microservices_project/statistics/pkg/def"
	"time"
)

type Client struct {
	repo          ClientRepo
	inMemoryCache ClientCache
	redisCache    RedisCache
}

func NewClient(repo ClientRepo, inMemoryCache ClientCache, redisCache RedisCache) *Client {
	return &Client{
		repo:          repo,
		inMemoryCache: inMemoryCache,
		redisCache:    redisCache,
	}
}

func (c *Client) Create(ctx context.Context, client model.Client) error {
	client.UpdatedAt = time.Now().UTC()
	err := c.repo.Upsert(ctx, client, model.ClientFilter{ID: def.Pointer(client.ID)})
	if err != nil {
		return fmt.Errorf("c.repo.Upsert: %w", err)
	}

	c.inMemoryCache.Set(client)
	err = c.redisCache.Set(ctx, client)
	if err != nil {
		return fmt.Errorf("c.redisCache.Set: %w", err)
	}

	return nil
}

func (c *Client) Get(ctx context.Context, id uint64) (model.Client, error) {
	//client, exists := c.inMemoryCache.Get(id)
	//if !exists {
	//	return model.Client{}, fmt.Errorf("client with ID: %d not found", id)
	//}

	return c.redisCache.Get(ctx, id)
}

func (c *Client) List(ctx context.Context) ([]model.Client, error) {
	return c.repo.GetListWithFilter(ctx, model.ClientFilter{})
}
