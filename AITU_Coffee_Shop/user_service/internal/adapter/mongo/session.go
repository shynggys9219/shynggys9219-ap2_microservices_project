package mongo

import (
	"context"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/adapter/mongo/dao"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RefreshToken struct {
	conn       *mongo.Database
	collection string
}

const (
	collectionRefreshTokens = "refreshTokens"
)

func NewRefreshToken(conn *mongo.Database) *RefreshToken {
	return &RefreshToken{
		conn:       conn,
		collection: collectionRefreshTokens,
	}
}

func (r *RefreshToken) Create(ctx context.Context, session model.Session) error {
	_, err := r.conn.Collection(r.collection).InsertOne(ctx, dao.FromSession(session))

	return err
}

func (r *RefreshToken) GetByToken(ctx context.Context, token string) (model.Session, error) {
	var session dao.Session
	err := r.conn.Collection(r.collection).FindOne(ctx, bson.M{"refreshToken": token}).Decode(&session)
	if err != nil {
		return model.Session{}, err
	}

	return dao.ToSession(session), nil
}

func (r *RefreshToken) DeleteByToken(ctx context.Context, token string) error {
	_, err := r.conn.Collection(r.collection).DeleteOne(ctx, bson.M{"refreshToken": token})

	return err
}
