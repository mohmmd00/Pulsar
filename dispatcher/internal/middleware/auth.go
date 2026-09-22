package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mohmmd00/Pulsar/dispatcher/internal/webToken"
)

func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// 1- read header
		header := c.GetHeader("Authorization")
		if header == "" {
			c.JSON(401, gin.H{"error": "header is empty"})
			c.Abort() 
			return
		}
		//2- strip the 'Bearer ' make jwt raw !!!

		rawToken := strings.TrimPrefix(header,"Bearer ") // the space is important !

		claims,validErr := webToken.ValidateToken(rawToken , secret) // passing raw jwt with secret from .env file
		if validErr != nil {
			c.JSON(401 , gin.H{"error" : "failed to validate token"})
			c.Abort()
			return
		}

		//3- put user id into context 
		c.Set("user_id" , claims.UserID)
		c.Next()// going intoo next midlleware 

	}
}
