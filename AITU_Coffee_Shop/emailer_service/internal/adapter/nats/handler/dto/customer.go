package dto

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"

	events "github.com/shynggys9219/ap2-apis-gen-user-service/events/v1"
	"github.com/shynggys9219/ap2_microservices_project/emailer_service/internal/model"
)

func ToCustomer(msg *nats.Msg) (model.Customer, error) {
	var pbCustomer events.Customer
	err := proto.Unmarshal(msg.Data, &pbCustomer)
	if err != nil {
		return model.Customer{}, fmt.Errorf("proto.Unmarshall: %w", err)
	}

	return model.Customer{
		ID:        pbCustomer.Id,
		Name:      pbCustomer.Name,
		Phone:     pbCustomer.Phone,
		Email:     pbCustomer.Email,
		CreatedAt: pbCustomer.CreatedAt.AsTime(),
		UpdatedAt: pbCustomer.UpdatedAt.AsTime(),
		IsDeleted: pbCustomer.IsDeleted,
	}, nil
}
