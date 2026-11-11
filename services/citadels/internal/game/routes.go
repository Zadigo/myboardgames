package game

import (
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func (a *Application) loadRoutes() {
	a.router = chi.NewRouter()

	a.router.Use(middleware.RequestID)
	a.router.Use(middleware.RealIP)
	// a.router.Use(Cors)
	// a.router.Use(Authorization)
	a.router.Use(middleware.AllowContentType("application/json"))
	a.router.Use(middleware.Throttle(1000))
	a.router.Use(middleware.Logger)
	a.router.Use(middleware.Recoverer)
	// a.router.Use(JsonHeartbeat("/health"))
	a.router.Use(middleware.Timeout(60 * time.Second))

	baseHandler := &GameHandler{}
	a.router.Get("/new", baseHandler.Get)
}
