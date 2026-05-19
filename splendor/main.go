package main

import (
	"net/http"

	"github.com/Zadigo/splendor/internal/backend"
	"github.com/Zadigo/splendor/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	serverRegistry := backend.NewServerRegistry()
	router := chi.NewRouter()

	router.Get("/ws/splendor/live", func(w http.ResponseWriter, r *http.Request) {
		handlers.LiveGameHandler(w, r, serverRegistry)
	})

	http.ListenAndServe(":9000", router)
}
