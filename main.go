package main

import (
	"context"
	"database/sql"
	"log"
	"simple-bank/api"
	db "simple-bank/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5431/bank?sslmode=disable"
	addr     = "0.0.0.0:4000"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)
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

	err = server.Start(addr)
	if err != nil {
		log.Fatal(err)
	}

}
