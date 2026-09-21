package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	DatabasePool *pgxpool.Pool
	JWTSecret    string
}

func (s *Server) Ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong!"}) // ping pong !
}
