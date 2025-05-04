package model

import "time"

type Customer struct {
	ID              uint64
	Name            string
	Phone           string
	Email           string
	CurrentPassword string
	NewPassword     string
	PasswordHash    string
	NewPasswordHash string
	CreatedAt       time.Time
	UpdatedAt       time.Time

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
