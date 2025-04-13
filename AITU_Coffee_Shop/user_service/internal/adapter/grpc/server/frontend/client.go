package frontend

import (
	"context"
	"fmt"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/adapter/grpc/server/frontend/dto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	svc "github.com/shynggys9219/ap2-apis-gen-user-service/service/frontend/client/v1"
)

type Client struct {
	svc.UnimplementedClientServiceServer

	uc ClientUsecase
}

func NewClient(uc ClientUsecase) *Client {
	return &Client{
		uc: uc,
	}
}

func (c *Client) Create(ctx context.Context, req *svc.CreateRequest) (*svc.CreateResponse, error) {
	//TODO: validate request

	client, err := dto.ToClientFromCreateRequest(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	newClient, err := c.uc.Create(ctx, client)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &svc.CreateResponse{Id: newClient.ID}, nil
}

func (c *Client) Update(ctx context.Context, req *svc.UpdateRequest) (*svc.UpdateResponse, error) {
	//TODO: validate request

	client, err := dto.ToClientFromUpdateRequest(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	newClient, err := c.uc.Update(ctx, client)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &svc.UpdateResponse{
		Client: dto.FromClient(newClient),
	}, nil
}

func (c *Client) Get(ctx context.Context, req *svc.GetRequest) (*svc.GetResponse, error) {
	if req.Id <= 0 {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("wrong ID: %d", req.Id))
	}
	client, err := c.uc.Get(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &svc.GetResponse{
		Client: dto.FromClient(client),
	}, nil
}
