package main

import (
	"context"
	"myboardgames/minimalgo/internal/models"
	"myboardgames/minimalgo/internal/server"
	"os"
	"os/signal"
)

func main() {
	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer cancel()

	// app := app.NewHttpApp(ctx)
	// app.Start()

	server := server.NewServerApp(ctx, &models.ServerAppOptions{})
	server.Start(ctx)
}
