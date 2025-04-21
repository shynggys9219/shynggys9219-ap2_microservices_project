package nats

import (
	"context"
	"fmt"

	"github.com/nats-io/nats.go"
)

type MsgHandler func(ctx context.Context, msg *nats.Msg) error

type Client struct {
	Conn *nats.Conn
}

func NewClient(ctx context.Context, hosts []string, nkey string, isTest bool) (*Client, error) {

	opts, err := setOptions(ctx, hosts, nkey, isTest)
	if err != nil {
		return nil, fmt.Errorf("setOptions: %w", err)
	}

	nc, err := opts.Connect()
	if err != nil {
		return nil, fmt.Errorf("opts.Connect: %w", err)
	}

	return &Client{
		Conn: nc,
	}, nil
}

func (c *Client) Subscribe(subject string, handler MsgHandler) (*nats.Subscription, error) {
	sub, err := c.Conn.Subscribe(subject, func(msg *nats.Msg) {
		ctx, cancel := context.WithTimeout(context.Background(), nats.DefaultTimeout)
		defer cancel()

		if err := handler(ctx, msg); err != nil {
			fmt.Printf("Error handling message from subject %s: %v\n", subject, err)
		}
	})
	if err != nil {
		return nil, fmt.Errorf("subscribe failed: %w", err)
	}

	return sub, nil
}

func (nc *Client) CloseConnect() {
	nc.Conn.Close()
}
