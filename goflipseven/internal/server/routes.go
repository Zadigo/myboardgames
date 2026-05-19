package server

import (
	"time"

	"github.com/Zadigo/goflipseven/internal/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func (a *App) loadUrls() {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(handlers.Cors)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	router.Route("/ws/splendor", a.loadGameUrls)

	a.router = router
}

func (a *App) loadGameUrls(r chi.Router) {
	handlers := handlers.LiveGame{}
	r.Get("/live", handlers.Connect)
}
