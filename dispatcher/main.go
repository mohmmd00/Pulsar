package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/mohmmd00/Pulsar/dispatcher/internal/handlers"
)

func main() {
	envLoadErr := godotenv.Load(".env")
	if envLoadErr != nil {
		log.Fatal("Error loading .env file")
	}

	databaseUrls := os.Getenv("DATABASE_URL")
	dbPool, err := pgxpool.New(context.Background(), databaseUrls)
	if err != nil {
		log.Fatal("Error crafting connection to db")
	}

	srv := &handlers.Server{DatabasePool: dbPool}

	defer dbPool.Close() // open until main returns

	router := gin.Default()
	router.GET("/ping", srv.Ping)

	router.Run(os.Getenv("API_PORT"))

}
