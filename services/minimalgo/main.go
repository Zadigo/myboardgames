package main

import (
	"context"
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

	server := server.NewServerApp(ctx, &models.ServerAppOptions{})
	server.Start(ctx)
}
