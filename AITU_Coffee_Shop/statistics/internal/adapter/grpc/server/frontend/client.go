package frontend

import (
	"context"
	base "github.com/shynggys9219/ap2-apis-gen-statistics-service/base/frontend/v1"
	"github.com/shynggys9219/ap2_microservices_project/statistics/internal/adapter/grpc/server/frontend/dto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	svc "github.com/shynggys9219/ap2-apis-gen-statistics-service/service/frontend/client_stats/v1"
)

type Client struct {
	svc.UnimplementedClientStatisticsServiceServer

	uc ClientUsecase
}

func NewClient(uc ClientUsecase) *Client {
	return &Client{
		uc: uc,
	}
}

func (c *Client) Get(ctx context.Context, req *svc.GetRequest) (*svc.GetResponse, error) {
	//TODO: validate request

	if req.Id < 1 {
		return nil, status.Error(codes.InvalidArgument, "invalid id")
	}

	client, err := c.uc.Get(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &svc.GetResponse{Client: dto.FromClientToStatistic(client)}, nil
}

func (c *Client) List(ctx context.Context, _ *svc.ListRequest) (*svc.ListResponse, error) {
	// TODO: check if user is authorized or not
	clientStats, err := c.uc.List(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	pbClientsStatistic := make([]*base.Statistic, len(clientStats))
	for i := range clientStats {
		pbClientsStatistic = append(pbClientsStatistic, dto.FromClientToStatistic(clientStats[i]))
	}

	return &svc.ListResponse{
		Clients: pbClientsStatistic,
	}, nil
}
