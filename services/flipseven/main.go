package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/Zadigo/flipseven/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	app := app.NewApp(ctx, ".")
	err := app.Start()

	if err != nil {
		panic(err)
	}
}
