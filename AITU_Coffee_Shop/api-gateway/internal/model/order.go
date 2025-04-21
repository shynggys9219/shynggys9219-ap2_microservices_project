package model

import "time"

type OrderType int8
type OrderStatus int8

const (
	OrderTypeUnspecified OrderType = iota
	OrderTypePreOrder
	OrderTypeInstant
)

const (
	OrderStatusUnspecified OrderStatus = iota
	OrderStatusCreated
	OrderStatusPending
	OrderStatusCompleted
)

type Order struct {
	ID        uint64
	ClientID  uint64
	OrderType OrderType
	Sum       float32
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    OrderStatus
}
