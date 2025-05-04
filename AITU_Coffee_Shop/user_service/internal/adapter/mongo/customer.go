package mongo

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/adapter/mongo/dao"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/model"
)

type Customer struct {
	conn       *mongo.Database
	collection string
}

const (
	collectionCustomers = "customers"
)

func NewCustomer(conn *mongo.Database) *Customer {
	return &Customer{
		conn:       conn,
		collection: collectionCustomers,
	}
}

func (a *Customer) EnsureIndexes(ctx context.Context) error {
	indexModel := mongo.IndexModel{
		Keys: bson.M{"email": 1},
		Options: options.Index().
			SetUnique(true).
			SetName("unique_email"),
	}

	_, err := a.conn.Collection(a.collection).Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return fmt.Errorf("failed to create unique index on email: %w", err)
	}

	return nil
}

func (a *Customer) Create(ctx context.Context, customer model.Customer) error {
	_, err := a.conn.Collection(a.collection).InsertOne(ctx, dao.FromCustomer(customer))
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return model.ErrEmailAlreadyRegistered
		}

		return fmt.Errorf("customer with ID %d has not been created: %w", customer.ID, err)
	}

	return nil
}

func (a *Customer) Update(ctx context.Context, filter model.CustomerFilter, update model.CustomerUpdateData) error {
	_, err := a.conn.Collection(a.collection).UpdateOne(
		ctx,
		dao.FromCustomerFilter(filter),
		dao.FromCustomerUpdateData(update),
	)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return model.ErrEmailAlreadyRegistered
		}

		return fmt.Errorf("customer has not been updated with filter: %v, err: %w", filter, err)
	}

	return nil
}

func (a *Customer) GetWithFilter(ctx context.Context, filter model.CustomerFilter) (model.Customer, error) {
	var customerDAO dao.Customer

	err := a.conn.Collection(a.collection).FindOne(ctx, dao.FromCustomerFilter(filter)).Decode(&customerDAO)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.Customer{}, model.ErrNotFound
		}

		return model.Customer{}, fmt.Errorf("collection.FindOne: %w", err)
	}

	return dao.ToCustomer(customerDAO), nil
}
func (a *Customer) GetListWithFilter(ctx context.Context, filter model.CustomerFilter) ([]model.Customer, error) {
	return nil, nil
}
