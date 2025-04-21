package model

import (
	"time"
)

type Client struct {
	ID          uint64
	OrdersCount uint64
	Orders      []*Order
	Phone       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsDeleted   bool
}

type ClientFilter struct {
	ID    *uint64
	Name  *string
	Phone *string
	Email *string

	IsDeleted *bool
}

type ClientUpdateData struct {
	ID        *uint64
	Name      *string
	Phone     *string
	Email     *string
	UpdatedAt *time.Time

	IsDeleted *bool
}
