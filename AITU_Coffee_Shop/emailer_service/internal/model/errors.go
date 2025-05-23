package model

import "errors"

var (
	ErrInvalidEmail           = errors.New("invalid email")
	ErrInvalidPassword        = errors.New("invalid password")
	ErrInvalidID              = errors.New("invalid id")
	ErrNotFound               = errors.New("not found")
	ErrEmailAlreadyRegistered = errors.New("email already registered")
	ErrRefreshTokenExpired    = errors.New("refresh token is expired")
)
