package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/arathi/wealth-manager/internal/config"
)

var app config.AppConfig

const portNumber = ":8080"

func main() {
	log.Println("Starting wealth manager application...")
	run()

	server := http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	app.InfoLog.Printf("Starting application on port %s", portNumber)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() {
	// Read Application Level Flags
	productionEnv := flag.Bool("production", true, "Application is in production")

	flag.Parse()

	app.ProductionEnv = *productionEnv

	// Initialize Loggers
	app.InfoLog = log.New(os.Stdout, "INFO\t", log.Default().Flags())
	app.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Default().Flags())

	// Session Configuration
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session

}
