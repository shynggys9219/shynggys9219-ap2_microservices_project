package model

import "time"

const (
	CustomerRole = "customer"
)

type Customer struct {
	ID        uint64
	Name      string
	Phone     string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time

	IsDeleted bool
}

type CustomerFilter struct {
	ID           *uint64
	Name         *string
	Phone        *string
	Email        *string
	PasswordHash *string

	IsDeleted *bool
}

type CustomerUpdateData struct {
	ID           *uint64
	Name         *string
	Phone        *string
	Email        *string
	PasswordHash *string
	UpdatedAt    *time.Time

	IsDeleted *bool
}
