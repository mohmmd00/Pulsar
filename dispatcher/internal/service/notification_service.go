package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"uuid"

	"github.com/mohmmd00/Pulsar/dispatcher/internal/models"
	"github.com/mohmmd00/Pulsar/dispatcher/internal/repository"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nats-io/nats.go"
)

func SubmitNotification(userID uuid.UUID, channel string, recipient string, message string, pool *pgxpool.Pool, natsConn *nats.Conn, ctx context.Context) (models.Notification, error) {
	errMsg := "Couldnt submit Notification"

	if userID == uuid.Nil() || recipient == "" || message == "" {
		return models.Notification{}, errors.New(errMsg + "invalid inputs")
	}

	newNotification := models.Notification{}

	switch channel {
	case "email", "sms", "push":
		newNotification = models.NewNotification(userID, channel, recipient, message)

		//repository layer
		crtNoteErr := repository.CreateNotification(newNotification, pool, ctx)
		if crtNoteErr != nil {
			return models.Notification{}, fmt.Errorf(errMsg+"%w", crtNoteErr)
		}

		//nats
		event := models.NewNotificationEvent(newNotification.ID, newNotification.Channel)

		marshEvent, marshErr := json.Marshal(event)
		if marshErr != nil {
			return newNotification, fmt.Errorf(errMsg+"%w", marshErr)
		}

		pubErr := natsConn.Publish("notifications."+channel, marshEvent)
		if pubErr != nil {
			return newNotification, fmt.Errorf(errMsg+"%w", pubErr)
		}

	default:
		return models.Notification{}, errors.New(errMsg + "invalid channel")
	}
	return newNotification, nil

}
