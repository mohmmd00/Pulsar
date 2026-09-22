package models

import "uuid"

type NotificationEvent struct {
	NotificationID uuid.UUID `json:"notification_id"`
	Channel        string    `json:"channel"`
}

func NewNotificationEvent(notificationId uuid.UUID, channel string) NotificationEvent {
	return NotificationEvent{
		NotificationID: notificationId,
		Channel:        channel,
	}
}
