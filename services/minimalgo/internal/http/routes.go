package http

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (h *HttpApp) loadRoutes() {
	h.router = chi.NewRouter()

	h.router.Use(middleware.RequestID)
	h.router.Use(middleware.ClientIPFromRemoteAddr)
	h.router.Use(Cors)
	// h.router.Use(Authorization)
	h.router.Use(middleware.AllowContentType("application/json"))
	h.router.Use(middleware.Throttle(1000))
	h.router.Use(middleware.Logger)
	h.router.Use(middleware.Recoverer)
	// h.router.Use(JsonHeartbeat("/health"))
	h.router.Use(middleware.Timeout(60 * time.Second))

	h.router.Route("/v1/", h.loadBaseRoutes)
}

func (h *HttpApp) loadBaseRoutes(r chi.Router) {
	base := &BaseHandlers{app: h}
	r.Get("/", base.JoinWebsite)
	r.Route("/{tableId}", func(r chi.Router) {
		r.Get("/join", base.JoinTable)
		r.Get("/observe", base.ObserveTable)
	})
}
