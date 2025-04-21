package model

import "time"

type Client struct {
	ID              uint64
	Name            string
	Phone           string
	Email           string
	CurrentPassword string
	NewPassword     string
	CreatedAt       time.Time
	UpdatedAt       time.Time

	IsDeleted bool
}

type ClientStatistic struct {
	ID          uint64
	OrdersCount uint64
	Orders      []*Order
	Phone       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsDeleted   bool
}
