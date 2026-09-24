package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/mohmmd00/Pulsar/worker/internal/handler"
	"github.com/mohmmd00/Pulsar/worker/internal/models"
	"github.com/mohmmd00/Pulsar/worker/internal/service"

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

	worker := handler.Worker{DatabasePool: dbPool, NatsConn: natsConn}

	_, err = worker.NatsConn.Subscribe("notifications.*", func(msg *nats.Msg) {

		event := models.NotificationEvent{}
		unmarshErr := json.Unmarshal(msg.Data, &event)
		if unmarshErr != nil {
			fmt.Println("failed to un marshal recieved notification into binding model : ", unmarshErr)
		}

		processErr := service.ProcessNotification(event, worker.DatabasePool, context.Background())
		if processErr != nil {
			fmt.Println(processErr)
		}

		fmt.Printf("notification %s status has been changed !\n", event.NotificationID)

	})
	if err != nil {
		log.Fatal("Error subscribing to NATS subject")
	}

	select {} // block forever, keeping main() alive so the subscription keeps listening

}
