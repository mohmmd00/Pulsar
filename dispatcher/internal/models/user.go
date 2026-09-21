package models

import (
	"time"
	"uuid"
)

type User struct {
	ID           uuid.UUID
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}

func NewUser(email string, passwordHash string) User {

	return User{
		ID:           uuid.New(),
		Email:        email,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now(),
	}
}
