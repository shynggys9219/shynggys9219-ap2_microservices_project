package dto

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
	"github.com/shynggys9219/ap2-apis-gen-user-service/events/v1"
	"github.com/shynggys9219/ap2_microservices_project/statistics/internal/model"
)

func ToClient(msg *nats.Msg) (model.Client, error) {
	var pbClient events.Client
	err := proto.Unmarshal(msg.Data, &pbClient)
	if err != nil {
		return model.Client{}, fmt.Errorf("proto.Unmarshall: %w", err)
	}

	return model.Client{
		ID:        pbClient.Id,
		Phone:     pbClient.Phone,
		CreatedAt: pbClient.CreatedAt.AsTime(),
		IsDeleted: pbClient.IsDeleted,
	}, nil
}
