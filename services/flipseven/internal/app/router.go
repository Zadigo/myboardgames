package app

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *App) loadRouter() {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(Cors)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(JsonHeartbeat("/health"))
	router.Use(middleware.Timeout(60 * time.Second))
	router.Use(Authorization)

	router.Route("/v1/", app.loadServerRoutes)

	app.router = router
}

func (app *App) loadServerRoutes(r chi.Router) {

}
