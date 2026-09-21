package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mohmmd00/Pulsar/dispatcher/internal/models"
)

func CreateUser(baseUser models.User, pool *pgxpool.Pool, ctx context.Context) error {

	_, err := pool.Exec(ctx, "INSERT INTO users (id , email, password_hash , created_at) VALUES ($1 , $2 , $3 , $4)", baseUser.ID, baseUser.Email, baseUser.PasswordHash, baseUser.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}
func GetUserByEmail(email string, pool *pgxpool.Pool, ctx context.Context) (fetchedUser models.User, err error) {
	row := pool.QueryRow(ctx, "SELECT id, email, password_hash, created_at FROM users WHERE email = $1", email) // single row -> queryrow , many rows -> query , better to use every column insted of *
	scanErr := row.Scan(&fetchedUser.ID, &fetchedUser.Email, &fetchedUser.PasswordHash, &fetchedUser.CreatedAt)

	if scanErr != nil {
		return models.User{}, scanErr
	}

	return fetchedUser, nil
}
