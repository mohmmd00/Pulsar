package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/mohmmd00/Pulsar/dispatcher/internal/models"
	"github.com/mohmmd00/Pulsar/dispatcher/internal/repository"
	"github.com/mohmmd00/Pulsar/dispatcher/internal/webToken"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(email string, password string, pool *pgxpool.Pool, ctx context.Context) (models.User, error) {

	errMsg := "couldnt register new user : "

	if email == "" || password == "" {
		return models.User{}, errors.New(errMsg + "invalid inputs")
	}

	//hash

	hashedPassword, cryptErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // hash pass word on default cost
	if cryptErr != nil {
		return models.User{}, fmt.Errorf(errMsg+"%w", cryptErr) //send back actual error not plain text
	}

	newUser := models.NewUser(email, string(hashedPassword))

	//repository layer

	repoErr := repository.CreateUser(newUser, pool, ctx)
	if repoErr != nil {
		return models.User{}, fmt.Errorf(errMsg+"%w", repoErr) //send back actual error not plain text
	}

	return newUser, nil

}

func LoginUser(email string, password string, pool *pgxpool.Pool, ctx context.Context, secret string) (jwt string, err error) {
	errMsg := "couldnt login user : "

	if email == "" || password == "" {
		return "", errors.New(errMsg + "invalid inputs")
	}

	//repository layer 
	fetchedUser, repoErr := repository.GetUserByEmail(email, pool, ctx)

	if repoErr != nil {
		return "", fmt.Errorf(errMsg+"%w", repoErr)
	}

	//compare hash
	cryptErr := bcrypt.CompareHashAndPassword([]byte(fetchedUser.PasswordHash), []byte(password))

	if cryptErr != nil {
		return "", fmt.Errorf(errMsg+"%w", cryptErr) //send back actual error not plain text
	}

	tknAssigned, tknErr := webToken.GenerateToken(fetchedUser.ID, secret)
	if tknErr != nil {
		return "", fmt.Errorf(errMsg+"%w", tknErr)
	}

	return tknAssigned, nil
}
