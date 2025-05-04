package dto

import (
	base "github.com/shynggys9219/ap2-apis-gen-user-service/base/frontend/v1"
	svc "github.com/shynggys9219/ap2-apis-gen-user-service/service/frontend/customer/v1"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToCustomerFromRegisterRequest(req *svc.RegisterRequest) (model.Customer, error) {
	return model.Customer{
		Email:           req.Email,
		CurrentPassword: req.Password,
	}, nil
}

func ToCustomerFromUpdateRequest(req *svc.UpdateRequest) (model.Customer, error) {
	return model.Customer{
		ID:              req.Id,
		Name:            req.Name,
		Phone:           req.Phone,
		Email:           req.Email,
		CurrentPassword: req.OldPassword,
	}, nil
}

func FromCustomer(client model.Customer) *base.Customer {
	return &base.Customer{
		Id:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		Phone:     client.Phone,
		CreatedAt: timestamppb.New(client.CreatedAt),
		UpdatedAt: timestamppb.New(client.UpdatedAt),
		IsDeleted: client.IsDeleted,
	}
}
