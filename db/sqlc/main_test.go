package db

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5431/bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	err = conn.PingContext(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}
