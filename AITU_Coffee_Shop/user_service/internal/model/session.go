package model

import "time"

type Session struct {
	UserID       uint64
	RefreshToken string
	ExpiresAt    time.Time
	CreatedAt    time.Time
}
