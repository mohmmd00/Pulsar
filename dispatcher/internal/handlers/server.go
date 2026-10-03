package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nats-io/nats.go"
)

type Server struct {
	DatabasePool   *pgxpool.Pool
	JWTSecret      string
	NatsConn       *nats.Conn
	JwtAccTime     time.Duration
	JwtRefreshTime time.Duration
}

func (s *Server) Ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong!"}) // ping pong !
}
func (s *Server) MeHandler(c *gin.Context) {

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(500, gin.H{"error": "user_id not found in context"})
		return
	}
	c.JSON(200, gin.H{"user_id": userID})
}
