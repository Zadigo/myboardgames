package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Zadigo/flipseven2/internal/backend"
	"github.com/Zadigo/flipseven2/internal/backend/broadcasting"
	"github.com/Zadigo/flipseven2/internal/handlers"
	"github.com/Zadigo/flipseven2/internal/models"
)

func main() {
	log.Print("🚀 Starting Flip 7 server...")

	serverBaseConfig := &backend.ServerBaseConfig{
		Backends: &backend.ServerBackendsConfig{
			Redis: &backend.ServerBackendConfig{
				Url: "redis://:@localhost:6379/0",
			},
		},
	}

	redisClient, err := backend.CreateRedisClient(serverBaseConfig.Backends.Redis)

	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := broadcasting.NewSubscription(redisClient, ctx)
	b := broadcasting.NewBroadcastingRegistry(s, ctx)

	baseRegistry := models.CreateBaseRegistry()
	log.Print("🟢 Backends initialized successfully")

	handler := http.HandlerFunc(handlers.Cors(func(w http.ResponseWriter, r *http.Request) {
		handlers.GameEngine(w, r, redisClient, b, baseRegistry)
	}))

	http.HandleFunc("/ws/flip-seven", handler)

	// handlerCandidate := http.HandlerFunc(handlers.Cors(func(w http.ResponseWriter, r *http.Request) {
	// 	handlers.GameCandidate(w, r, redisClient, b, baseRegistry)
	// }))

	// http.HandleFunc("/ws/flip-seven/candidate", handlerCandidate)

	err = http.ListenAndServe(":9000", nil)
	if err != nil {
		panic(err)
	}
}
