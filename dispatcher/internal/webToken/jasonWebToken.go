package webToken

import (
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
func ValidateTokken() {}
