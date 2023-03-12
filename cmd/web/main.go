package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/wulaiyin789/go-web-bookings/pkg/config"
	"github.com/wulaiyin789/go-web-bookings/pkg/handlers"
	"github.com/wulaiyin789/go-web-bookings/pkg/render"
)

// go run . (Windows)
// go run *.go (Mac)
const PORT = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {
	// Change if production
	app.InProd = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProd

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", PORT))

	srv := &http.Server{
		Addr: PORT,
		// pat modules for routing
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
