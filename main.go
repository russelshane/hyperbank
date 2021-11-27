package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/russelshane/hyperbank/api"
	db "github.com/russelshane/hyperbank/db/sqlc"
	"github.com/russelshane/hyperbank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Error while connecting to database:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("failed to start server", err)
	}


	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("failed to start server:", err)
	}
}