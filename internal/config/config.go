package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/validator/v10"
)

// AppConfig holds the application config
type AppConfig struct {
	InfoLog        *log.Logger
	ErrorLog       *log.Logger
	SessionManager *scs.SessionManager
	TemplateCache  map[string]*template.Template
	ProductionEnv  bool
	Validate       *validator.Validate
}
