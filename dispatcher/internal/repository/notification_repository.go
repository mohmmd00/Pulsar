package repository

import (
	"context"

	"github.com/mohmmd00/Pulsar/dispatcher/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateNotification(note models.Notification, pool *pgxpool.Pool, ctx context.Context) error {

	_, err := pool.Exec(ctx, "INSERT INTO notifications (id , user_id , channel , recipient , message , status , attempts , created_at , updated_at) Values ($1,$2,$3,$4,$5,$6,$7,$8,$9)",
		note.ID, note.UserID, note.Channel, note.Recipient, note.Message, note.Status, note.Attempts, note.CreatedAt, note.UpdatedAt)

	if err != nil {
		return err
	}
	return nil
}
