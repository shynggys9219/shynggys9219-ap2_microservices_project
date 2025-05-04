package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type Transactor struct {
	dbClient *mongo.Client
}

func NewTransactor(dbClient *mongo.Client) *Transactor {
	return &Transactor{
		dbClient: dbClient,
	}
}

// WithinTransaction starts new mongo session
// and wraps callback within necessary SessionContext and executes it in transaction
func (t *Transactor) WithinTransaction(ctx context.Context, fn func(fnCtx context.Context) error) error {
	session, err := t.dbClient.StartSession()
	if err != nil {
		return fmt.Errorf("StartSession: %w", err)
	}
	defer session.EndSession(ctx)

	_, err = session.WithTransaction(ctx, func(sessionContext mongo.SessionContext) (interface{}, error) {
		return nil, fn(sessionContext)
	})
	if err != nil {
		return fmt.Errorf("failed to make transaction: %w", err)
	}

	return nil
}
