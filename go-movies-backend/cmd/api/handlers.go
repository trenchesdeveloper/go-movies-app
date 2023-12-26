package main

import (
	"fmt"
	"log"
	"net/http"
)

func (a *application) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from %s\n", a.Domain)
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.DB.AllMovies()

	if err != nil {
		app.errorJson(w, err)
		return
	}

	// Marshal the payload to JSON
	_ = app.writeJson(w, http.StatusOK, movies)

}

func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {
	// read a json payload

	// validate user against database

	// check password

	// create a jwt user
	u := jwtUser{
		ID:       1,
		FirstName: "John",
		LastName: "Doe",
	}

	//generate token
	tokens, err := app.auth.GenerateTokenPair(&u)

	if err != nil {
		app.errorJson(w, err)
		return
	}

	log.Println(tokens.Token)

	refreshCookie := app.auth.GetRefreshCookie(
		tokens.RefreshToken,
	)

	http.SetCookie(w, refreshCookie)

	// send token to user
	_ = app.writeJson(w, http.StatusOK, tokens)

}
