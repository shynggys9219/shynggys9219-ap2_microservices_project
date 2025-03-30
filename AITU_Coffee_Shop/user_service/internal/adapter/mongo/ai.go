package mongo

import (
	"context"
	"fmt"
	"log"

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

func (a *Ai) Next(ctx context.Context, coll string) (uint64, error) {
	log.Println("starting mongo session for collection:", collectionAi)
	session, err := a.collection.Database().Client().StartSession()
	if err != nil {
		return 0, fmt.Errorf("StartSession: %w", err)
	}

	defer session.EndSession(ctx)

	result := struct {
		ID      string `bson:"_id"`
		Counter uint64 `bson:"counter"`
	}{}

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		filter := bson.D{{Key: "_id", Value: coll}}
		update := bson.D{{Key: "$inc", Value: bson.D{{Key: "counter", Value: 1}}}}
		after := options.After
		opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(after)
		err = a.collection.FindOneAndUpdate(sessCtx, filter, update, opts).Decode(&result)
		if err != nil {
			return nil, fmt.Errorf("FindOneAndUpdate: %w", err)
		}

		return result.Counter, nil
	}
	_, err = session.WithTransaction(ctx, callback)
	if err != nil {
		return 0, fmt.Errorf("WithTransaction: %w", err)
	}

	return result.Counter, nil
}
