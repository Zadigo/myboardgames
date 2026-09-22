package main

import (
	"net/http"

	"github.com/Zadigo/oriflamme/internal/backend"
	"github.com/Zadigo/oriflamme/internal/backend/redisclient"
	"github.com/Zadigo/oriflamme/internal/handlers"
)

func main() {
	// Base dependencies for the server
	redisClient := redisclient.NewRedisClient("redis://@localhost:6379/0")
	serverRegistry := backend.NewServerRegistry(redisClient)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.LiveGameHandler(w, r, serverRegistry)
	})

	http.Handle("/oriflamme/live", handler)
	http.ListenAndServe(":9000", nil)
}
