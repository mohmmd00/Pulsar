# Pulsar

An authenticated notification dispatch system built as a Go microservices practice project.

- **dispatcher/** — HTTP API: auth, accepts notification jobs, publishes to NATS
- **worker/** — subscribes to NATS, processes jobs concurrently, updates Postgres

## Local dev infra
Postgres + NATS run via Docker Compose:
    docker compose up -d
