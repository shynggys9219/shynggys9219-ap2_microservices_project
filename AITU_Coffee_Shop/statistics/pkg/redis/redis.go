package redis

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Hosts        []string
	Password     string
	TLSEnable    bool
	DialTimeout  time.Duration
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
}

type Client struct {
	client *redis.ClusterClient
}

func NewClient(ctx context.Context, cfg Config) (*Client, error) {
	// TODO: add telemetry later

	clusterOptions := &redis.ClusterOptions{
		Addrs:        cfg.Hosts,
		DialTimeout:  cfg.DialTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		Password:     cfg.Password,
	}

	if cfg.TLSEnable {
		clusterOptions.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}

	client := redis.NewClusterClient(clusterOptions)

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

func (c *Client) Unwrap() *redis.ClusterClient {
	return c.client
}
