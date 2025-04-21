package dto

import (
	svc "github.com/shynggys9219/ap2-apis-gen-user-service/service/frontend/client/v1"
	"github.com/shynggys9219/ap2_microservices_project/api-gateway/internal/model"
)

func FromGRPCClientCreateResponse(resp *svc.CreateResponse) model.Client {
	return model.Client{ID: resp.Id}
}

func FromGRPCClientGetResponse(resp *svc.GetResponse) model.Client {
	return model.Client{
		ID:        resp.Client.Id,
		Name:      resp.Client.Name,
		Phone:     resp.Client.Phone,
		Email:     resp.Client.Email,
		CreatedAt: resp.Client.CreatedAt.AsTime(),
		IsDeleted: resp.Client.IsDeleted,
	}
}

func FROMGRPCClientUpdateResponse(resp *svc.UpdateResponse) model.Client {
	return model.Client{
		ID:        resp.Client.Id,
		Name:      resp.Client.Name,
		Phone:     resp.Client.Phone,
		Email:     resp.Client.Email,
		CreatedAt: resp.Client.CreatedAt.AsTime(),
		IsDeleted: resp.Client.IsDeleted,
	}
}
