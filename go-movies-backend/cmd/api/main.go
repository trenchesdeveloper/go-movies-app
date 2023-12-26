package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/trenchesdeveloper/go-backend/internal/repository"
	"github.com/trenchesdeveloper/go-backend/internal/repository/dbrepo"
)

const port = 8080

type application struct {
	Domain       string
	DSN          string
	DB           repository.DatabaseRepo
	auth         Auth
	JWTSecret    string
	JwtIssuer    string
	JWTAudience  string
	CookieDomain string
}

func main() {
	// set application config
	var app application

	// read from command line
	flag.StringVar(&app.Domain, "domain", "example.com", "Domain for the application")
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 dbname=movies user=postgres password=postgres sslmode=disable connect_timeout=5", "Postgres connection string")
	flag.StringVar(&app.JWTSecret, "jwt-secret", "secret", "Secret key for JWT")
	flag.StringVar(&app.JwtIssuer, "jwt-issuer", "example.com", "Issuer for JWT")
	flag.StringVar(&app.JWTAudience, "jwt-audience", "example.com", "Audience for JWT")
	flag.StringVar(&app.CookieDomain, "cookie-domain", "localhost", "Domain for the cookie")
	flag.Parse()

	// connect to database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()

	app.auth = Auth{
		Issuer:        app.JwtIssuer,
		Audience:      app.JWTAudience,
		Secret:        app.JWTSecret,
		TokenExpiry:   time.Minute * 15,
		RefreshExpiry: time.Hour * 24,
		CookieDomain:  app.CookieDomain,
		CookiePath:    "/",
		CookieName:    "__Host-refresh-token",

	}

	router := app.routes()

	log.Println("Starting server on port", port)
	// start the server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), router)

	if err != nil {
		log.Fatal(err)
	}
}
