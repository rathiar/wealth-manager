package main

import (
	"net/http"

	"github.com/arathi/wealth-manager/internal/config"
	"github.com/arathi/wealth-manager/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	// Static content handling
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/sign-up", handlers.Repo.ShowSignUp)
	mux.Get("/login", handlers.Repo.ShowLogin)

	return mux
}
