package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Zadigo/flipseven2/internal/backend"
	"github.com/Zadigo/flipseven2/internal/backend/broadcasting"
	"github.com/Zadigo/flipseven2/internal/handlers"
	"github.com/Zadigo/flipseven2/internal/models"
)

func main() {
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

	fmt.Print(ctx, b, baseRegistry)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.GameEngine(w, r, redisClient, b, baseRegistry)
	})

	http.HandleFunc("/ws/live/game", handler)
	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
