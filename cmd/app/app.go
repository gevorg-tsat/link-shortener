package main

import (
	"github.com/gevorg-tsat/link-shortener/internal/storage"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func main() {
	_, err := storage.NewDB(host, user, password, dbname, port)
	if err != nil {
		log.Fatal(err)
	}
}
