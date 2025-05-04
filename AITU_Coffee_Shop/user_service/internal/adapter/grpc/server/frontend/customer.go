package frontend

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	svc "github.com/shynggys9219/ap2-apis-gen-user-service/service/frontend/customer/v1"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/adapter/grpc/server/frontend/dto"
)

type Customer struct {
	svc.UnimplementedCustomerServiceServer

	customerUsecase CustomerUsecase
}

func NewCustomer(uc CustomerUsecase) *Customer {
	return &Customer{
		customerUsecase: uc,
	}
}

func (c *Customer) Register(ctx context.Context, req *svc.RegisterRequest) (*svc.RegisterResponse, error) {
	//TODO: validate request

	customer, err := dto.ToCustomerFromRegisterRequest(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	id, err := c.customerUsecase.Register(ctx, customer)
	if err != nil {
		return nil, dto.FromError(err)
	}

	return &svc.RegisterResponse{Id: id}, nil
}

func (c *Customer) Update(ctx context.Context, req *svc.UpdateRequest) (*svc.UpdateResponse, error) {
	//TODO: validate request

	client, err := dto.ToCustomerFromUpdateRequest(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	newClient, err := c.customerUsecase.Update(ctx, client)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &svc.UpdateResponse{
		Customer: dto.FromCustomer(newClient),
	}, nil
}

func (c *Customer) Get(ctx context.Context, req *svc.GetRequest) (*svc.GetResponse, error) {
	if req.Id <= 0 {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("wrong ID: %d", req.Id))
	}
	client, err := c.customerUsecase.Get(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &svc.GetResponse{
		Customer: dto.FromCustomer(client),
	}, nil
}

func (c *Customer) Login(ctx context.Context, req *svc.LoginRequest) (*svc.LoginResponse, error) {

	return &svc.LoginResponse{
		AccessToken:  "",
		RefreshToken: "",
	}, nil
}

func (c *Customer) RefreshToken(
	ctx context.Context, req *svc.RefreshTokenRequest,
) (*svc.RefreshTokenResponse, error) {

	return &svc.RefreshTokenResponse{
		AccessToken:  "",
		RefreshToken: "",
	}, nil
}
