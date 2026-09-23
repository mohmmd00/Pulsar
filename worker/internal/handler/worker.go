package handler

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nats-io/nats.go"
)

type Worker struct { //dependency container
	DatabasePool *pgxpool.Pool
	NatsConn     *nats.Conn
}
