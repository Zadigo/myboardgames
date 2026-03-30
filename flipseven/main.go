package main

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/Zadigo/flipseven/internal/backend"
	"github.com/Zadigo/flipseven/internal/handlers"
	"github.com/redis/go-redis/v9"
)

func beforeStart() *redis.Client {
	baseConfig := backend.ServerConfig{
		Config: backend.ServerBaseConfig{
			Backends: &backend.ServerBackendsConfig{
				Redis: &backend.ServerBackendConfig{
					Url: "redis://:@localhost:6379/0",
				},
			},
		},
	}

	client, err := backend.CreateRedisClient(baseConfig.Config.Backends.Redis)

	if err != nil {
		log.Panicln("❌ Failed to create Redis client", err.Error())
	}

	return client
}

func main() {
	log.Println("🚀 Starting Flip 7 Webserver...")
	log.Println("✅ Server started on 127.0.0.1:9000")

	// Pass a context to the handler with a cancel method
	// in order to close the current task
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	redisClient := beforeStart()
	defer redisClient.Close()

	// Initialise the pub/sub layer and broadcaster registry
	subscription := backend.NewSubscription(redisClient, ctx)
	broadcasterRegistry := backend.NewRegistry(subscription, ctx)

	registry := handlers.NewBaseRegistry()

	http.HandleFunc("/ws/flip-seven", handlers.Cors(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.LiveGameHandler(w, r, ctx, redisClient, broadcasterRegistry, registry)
	})))

	http.HandleFunc("/v1/flip-seven/create", handlers.Cors(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateTableHandler(w, r, redisClient)
	})))

	err := http.ListenAndServe(":9000", nil)

	if errors.Is(err, http.ErrServerClosed) {
		log.Println("❌ Server closed")
	} else {
		log.Fatalln("❌ Could not start server")
	}
}
