package dto

import (
	base "github.com/shynggys9219/ap2-apis-gen-user-service/base/frontend/v1"
	svc "github.com/shynggys9219/ap2-apis-gen-user-service/service/frontend/client/v1"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToClientFromCreateRequest(req *svc.CreateRequest) (model.Client, error) {
	return model.Client{
		Email:           req.Email,
		CurrentPassword: req.Password,
	}, nil
}

func ToClientFromUpdateRequest(req *svc.UpdateRequest) (model.Client, error) {
	return model.Client{
		ID:              req.Id,
		Name:            req.Name,
		Phone:           req.Phone,
		Email:           req.Email,
		CurrentPassword: req.OldPassword,
	}, nil
}

func FromClient(client model.Client) *base.Client {
	return &base.Client{
		Id:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		Phone:     client.Phone,
		CreatedAt: timestamppb.New(client.CreatedAt),
		IsDeleted: client.IsDeleted,
	}
}
