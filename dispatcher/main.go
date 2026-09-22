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
	"github.com/nats-io/nats.go"
)

func main() {

	envLoadErr := godotenv.Load(".env")
	if envLoadErr != nil {
		log.Fatal("Error loading .env file")
	}

	natsConn, err := nats.Connect(os.Getenv("NATS_URL"))
	if err != nil {
		log.Fatal("Error connecting to NATS.")
	}
	defer natsConn.Close()

	dbPool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Error crafting connection to db.") // this loog using os.exit(1) so it will crash if used
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("Error recieving Jwt secret.")
	}

	srv := &handlers.Server{DatabasePool: dbPool, JWTSecret: jwtSecret, NatsConn: natsConn}

	defer dbPool.Close() // open until main returns

	router := gin.Default()

	//endpoints
	router.GET("/ping", srv.Ping)

	router.GET("/me", middleware.AuthMiddleware(srv.JWTSecret), srv.MeHandler)

	router.POST("/register", srv.RegisterHandler)

	router.POST("/login", srv.LoginHandler)

	router.POST("/notifications", middleware.AuthMiddleware(srv.JWTSecret), srv.CreateNotificationHandler)

	// router run
	router.Run(os.Getenv("API_PORT"))

}
