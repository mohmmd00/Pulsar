package webToken

import (
	"errors"
	"fmt"
	"time"
	"uuid"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserID uuid.UUID `json:"user_id"` // bind to user id
	jwt.RegisteredClaims
}

func GenerateToken(userId uuid.UUID, secret string) (string, error) {

	claims := CustomClaims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedString, err := token.SignedString([]byte(secret))
	return signedString, err
}
func ValidateToken(tokenString string, secret string) (*CustomClaims, error) {
	if tokenString == "" {
		return nil, errors.New("token is empty")
	}

	if secret == "" {
		return nil, errors.New("JWT secret is empty")
	}

	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		// Make sure the token was signed using HS256.
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf(
				"unexpected signing method: %s",
				token.Method.Alg(),
			)
		}

		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token is invalid")
	}

	if claims.UserID == uuid.Nil() {
		return nil, errors.New("token does not contain a valid user_id")
	}

	return claims, nil
}
