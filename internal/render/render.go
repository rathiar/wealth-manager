package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/arathi/wealth-manager/internal/config"
	"github.com/arathi/wealth-manager/internal/models"
	"github.com/gorilla/csrf"
)

var templatePath = "./templates"
var tFuncs = template.FuncMap{}
var app *config.AppConfig

// Initializes Renderer
func InitRenderer(a *config.AppConfig) {
	app = a
}

func addDefaultData(td *models.Data, r *http.Request) *models.Data {
	td.CSRFToken = csrf.Token(r)
	return td
}

// Template renders template
func Template(tmpl string, rw http.ResponseWriter, r *http.Request, data *models.Data) error {
	var tc = map[string]*template.Template{}
	if app.ProductionEnv {
		tc = app.TemplateCache
	} else {
		// In development mode, build cache every time so that template changes are effective without server restart
		tc, _ = CacheTemplates()
	}

	t, ok := tc[tmpl]

	if !ok {
		app.ErrorLog.Printf("can't find %s template in cache", tmpl)
		return fmt.Errorf("can't find %s template in cache", tmpl)
	}

	buf := new(bytes.Buffer)

	data = addDefaultData(data, r)

	err := t.Execute(buf, data)
	if err != nil {
		log.Fatal(err)
	}

	app.InfoLog.Printf("Buffer size: %v", buf.Len())

	_, err = buf.WriteTo(rw)
	if err != nil {
		app.ErrorLog.Println("Error writing template to browser", err)
		return err
	}

	return nil
}

// CacheTemplates caches templates
func CacheTemplates() (map[string]*template.Template, error) {
	tCache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", templatePath))

	if err != nil {
		return tCache, err
	}

	for _, page := range pages {
		tName := filepath.Base(page)
		ts, err := template.New(tName).Funcs(tFuncs).ParseFiles(page)

		if err != nil {
			return tCache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", templatePath))
		if err != nil {
			return tCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", templatePath))
			if err != nil {
				return tCache, err
			}
		}

		tCache[tName] = ts
	}
	app.InfoLog.Printf("Built template cache with %d templates", len(tCache))
	return tCache, nil

}
