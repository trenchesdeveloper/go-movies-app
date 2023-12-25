package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
	DSN    string
	DB     *sql.DB
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

	app.DB = conn

	app.Domain = "example.com"

	router := app.routes()

	log.Println("Starting server on port", port)
	// start the server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), router)

	if err != nil {
		log.Fatal(err)
	}
}
