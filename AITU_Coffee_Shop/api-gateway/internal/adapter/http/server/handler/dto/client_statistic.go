package dto

import (
	"github.com/shynggys9219/ap2_microservices_project/api-gateway/internal/model"
	"time"
)

type ClientStatistic struct {
	ID          uint64    `json:"id"`
	OrdersCount uint64    `json:"ordersCount"`
	Orders      []*Order  `json:"orders"`
	Phone       string    `json:"phone"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"UpdatedAt"`
	IsDeleted   bool      `json:"isDeleted"`
}

func FromClientStatistic(client model.ClientStatistic) ClientStatistic {
	return ClientStatistic{
		ID:          client.ID,
		OrdersCount: client.OrdersCount,
		Orders:      toClientOrders(client.Orders),
		Phone:       client.Phone,
		CreatedAt:   client.CreatedAt,
		UpdatedAt:   client.UpdatedAt,
		IsDeleted:   client.IsDeleted,
	}
}

func FromClientStatisticToList(stats []model.ClientStatistic) []ClientStatistic {
	list := make([]ClientStatistic, 0)
	for i := range stats {
		list = append(list, ClientStatistic{
			ID:          stats[i].ID,
			OrdersCount: stats[i].OrdersCount,
			Orders:      toClientOrders(stats[i].Orders),
			Phone:       stats[i].Phone,
			CreatedAt:   stats[i].CreatedAt,
			UpdatedAt:   stats[i].UpdatedAt,
			IsDeleted:   stats[i].IsDeleted,
		})
	}

	return list
}

func toClientOrders(orders []*model.Order) []*Order {
	ordersJSON := make([]*Order, 0)
	for i := range orders {
		ordersJSON = append(ordersJSON, toClientOrder(orders[i]))
	}

	return ordersJSON
}

func toClientOrder(order *model.Order) *Order {
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
