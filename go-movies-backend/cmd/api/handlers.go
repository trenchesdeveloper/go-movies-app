package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (a *application) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from %s\n", a.Domain)
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.DB.AllMovies()

	if err != nil {
		fmt.Println(err)
		return
	}

	// Marshal the payload to JSON
	out, _ := json.MarshalIndent(movies, "", "\t")

	// Write the JSON payload to the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}
