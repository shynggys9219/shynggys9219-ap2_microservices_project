package mongo

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/shynggys9219/ap2_microservices_project/statistics/internal/adapter/mongo/dao"
	"github.com/shynggys9219/ap2_microservices_project/statistics/internal/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	conn       *mongo.Database
	collection string
}

const (
	collectionClientStatistics = "clientStatistics"
)

func NewClient(conn *mongo.Database) *Client {
	return &Client{
		conn:       conn,
		collection: collectionClientStatistics,
	}
}

func (c *Client) Upsert(ctx context.Context, client model.Client, filter model.ClientFilter) error {
	update := bson.M{"$set": dao.FromClient(client)}
	opts := options.Update().SetUpsert(true)

	_, err := c.conn.Collection(collectionClientStatistics).UpdateOne(ctx, dao.FromClientFilter(filter), update, opts)
	if err != nil {
		return fmt.Errorf("c.conn.Collection(collectionClientStatistics).UpdateOne: %w", err)
	}

	return nil
}

func (c *Client) GetWithFilter(ctx context.Context, filter model.ClientFilter) (model.Client, error) {
	var clientDAO dao.Client

	err := c.conn.Collection(c.collection).FindOne(ctx, dao.FromClientFilter(filter)).Decode(&clientDAO)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.Client{}, model.ErrNotFound
		}

		return model.Client{}, fmt.Errorf("collection.FindOne: %w", err)
	}

	return dao.ToClient(clientDAO), nil
}

func (c *Client) GetListWithFilter(ctx context.Context, filter model.ClientFilter) ([]model.Client, error) {
	cur, err := c.conn.Collection(c.collection).Find(ctx, dao.FromClientFilter(filter))
	if err != nil {
		return nil, fmt.Errorf("collection.Find: %w", err)
	}
	defer cur.Close(ctx)

	var clients []model.Client
	for cur.Next(ctx) {
		var clientDAO dao.Client
		err = cur.Decode(&clientDAO)
		if err != nil {
			return nil, fmt.Errorf("cursor.Decode: %w", err)
		}
		clients = append(clients, dao.ToClient(clientDAO))
	}
	err = cur.Err()
	if err != nil {
		return nil, fmt.Errorf("cursor iteration error: %w", err)
	}

	if len(clients) == 0 {
		return nil, model.ErrNotFound
	}

	return clients, nil
}
