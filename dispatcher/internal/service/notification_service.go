package service

import (
	"context"
	"errors"
	"fmt"
	"uuid"

	"github.com/mohmmd00/Pulsar/dispatcher/internal/models"
	"github.com/mohmmd00/Pulsar/dispatcher/internal/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

func SubmitNotification(userID uuid.UUID, channel string, recipient string, message string, pool *pgxpool.Pool, ctx context.Context) (models.Notification, error) {
	errMsg := "Couldnt submit Notification"

	if userID == uuid.Nil() || recipient == "" || message == "" {
		return models.Notification{}, errors.New(errMsg + "invalid inputs")
	}

	newNotification := models.Notification{}

	switch channel {
	case "email", "sms", "push":
		newNotification = models.NewNotification(userID, channel, recipient, message)

		//repository layer
		err := repository.CreateNotification(newNotification, pool, ctx)
		if err != nil {
			return models.Notification{}, fmt.Errorf(errMsg+"%w", err)
		}
	default:
		return models.Notification{}, errors.New(errMsg + "invalid channel")
	}
	return newNotification, nil

}
