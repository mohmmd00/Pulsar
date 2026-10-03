package service

import (
	"context"
	"fmt"
	"time"
	"uuid"

	"github.com/mohmmd00/Pulsar/dispatcher/internal/models"
	"github.com/mohmmd00/Pulsar/dispatcher/internal/repository"
	"github.com/mohmmd00/Pulsar/dispatcher/internal/webToken"

	"github.com/jackc/pgx/v5/pgxpool"
)

func IssueRefreshToken(userID uuid.UUID, refreshTTL time.Duration, pool *pgxpool.Pool, ctx context.Context) (rawToken string, token models.RefreshToken, err error) {

	errMsg := "Couldnt create new refresh token"
	rndTkn, tknErr := webToken.GenerateRandomToken()
	if tknErr != nil {
		return "", models.RefreshToken{}, fmt.Errorf(errMsg+"%w", tknErr)
	}

	hashedTkn := webToken.HashToken(rndTkn)
	newTkn := models.NewRefreshToken(userID, hashedTkn, time.Now().Add(refreshTTL) , false )

	//repositoryLayer 
	repoErr := repository.CreateRefreshToken(newTkn , pool , ctx)
	if repoErr != nil {
		return "", models.RefreshToken{}, fmt.Errorf(errMsg+"%w", repoErr)
	}

	return rndTkn , newTkn , nil

}
