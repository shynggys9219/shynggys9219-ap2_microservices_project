package nats

import (
	"context"
	"fmt"

	"github.com/nats-io/nats.go"
)

type MsgHandler func(msg *nats.Msg) error

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

func (nc *Client) CloseConnect() {
	nc.Conn.Close()
}
