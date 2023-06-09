package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/wulaiyin789/go-web-bookings/pkg/config"
	"github.com/wulaiyin789/go-web-bookings/pkg/handlers"
)

// Pat module
// func routes(app *config.AppConfig) http.Handler {
// 	mux := pat.New()

// 	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
// 	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

// 	return mux
// }

// Chi module
func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	// mux.Use(WriteToConsole)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
