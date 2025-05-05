package frontend

import (
	"context"
	"fmt"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/pkg/security"

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

	token, ok := security.TokenFromCtx(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}

	newClient, err := c.customerUsecase.Update(ctx, token, client)
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

	token, ok := security.TokenFromCtx(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}

	client, err := c.customerUsecase.Get(ctx, token, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &svc.GetResponse{
		Customer: dto.FromCustomer(client),
	}, nil
}

func (c *Customer) Login(ctx context.Context, req *svc.LoginRequest) (*svc.LoginResponse, error) {
	if req.Email == "" || req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "email or password not provided")
	}

	token, err := c.customerUsecase.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &svc.LoginResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}

func (c *Customer) RefreshToken(
	ctx context.Context, req *svc.RefreshTokenRequest,
) (*svc.RefreshTokenResponse, error) {
	if req.RefreshToken == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid refresh token")
	}

	token, err := c.customerUsecase.RefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &svc.RefreshTokenResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}
