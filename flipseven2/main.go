package main

import (
	"context"

	"github.com/Zadigo/flipseven2/internal/backend"
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

	ctx := context.Background()
	backend.CreateRedisSubcription(redisClient, ctx)
}
