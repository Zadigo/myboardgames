package main

import (
	"context"
	"fmt"

	"github.com/Zadigo/flipseven2/internal/backend"
	"github.com/Zadigo/flipseven2/internal/backend/broadcasting"
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

	defaultContext := context.Background()
	s := broadcasting.NewSubscription(redisClient, defaultContext)
	b := broadcasting.NewBroadcastingRegistry(s, defaultContext)

	fmt.Print(ctx, b)

}
