package model

import "time"

type Client struct {
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

type ClientFilter struct {
	ID           *uint64
	Name         *string
	Phone        *string
	Email        *string
	PasswordHash *string

	IsDeleted *bool
}

type ClientUpdateData struct {
	ID           *uint64
	Name         *string
	Phone        *string
	Email        *string
	PasswordHash *string
	UpdatedAt    *time.Time

	IsDeleted *bool
}
