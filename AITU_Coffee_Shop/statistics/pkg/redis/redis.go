package redis

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host         string
	Password     string
	TLSEnable    bool
	DialTimeout  time.Duration
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
}

type Client struct {
	client *redis.Client
}

func NewClient(ctx context.Context, cfg Config) (*Client, error) {
	// TODO: add telemetry later

	opts := &redis.Options{
		Addr:         cfg.Host,
		DialTimeout:  cfg.DialTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		Password:     cfg.Password,
	}

	if cfg.TLSEnable {
		opts.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}

	client := redis.NewClient(opts)

	err := client.Ping(ctx).Err()
	if err != nil {
		return nil, fmt.Errorf("ping: %w", err)
	}

	return &Client{client: client}, nil
}

func (c *Client) Close() error {
	err := c.client.Close()
	if err != nil {
		return fmt.Errorf("close: %w", err)
	}

	return nil
}

func (c *Client) Ping(ctx context.Context) error {
	err := c.client.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("redis: %w", err)
	}

	return nil
}

func (c *Client) Unwrap() *redis.Client {
	return c.client
}
