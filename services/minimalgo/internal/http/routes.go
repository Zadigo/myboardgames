package http

import "github.com/go-chi/chi/v5"

func (h *HttpApp) loadRoutes() {
	h.router = chi.NewRouter()
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
