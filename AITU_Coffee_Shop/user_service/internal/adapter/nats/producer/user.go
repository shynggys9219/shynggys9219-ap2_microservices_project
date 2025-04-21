package producer

import (
	"context"
	"fmt"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/adapter/grpc/server/frontend/dto"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/model"
	"google.golang.org/protobuf/proto"
	"log"
	"time"

	"github.com/shynggys9219/ap2_microservices_project/user_svc/pkg/nats"
)

const PushTimeout = time.Second * 30

type Client struct {
	client  *nats.Client
	subject string
}

func NewClientProducer(
	client *nats.Client,
	subject string,
) *Client {
	return &Client{
		client:  client,
		subject: subject,
	}
}

func (c *Client) Push(ctx context.Context, client model.Client) error {
	ctx, cancel := context.WithTimeout(ctx, PushTimeout)
	defer cancel()

	clientPb := dto.FromClient(client)
	data, err := proto.Marshal(clientPb)
	if err != nil {
		return fmt.Errorf("proto.Marshal: %w", err)
	}

	err = c.client.Conn.Publish(c.subject, data)
	if err != nil {
		return fmt.Errorf("c.client.Conn.Publish: %w", err)
	}
	log.Println("client is pushed:", client)

	return nil
}
