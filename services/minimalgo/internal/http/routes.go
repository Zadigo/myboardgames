package http

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (h *HttpApp) loadRoutes() {
	h.router = chi.NewRouter()

	// Applied to everything
	h.router.Use(middleware.RequestID)
	h.router.Use(middleware.ClientIPFromRemoteAddr)
	h.router.Use(middleware.Logger)
	h.router.Use(middleware.Recoverer)
	h.router.Use(Cors)

	base := &BaseHandlers{app: h}
	h.router.Route("/v1", func(r chi.Router) {
		// Long-lived: no Timeout / Throttle / AllowContentType
		r.Get("/", base.JoinWebsite)

		// Regular short HTTP routes
		r.Group(func(r chi.Router) {
			r.Use(middleware.AllowContentType("application/json"))
			r.Use(middleware.Throttle(1000))
			r.Use(middleware.Timeout(60 * time.Second))

			r.Route("/{tableId}", func(r chi.Router) {
				r.Get("/join", base.JoinTable)
			})
		})
	})
}
