package mongo

import (
	"context"
	"fmt"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/adapter/mongo/dao"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	conn       *mongo.Database
	collection string
}

const (
	collectionUsers = "clients"
)

func NewClient(conn *mongo.Database) *Client {
	return &Client{
		conn:       conn,
		collection: collectionUsers,
	}
}

func (a *Client) Create(ctx context.Context, client model.Client) error {
	achievement := dao.FromClient(client)
	_, err := a.conn.Collection(a.collection).InsertOne(ctx, achievement)
	if err != nil {
		return fmt.Errorf("client with ID %d has not been created: %w", client.ID, err)
	}

	return nil
}

func (a *Client) Update(ctx context.Context, filter model.ClientFilter, update model.ClientUpdateData) error {
	res, err := a.conn.Collection(a.collection).UpdateOne(
		ctx,
		dao.FromClientFilter(filter),
		dao.FromClientUpdateData(update),
	)
	if err != nil {
		return fmt.Errorf("achievement has not been updated with filter: %v, err: %w", filter, err)
	}

	if res.ModifiedCount == 0 {
		return fmt.Errorf("achievement has not been updated with filter: %v", filter)
	}

	return nil
}

func (a *Client) GetWithFilter(ctx context.Context, filter model.ClientFilter) (model.Client, error) {
	return model.Client{}, nil
}
func (a *Client) GetListWithFilter(ctx context.Context, filter model.ClientFilter) ([]model.Client, error) {
	return nil, nil
}
