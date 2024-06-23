package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/raj23manj/techshool-go/api"
	db "github.com/raj23manj/techshool-go/db/sqlc"
	"github.com/raj23manj/techshool-go/util"
)

func main() {
	config, err := util.LoadConfig(".") // . mean current folder
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
