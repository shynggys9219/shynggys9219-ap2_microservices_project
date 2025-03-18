package model

import "time"

type Client struct {
	ID           uint64
	Name         string
	Phone        string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	IsDeleted    bool
}
