package handlers

import (
	"log"
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

// ShowLogin is handler for showing login page
func (m *Repository) ShowLogin(rw http.ResponseWriter, r *http.Request) {
	render.Template("login.page.tmpl", rw, r, &models.Data{
		Form: forms.New(nil),
	})
}

// Login is handler for logging in user
func (m *Repository) Login(rw http.ResponseWriter, r *http.Request) {
	if !renewToken(rw, r) {
		return
	}
	err := r.ParseForm()
	if err != nil {
		m.App.ErrorLog.Println(err)
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	form := forms.New(r.PostForm)
	form.ValidateField("required", "email", email, "email is required")
	form.ValidateField("email", "email", email, "email format not valid")
	form.ValidateField("required", "password", password, "password is required")
	log.Println(form.Errors)
	if !form.Valid() {
		m.App.ErrorLog.Println("There are validation errors. Rendering login page again")
		render.Template("login.page.tmpl", rw, r, &models.Data{
			Form: form,
		})
		return
	}

	// TODO: temporary logic for authentication till DB is integrated
	if email == "test@test.com" && password == "test" {
		m.App.SessionManager.Put(r.Context(), "user_id", 100)
		m.App.SessionManager.Put(r.Context(), "SuccessMsg", "Logged in successfully")
		http.Redirect(rw, r, "/", http.StatusSeeOther)
	} else {
		// Authentication Failed
		m.App.SessionManager.Put(r.Context(), "ErrorMsg", "Invalid credentials")
		http.Redirect(rw, r, "/login", http.StatusSeeOther)
		return
	}
}

// renewToken renews token to prevent session fixation
func renewToken(rw http.ResponseWriter, r *http.Request) bool {
	err := Repo.App.SessionManager.RenewToken(r.Context())
	if err != nil {
		http.Error(rw, "Error renewing token", http.StatusInternalServerError)
		return false
	}
	return true
}
