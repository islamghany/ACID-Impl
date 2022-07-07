package main

import (
	"context"
	"database/sql"
	"github/islmaghany/bank/api"
	db "github/islmaghany/bank/db/sqlc"
	"github/islmaghany/bank/util"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	err = conn.PingContext(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal(err)
	}

}
