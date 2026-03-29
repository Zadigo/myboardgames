package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/Zadigo/flipseven/internal"
	"github.com/Zadigo/flipseven/internal/backend"
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

	http.HandleFunc("/ws/flip-seven", internal.Cors(internal.LiveGameHandler))
	http.HandleFunc("/v1/flip-seven/create", internal.Cors(internal.CreateTableHandler))

	redisClient := beforeStart()
	defer redisClient.Close()

	err := http.ListenAndServe(":9000", nil)

	if errors.Is(err, http.ErrServerClosed) {
		log.Println("❌ Server closed")
	} else {
		log.Fatalln("❌ Could not start server")
	}
}
