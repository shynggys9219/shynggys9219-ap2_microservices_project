package model

import "time"

type Employee struct {
	ID           uint64
	FirstName    string
	SecondName   string
	Position     Position
	Email        string
	Phone        string
	PasswordHash string
	HiredBy      uint64 // user id that hired an employee
	HiredAt      time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

type Position struct {
	Name          string
	Salary        string
	DirectManager uint64   // a manager id
	Subordinates  []uint64 // ids of colleagues
	PromotedAt    time.Time
}

func (e *Employee) IsDeleted() bool {
	// if !=nil employee is deleted
	return e.DeletedAt != nil
}
