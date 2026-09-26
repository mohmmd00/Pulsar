package repository

import (
	"context"
	"time"
	"uuid"

	"github.com/mohmmd00/Pulsar/worker/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetNotificationByID(noteID uuid.UUID, pool *pgxpool.Pool, ctx context.Context) (note models.Notification, err error) {

	row := pool.QueryRow(ctx, "SELECT id,user_id,channel,recipient,message,status,attempts,created_at,updated_at FROM notifications WHERE id = $1", noteID)
	scanErr := row.Scan(&note.ID, &note.UserID, &note.Channel, &note.Recipient, &note.Message, &note.Status, &note.Attempts, &note.CreatedAt, &note.UpdatedAt) //scan needs pointer
	if scanErr != nil {
		return models.Notification{}, scanErr
	}
	return note, nil
}

func UpdateNotificationStatus(noteID uuid.UUID, status string, incrementAttempts bool, pool *pgxpool.Pool, ctx context.Context) error { //increment attemps as the numbers this functions called

	if incrementAttempts {
		_, execErr := pool.Exec(ctx, "UPDATE notifications SET status = $1 ,updated_at = $2 , attempts = attempts + 1  WHERE id = $3", status, time.Now(), noteID)
		return execErr
	}
	_, execErr := pool.Exec(ctx, "UPDATE notifications SET status = $1, updated_at = $2 WHERE id = $3", status, time.Now(), noteID)
	return execErr
}

func CreateNotificationLog(notiflog models.NotificationLog, pool *pgxpool.Pool, ctx context.Context) error {
	_, err := pool.Exec(ctx, "INSERT INTO notification_logs (id , notification_id , result , error_message , created_at) VALUES ($1,$2,$3,$4,$5)",
		notiflog.ID, notiflog.NotificationID, notiflog.Result, notiflog.ErrorMessage, notiflog.CreatedAt)

	return err
}
