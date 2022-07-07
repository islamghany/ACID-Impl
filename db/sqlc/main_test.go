package db

import (
	"context"
	"database/sql"
	"github/islmaghany/bank/util"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
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
	testQueries = New(conn)
	os.Exit(m.Run())
}
