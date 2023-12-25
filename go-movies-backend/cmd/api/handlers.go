package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/trenchesdeveloper/go-backend/internal/models"
)

func (a *application) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from %s\n", a.Domain)
}

func (a *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movie
	rd, _ := time.Parse("2006-01-02", "1981-06-12")

	// Create a new movie instance
	highlander := models.Movie{
		ID:          1,
		Title:       "Highlander",
		ReleaseDate: rd,
		Runtime:     116,
		MPAARating:  "R",
		Description: "A random description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, highlander)

	rota := models.Movie{
		ID:          2,
		Title:       "Rota de Fuga",
		ReleaseDate: rd,
		Runtime:     116,
		MPAARating:  "R",
		Description: "A random description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, rota)

	// Marshal the payload to JSON

	out, err := json.MarshalIndent(movies, "", "\t")

	if err != nil {
		fmt.Println(err)
	}

	// Write the JSON payload to the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}
