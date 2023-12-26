package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	mux.Get("/", app.Hello)

	mux.Post("/authenticate", app.authenticate)

	mux.Get("/refresh", app.refreshToken)

	mux.Get("/movies", app.AllMovies)
	return mux

}
