package db

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/raj23manj/techshool-go/util"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("./../..") // . mean current folder
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	// inform os that connnection was successfulr nad terminate the test case
	os.Exit(m.Run())
}

// setup and tear down
func CreateForEach(setUp func(), tearDown func(ctx context.Context)) func(func()) {
	ctx := context.Background()
	return func(testFunc func()) {
		setUp()
		testFunc()
		tearDown(ctx)
	}
}
