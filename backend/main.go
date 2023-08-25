package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/eduns/oncar-challenge/backend/src/infra"

	_ "github.com/lib/pq"
)

func main() {
	DB_URL := os.Getenv("DB_URL")
	connection, err := sql.Open("postgres", DB_URL)

	if err != nil {
		log.Fatal("[DATABASE CONNECTION ERROR] |>", err.Error())
	}

	defer connection.Close()

	router := infra.SetupServer(connection)
	router.Run(":8080")
}
