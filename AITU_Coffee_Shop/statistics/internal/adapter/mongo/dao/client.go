package dao

import (
	"github.com/shynggys9219/ap2_microservices_project/statistics/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type Client struct {
	ID          uint64    `bson:"_id"`
	OrdersCount uint64    `bson:"ordersCount"`
	Orders      []*Order  `bson:"orders"`
	Phone       string    `bson:"phone"`
	CreatedAt   time.Time `bson:"createdAt"`
	UpdatedAt   time.Time `bson:"updatedAt"`
	IsDeleted   bool      `bson:"isDeleted"`
}

func FromClient(client model.Client) Client {
	orders := make([]*Order, 0)
	for i := range client.Orders {
		orders = append(orders, fromOrder(client.Orders[i]))
	}

	return Client{
		ID:          client.ID,
		OrdersCount: client.OrdersCount,
		Orders:      orders,
		Phone:       client.Phone,
		CreatedAt:   client.CreatedAt,
		UpdatedAt:   client.UpdatedAt,
		IsDeleted:   false,
	}
}

func ToClient(client Client) model.Client {
	orders := make([]*model.Order, 0)
	for i := range client.Orders {
		orders = append(orders, toOrder(client.Orders[i]))
	}

	return model.Client{
		ID:          client.ID,
		OrdersCount: client.OrdersCount,
		Orders:      orders,
		Phone:       client.Phone,
		CreatedAt:   client.CreatedAt,
		UpdatedAt:   client.UpdatedAt,
		IsDeleted:   client.IsDeleted,
	}
}

func FromClientFilter(filter model.ClientFilter) bson.M {
	query := bson.M{}

	if filter.ID != nil {
		query["_id"] = *filter.ID
	}

	return query
}

func FromClientUpdateData(updateData model.ClientUpdateData) bson.M {
	query := bson.M{}

	if updateData.Name != nil {
		query["name"] = *updateData.Name
	}

	if updateData.Phone != nil {
		query["phone"] = *updateData.Phone
	}

	if updateData.Email != nil {
		query["email"] = *updateData.Email
	}

	if updateData.IsDeleted != nil {
		query["isDeleted"] = *updateData.IsDeleted
	}

	return bson.M{"$set": query}
}
