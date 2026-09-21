package main

import (
	"context"
	"log"
	"os"

	"github.com/mohmmd00/Pulsar/dispatcher/internal/handlers"

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
	dbPool, err := pgxpool.New(context.Background(), databaseUrls)
	if err != nil {
		log.Fatal("Error crafting connection to db") // this loog using os.exit(1) so it will crash if used
	}

	srv := &handlers.Server{DatabasePool: dbPool}

	defer dbPool.Close() // open until main returns

	router := gin.Default()

	//endpoints
	router.GET("/ping", srv.Ping)

	router.POST("/register", srv.RegisterHandler)

	// router run
	router.Run(os.Getenv("API_PORT"))

}
