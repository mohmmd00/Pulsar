package models

import (
	"time"
	"uuid"
)

type Notification struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Channel   string
	Recipient string
	Message   string
	Status    string
	Attempts  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
