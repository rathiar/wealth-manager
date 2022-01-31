package main

import (
	"crypto/rand"
	"io"
	"net/http"

	"github.com/arathi/wealth-manager/internal/config"
	"github.com/arathi/wealth-manager/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/csrf"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(CSRFHandler)
	// SessionManager.LoadAndSave automatically loads and saves session data for the current request,
	// 	and communicates the session token to and from the client in a cookie.
	mux.Use(app.SessionManager.LoadAndSave)

	// Static content handling
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/sign-up", handlers.Repo.ShowSignUp)
	mux.Get("/login", handlers.Repo.ShowLogin)
	mux.Post("/login", handlers.Repo.Login)

	return mux
}

func CSRFHandler(next http.Handler) http.Handler {
	return csrf.Protect(
		generateToken(),
		csrf.Secure(app.ProductionEnv),
		csrf.SameSite(csrf.SameSiteLaxMode),
		csrf.Path("/"),
		csrf.HttpOnly(true),
		csrf.FieldName("csrfField"),
	)(next)
}

// generateToken generates token of specified length
func generateToken() []byte {
	bytes := make([]byte, tokenLength)

	_, err := io.ReadFull(rand.Reader, bytes)
	if err != nil {
		panic(err)
	}
	return bytes
}
