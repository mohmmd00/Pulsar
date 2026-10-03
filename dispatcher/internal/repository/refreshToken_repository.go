package repository

import (
	"context"
	"uuid"

	"github.com/mohmmd00/Pulsar/dispatcher/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateRefreshToken(token models.RefreshToken, pool *pgxpool.Pool, ctx context.Context) error {

	_, err := pool.Exec(ctx, "INSERT INTO refresh_tokens (id,user_id,token_hash,expires_at,revoked,created_at) VALUES ($1,$2,$3,$4,$5,$6)",
		token.ID, token.UserID, token.TokenHash, token.ExpiresAt, token.Revoked, token.CreatedAt)
	return err
}
func GetRefreshTokenByHash(tokenHash string, pool *pgxpool.Pool, ctx context.Context) (token models.RefreshToken, err error) {
	effectedRow := pool.QueryRow(ctx, "SELECT id,user_id,token_hash,expires_at,revoked,created_at FROM refresh_tokens WHERE token_hash = $1", tokenHash)
	scanErr := effectedRow.Scan(&token.ID, &token.UserID, &token.TokenHash, &token.ExpiresAt, &token.Revoked, &token.CreatedAt)
	if scanErr != nil {
		return models.RefreshToken{}, scanErr
	}
	return token, nil
}
func RevokeRefreshToken(tokenID uuid.UUID, pool *pgxpool.Pool, ctx context.Context) error {
	_, err := pool.Exec(ctx, "UPDATE refresh_tokens SET revoked = true WHERE id = $1", tokenID)
	return err
}
