package handlers

import (
	"net/http"

	"github.com/arathi/wealth-manager/internal/config"
	"github.com/arathi/wealth-manager/internal/forms"
	"github.com/arathi/wealth-manager/internal/models"
	"github.com/arathi/wealth-manager/internal/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// Creates repository used by handler
func CreateHandlerRepo(a *config.AppConfig) {
	Repo = &Repository{
		App: a,
	}
}

// Home is handler for home page
func (m *Repository) Home(rw http.ResponseWriter, r *http.Request) {
	render.Template("home.page.tmpl", rw, r, &models.Data{})
}

// ShowSignUp is handler for showing sign-up page
func (m *Repository) ShowSignUp(rw http.ResponseWriter, r *http.Request) {
	render.Template("sign-up.page.tmpl", rw, r, &models.Data{})
}

func (m *Repository) ShowLogin(rw http.ResponseWriter, r *http.Request) {
	render.Template("login.page.tmpl", rw, r, &models.Data{
		Form: forms.Init(nil),
	})
}
