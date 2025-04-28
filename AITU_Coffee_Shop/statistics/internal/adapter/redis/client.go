package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	goredis "github.com/redis/go-redis/v9"

	"github.com/shynggys9219/ap2_microservices_project/statistics/internal/model"
	"github.com/shynggys9219/ap2_microservices_project/statistics/pkg/redis"
)

const (
	keyPrefix = "client:%d"
)

// Client is our entity, client is redis client through which we are going to make queries to redis
type Client struct {
	client *redis.Client
	ttl    time.Duration
}

func NewClient(client *redis.Client, ttl time.Duration) *Client {
	return &Client{
		client: client,
		ttl:    ttl,
	}
}

func (c *Client) Set(ctx context.Context, client model.Client) error {
	data, err := json.Marshal(client)
	if err != nil {
		return fmt.Errorf("failed to marshal client: %w", err)
	}

	return c.client.Unwrap().Set(ctx, c.key(client.ID), data, c.ttl).Err()
}

func (c *Client) SetMany(ctx context.Context, clients []model.Client) error {
	pipe := c.client.Unwrap().Pipeline()
	for _, client := range clients {
		data, err := json.Marshal(client)
		if err != nil {
			return fmt.Errorf("failed to marshal client: %w", err)
		}
		pipe.Set(ctx, c.key(client.ID), data, c.ttl)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to set many clients: %w", err)
	}
	return nil
}

func (c *Client) Get(ctx context.Context, clientID uint64) (model.Client, error) {
	data, err := c.client.Unwrap().Get(ctx, c.key(clientID)).Bytes()
	if err != nil {
		if err == goredis.Nil {
			return model.Client{}, nil // not found
		}

		return model.Client{}, fmt.Errorf("failed to get client: %w", err)
	}

	var client model.Client
	err = json.Unmarshal(data, &client)
	if err != nil {
		return model.Client{}, fmt.Errorf("failed to unmarshal client: %w", err)
	}

	return client, nil
}

func (c *Client) Delete(ctx context.Context, clientID uint64) error {
	return c.client.Unwrap().Del(ctx, c.key(clientID)).Err()
}

func (c *Client) key(id uint64) string {
	return fmt.Sprintf(keyPrefix, id)
}
