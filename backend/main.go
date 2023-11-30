package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/vicheanath/go-bank/api"
	db "github.com/vicheanath/go-bank/db/sqlc"
	"github.com/vicheanath/go-bank/util"
)


func main() {

	config, err := util.LoadConfig(".")

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

