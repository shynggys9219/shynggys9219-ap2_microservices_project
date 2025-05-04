package dao

import (
	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type Customer struct {
	ID           uint64    `bson:"_id"`
	Name         string    `bson:"name"`
	Phone        string    `bson:"phone"`
	Email        string    `bson:"email"`
	PasswordHash string    `bson:"passwordHash"`
	CreatedAt    time.Time `bson:"createdAt"`
	UpdatedAt    time.Time `bson:"updatedAt"`

	IsDeleted bool `bson:"isDeleted"`
}

func FromCustomer(customer model.Customer) Customer {
	return Customer{
		ID:           customer.ID,
		Name:         customer.Name,
		Phone:        customer.Phone,
		Email:        customer.Email,
		PasswordHash: customer.PasswordHash,
		CreatedAt:    customer.CreatedAt,
		UpdatedAt:    customer.UpdatedAt,
		IsDeleted:    customer.IsDeleted,
	}
}

func ToCustomer(customer Customer) model.Customer {
	return model.Customer{
		ID:           customer.ID,
		Name:         customer.Name,
		Phone:        customer.Phone,
		Email:        customer.Email,
		PasswordHash: customer.PasswordHash,
		CreatedAt:    customer.CreatedAt,
		UpdatedAt:    customer.UpdatedAt,
		IsDeleted:    customer.IsDeleted,
	}
}

func FromCustomerFilter(filter model.CustomerFilter) bson.M {
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

func FromCustomerUpdateData(updateData model.CustomerUpdateData) bson.M {
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

	if updateData.PasswordHash != nil {
		query["passwordHash"] = *updateData.PasswordHash
	}

	if updateData.IsDeleted != nil {
		query["isDeleted"] = *updateData.IsDeleted
	}

	return bson.M{"$set": query}
}
