package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Ai struct {
	collection *mongo.Collection
}

const collectionAi = "ai"

func NewAi(conn *mongo.Database) *Ai {
	return &Ai{
		collection: conn.Collection(collectionAi),
	}
}

func (ai *Ai) Next(ctx context.Context, name string) (uint64, error) {
	result := struct {
		ID      string `bson:"_id"`
		Counter uint64 `bson:"counter"`
	}{}

	filter := bson.D{{Key: "_id", Value: name}}
	update := bson.D{{Key: "$inc", Value: bson.D{{Key: "counter", Value: 1}}}}
	after := options.After
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(after)
	err := ai.collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&result)
	if err != nil {
		return 0, fmt.Errorf("FindOneAndUpdate: %w", err)
	}

	return result.Counter, nil
}
