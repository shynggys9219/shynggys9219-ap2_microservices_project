package mongo

import (
	"context"
	"errors"
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
	daoClient := dao.FromClient(client)
	_, err := a.conn.Collection(a.collection).InsertOne(ctx, daoClient)
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
		return fmt.Errorf("client has not been updated with filter: %v, err: %w", filter, err)
	}

	if res.ModifiedCount == 0 {
		return fmt.Errorf("client has not been updated with filter: %v", filter)
	}

	return nil
}

func (a *Client) GetWithFilter(ctx context.Context, filter model.ClientFilter) (model.Client, error) {
	var clientDAO dao.Client

	err := a.conn.Collection(a.collection).FindOne(ctx, dao.FromClientFilter(filter)).Decode(&clientDAO)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.Client{}, model.ErrNotFound
		}

		return model.Client{}, fmt.Errorf("collection.FindOne: %w", err)
	}

	return dao.ToClient(clientDAO), nil
}
func (a *Client) GetListWithFilter(ctx context.Context, filter model.ClientFilter) ([]model.Client, error) {
	return nil, nil
}
