package main

import (
	"context"
	"log"
	"os"

	"github.com/mohmmd00/Pulsar/dispatcher/internal/handlers"
	"github.com/mohmmd00/Pulsar/dispatcher/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	envLoadErr := godotenv.Load(".env")
	if envLoadErr != nil {
		log.Fatal("Error loading .env file")
	}

	databaseUrls := os.Getenv("DATABASE_URL")
	jwtSecret := os.Getenv("JWT_SECRET")

	dbPool, err := pgxpool.New(context.Background(), databaseUrls)
	if err != nil {
		log.Fatal("Error crafting connection to db") // this loog using os.exit(1) so it will crash if used
	}

	srv := &handlers.Server{DatabasePool: dbPool, JWTSecret: jwtSecret}

	defer dbPool.Close() // open until main returns

	router := gin.Default()

	//endpoints
	router.GET("/ping", srv.Ping)

	router.GET("/me", middleware.AuthMiddleware(srv.JWTSecret), srv.MeHandler)

	router.POST("/register", srv.RegisterHandler)

	router.POST("/login", srv.LoginHandler)

	// router run
	router.Run(os.Getenv("API_PORT"))

}
