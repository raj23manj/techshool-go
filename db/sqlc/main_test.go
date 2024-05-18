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
	dbSource = "postgres://rajeshmanjunath:password@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(conn)

	// inform os that connnection was successfulr nad terminate the test case
	os.Exit(m.Run())
}

// setupa and tear down
func CreateForEach(setUp func(), tearDown func(ctx context.Context)) func(func()) {
	ctx := context.Background()
	return func(testFunc func()) {
		setUp()
		testFunc()
		tearDown(ctx)
	}
}
