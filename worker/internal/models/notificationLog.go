package models

import (
	"time"
	"uuid"
)

type NotificationLog struct {
	ID             uuid.UUID
	NotificationID uuid.UUID
	Result         string
	ErrorMessage   *string
	CreatedAt      time.Time
}

func NewNotificationLog(notificationID uuid.UUID, result string, errorMessage string) NotificationLog {

	var errMsgPtr *string
	if errorMessage != "" {
		errMsgPtr = &errorMessage
	}

	return NotificationLog{
		ID:             uuid.New(),
		NotificationID: notificationID,
		Result:         result,
		ErrorMessage:   errMsgPtr, //must set null in db
		CreatedAt:      time.Now(),
	}
}
