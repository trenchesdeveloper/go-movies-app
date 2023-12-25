package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
}

func main() {
	// set application config
	var app application

	// read from command line

	// connect to database

	app.Domain = "example.com"

	router := app.routes()

	log.Println("Starting server on port", port)
	// start the server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router)

	if err != nil {
		log.Fatal(err)
	}
}