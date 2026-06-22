package app

import (
	"time"

	"github.com/Zadigo/flipseven/internal/handlers"
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
	genericHandler := handlers.GenericHandler{}
	genericHandler.SetApp(app)

	r.Post("/create", genericHandler.CreateGameHandler)

	r.Route("/{tableId}", func(r chi.Router) {
		r.Use(TableIdMiddleware)
		
		r.Get("/join", genericHandler.JoinGameHandler)
		r.Get("/observe", genericHandler.ObserveGameHandler)
	})
}
