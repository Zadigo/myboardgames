package main

import (
	"context"
	"myboardgames/minimalgame/internal"
	"myboardgames/minimalgame/internal/models"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	app := internal.NewApp(ctx, &models.ServerOptions{Port: 9001})
	app.Start(ctx)
}
