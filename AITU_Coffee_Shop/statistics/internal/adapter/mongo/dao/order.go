package dao

import (
	"github.com/shynggys9219/ap2_microservices_project/statistics/internal/model"
	"time"
)

type Order struct {
	ID        uint64            `bson:"_id"`
	ClientID  uint64            `bson:"clientID"`
	OrderType model.OrderType   `bson:"orderType"`
	Sum       float32           `bson:"sum"`
	CreatedAt time.Time         `bson:"createdAt"`
	UpdatedAt time.Time         `bson:"updatedAt"`
	Status    model.OrderStatus `bson:"status"`
}

func fromOrder(order *model.Order) *Order {
	return &Order{
		ID:        order.ID,
		ClientID:  order.ClientID,
		OrderType: order.OrderType,
		Sum:       order.Sum,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
		Status:    order.Status,
	}
}

func toOrder(order *Order) *model.Order {
	return &model.Order{
		ID:        order.ID,
		ClientID:  order.ClientID,
		OrderType: order.OrderType,
		Sum:       order.Sum,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
		Status:    order.Status,
	}
}
