package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mohmmd00/Pulsar/dispatcher/internal/models"
)

func CreateUser(baseUser models.User, pool *pgxpool.Pool, ctx context.Context) error {

	_, err := pool.Exec(ctx, "INSERT INTO users (id ,email, password_hash , created_at) VALUES ($1 , $2 , $3 , $4)", baseUser.ID, baseUser.Email, baseUser.PasswordHash, baseUser.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
