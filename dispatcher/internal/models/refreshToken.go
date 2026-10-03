package models

import (
	"time"
	"uuid"
)

type RefreshToken struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	TokenHash string
	ExpiresAt time.Time
	Revoked   bool
	CreatedAt time.Time
}

func NewRefreshToken(userID uuid.UUID, tokenHash string, expiresAt time.Time, revoked bool) RefreshToken {

	return RefreshToken{
		ID:        uuid.New(),
		UserID:    userID,
		TokenHash: tokenHash,
		ExpiresAt: expiresAt,
		Revoked:   revoked,
		CreatedAt: time.Now(),
	}
}
