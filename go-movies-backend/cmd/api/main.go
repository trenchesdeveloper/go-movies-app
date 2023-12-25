package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/trenchesdeveloper/go-backend/internal/repository"
	"github.com/trenchesdeveloper/go-backend/internal/repository/dbrepo"
)

const port = 8080

type application struct {
	Domain string
	DSN    string
	DB     repository.DatabaseRepo
}

func main() {
	// set application config
	var app application

	// read from command line
	flag.StringVar(&app.Domain, "domain", "localhost", "Domain for the application")
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 dbname=movies user=postgres password=postgres sslmode=disable connect_timeout=5", "Postgres connection string")

	flag.Parse()

	// connect to database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	app.DB.Connection().Close()
	app.Domain = "example.com"

	router := app.routes()

	log.Println("Starting server on port", port)
	// start the server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), router)

	if err != nil {
		log.Fatal(err)
	}
}
