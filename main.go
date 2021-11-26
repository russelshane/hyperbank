package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/russelshane/hyperbank/api"
	db "github.com/russelshane/hyperbank/db/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/hyperbank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Error while connecting to database:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("failed to start server:", err)
	}
}