package main

import (
	"net/http"

	"github.com/arathi/wealth-manager/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Test Route"))
	})

	return mux
}
