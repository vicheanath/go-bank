package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/vicheanath/go-bank/api"
	db "github.com/vicheanath/go-bank/db/sqlc"
)

const (
	DBDriver      = "postgres"
	DBSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	ServerAddress = "0.0.0.0:80"
)

func main() {
	conn, err := sql.Open(DBDriver, DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(ServerAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

