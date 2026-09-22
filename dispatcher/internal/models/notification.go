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

func NewNotification(userId uuid.UUID, channel string, recipient string, message string) Notification {

	timeNow := time.Now() // now there is no difference between  createAt and UpdateAt

	return Notification{
		ID:        uuid.New(),
		UserID:    userId,
		Channel:   channel,
		Recipient: recipient,
		Message:   message,
		Status:    "pending",
		Attempts:  0,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}
}
