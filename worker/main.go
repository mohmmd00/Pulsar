package main

import (
	"context"
	"fmt"
	"log"
	"os"

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
		log.Fatal("Error crafting connection to db.")
	}
	defer dbPool.Close()

	_, err = natsConn.Subscribe("notifications.*", func(msg *nats.Msg) {
		fmt.Println("Received message on subject:", msg.Subject)
		fmt.Println("Payload:", string(msg.Data))
	})
	if err != nil {
		log.Fatal("Error subscribing to NATS subject")
	}

	select {} // block forever, keeping main() alive so the subscription keeps listening

}
