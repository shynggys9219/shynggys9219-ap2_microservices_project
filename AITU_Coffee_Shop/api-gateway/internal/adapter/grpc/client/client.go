package client

import (
	"context"
	svc "github.com/shynggys9219/ap2-apis-gen-user-service/service/frontend/client/v1"
	"github.com/shynggys9219/ap2_microservices_project/api-gateway/internal/adapter/grpc/client/dto"
	"github.com/shynggys9219/ap2_microservices_project/api-gateway/internal/model"
)

type Client struct {
	client svc.ClientServiceClient
}

func NewClient(client svc.ClientServiceClient) *Client {
	return &Client{
		client: client,
	}
}

func (c *Client) Create(ctx context.Context, request model.Client) (model.Client, error) {
	resp, err := c.client.Create(ctx, &svc.CreateRequest{
		Email:    request.Email,
		Password: request.CurrentPassword,
	})

	if err != nil {
		return model.Client{}, err
	}

	client := dto.FromGRPCClientCreateResponse(resp)

	return client, nil
}

func (c *Client) Update(ctx context.Context, request model.Client) (model.Client, error) {
	resp, err := c.client.Update(ctx, &svc.UpdateRequest{
		Id:          request.ID,
		Name:        request.Name,
		Email:       request.Email,
		Phone:       request.Phone,
		Password:    request.NewPassword,
		OldPassword: request.CurrentPassword,
	})

	if err != nil {
		return model.Client{}, err
	}

	return dto.FROMGRPCClientUpdateResponse(resp), nil
}

func (c *Client) Get(ctx context.Context, id uint64) (model.Client, error) {
	resp, err := c.client.Get(ctx, &svc.GetRequest{Id: id})
	if err != nil {
		return model.Client{}, err
	}

	return dto.FromGRPCClientGetResponse(resp), nil
}

func (c *Client) Delete(ctx context.Context, id uint64) error {
	return nil
}
