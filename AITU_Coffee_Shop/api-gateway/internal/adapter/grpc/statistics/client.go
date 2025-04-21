package statistics

import (
	"context"
	
	svc "github.com/shynggys9219/ap2-apis-gen-statistics-service/service/frontend/client_stats/v1"
	"github.com/shynggys9219/ap2_microservices_project/api-gateway/internal/adapter/grpc/statistics/dto"
	"github.com/shynggys9219/ap2_microservices_project/api-gateway/internal/model"
)

type Client struct {
	client svc.ClientStatisticsServiceClient
}

func NewClient(client svc.ClientStatisticsServiceClient) *Client {
	return &Client{
		client: client,
	}
}

func (c *Client) List(ctx context.Context) ([]model.ClientStatistic, error) {
	resp, err := c.client.List(ctx, &svc.ListRequest{})

	if err != nil {
		return nil, err
	}

	return dto.FROMGRPCClientListResponse(resp), nil
}

func (c *Client) Get(ctx context.Context, id uint64) (model.ClientStatistic, error) {
	resp, err := c.client.Get(ctx, &svc.GetRequest{Id: id})
	if err != nil {
		return model.ClientStatistic{}, err
	}

	return dto.FromGRPCClientGetResponse(resp), nil
}
