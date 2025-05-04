package dto

import (
	"github.com/shynggys9219/ap2-apis-gen-user-service/events/v1"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromCustomer(client model.Customer) *events.Customer {
	return &events.Customer{
		Id:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		Phone:     client.Phone,
		CreatedAt: timestamppb.New(client.CreatedAt),
		UpdatedAt: timestamppb.New(client.UpdatedAt),
		IsDeleted: client.IsDeleted,
	}
}
