package main

import (
	"context"
	"myboardgames/minimalgo/internal/backend"
	"myboardgames/minimalgo/internal/models"
	"myboardgames/minimalgo/internal/server"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer cancel()

	// Redis client
	if client, err := backend.CreateRedisConnection(); err != nil {
		panic(err)
	} else {
		// Main server
		server := server.NewServerApp(ctx, &models.ServerAppOptions{RedisClient: client})
		err = server.Start(ctx)
		if err != nil {
			panic(err)
		}
	}
}
