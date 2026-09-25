package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"

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

	natsConn, natsConErr := nats.Connect(os.Getenv("NATS_URL"))
	if natsConErr != nil {
		log.Fatal("Error connecting to NATS.")
	}
	defer natsConn.Close()

	dbPool, pgxCraftErr := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if pgxCraftErr != nil {
		log.Fatal("Error crafting connection to db.")
	}
	defer dbPool.Close()

	concurrencyStr := os.Getenv("WORKER_CONCURRENCY")
	concurrency, concurLoadErr := strconv.Atoi(concurrencyStr)
	if concurLoadErr != nil {
		log.Fatal("falied to load concurency number.")
	}

	worker := handler.Worker{DatabasePool: dbPool, NatsConn: natsConn}

	//signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	//waitgroup
	var wg sync.WaitGroup

	//channel
	eventsInChan := make(chan models.NotificationEvent, concurrency*10)

	sub, subErr := worker.NatsConn.Subscribe("notifications.*", func(msg *nats.Msg) {

		event := models.NotificationEvent{}
		unmarshErr := json.Unmarshal(msg.Data, &event)
		if unmarshErr != nil {
			fmt.Println("failed to un marshal recieved notification into binding model : ", unmarshErr)
			return
		}

		eventsInChan <- event //put every new message into channel queue for go routines to pick one
	})
	if subErr != nil {
		log.Fatal("Error subscribing to NATS subject")
	}

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for event := range eventsInChan {
				processErr := service.ProcessNotification(event, worker.DatabasePool, context.Background())
				if processErr != nil {
					fmt.Println(processErr)
				} else {
					fmt.Printf("%s notification %s status has been changed !\n", time.Now(), event.NotificationID)
				}
			}
		}()

	}

	<-sigChan // close up call
	fmt.Println("shutting down, no longer accepting new work...")
	sub.Unsubscribe()   // closing nats connection
	close(eventsInChan) // closing channel so no more events pile up in channel queue
	wg.Wait()           // doesnt close until every procees processes
	fmt.Println("all workers finished, exiting cleanly")

}
