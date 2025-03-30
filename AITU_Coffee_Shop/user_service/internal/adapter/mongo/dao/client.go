package dao

import (
	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type Client struct {
	ID           uint64    `bson:"_id"`
	Name         string    `bson:"name"`
	Phone        string    `bson:"phone"`
	Email        string    `bson:"email"`
	PasswordHash string    `bson:"passwordHash"`
	CreatedAt    time.Time `bson:"createdAt"`
	UpdatedAt    time.Time `bson:"updatedAt"`

	IsDeleted bool `bson:"isDeleted"`
}

func FromClient(client model.Client) Client {
	return Client{
		ID:           client.ID,
		Name:         client.Name,
		Phone:        client.Phone,
		Email:        client.Email,
		PasswordHash: client.PasswordHash,
		CreatedAt:    client.CreatedAt,
		UpdatedAt:    client.UpdatedAt,
		IsDeleted:    client.IsDeleted,
	}
}

func FromClientFilter(filter model.ClientFilter) bson.M {
	query := bson.M{}

	if filter.ID != nil {
		query["_id"] = *filter.ID
	}

	if filter.Name != nil {
		query["name"] = *filter.Name
	}

	if filter.Phone != nil {
		query["phone"] = *filter.Phone
	}

	if filter.Email != nil {
		query["email"] = *filter.Email
	}

	if filter.PasswordHash != nil {
		query["passwordHash"] = *filter.PasswordHash
	}

	if filter.IsDeleted != nil {
		query["isDeleted"] = *filter.IsDeleted
	}

	return query
}

func FromClientUpdateData(updateData model.ClientUpdateData) bson.M {
	query := bson.M{}

	if updateData.Name != nil {
		query["name"] = *updateData.Name
	}

	if updateData.Name != nil {
		query["name"] = *updateData.Name
	}

	if updateData.Phone != nil {
		query["phone"] = *updateData.Phone
	}

	if updateData.Email != nil {
		query["email"] = *updateData.Email
	}

	if updateData.PasswordHash != nil {
		query["passwordHash"] = *updateData.PasswordHash
	}

	if updateData.IsDeleted != nil {
		query["isDeleted"] = *updateData.IsDeleted
	}

	return bson.M{"$set": query}
}
