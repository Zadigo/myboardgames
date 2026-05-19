package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/Zadigo/gosplendor/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	rootDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	app := server.NewApp(server.LoadConfig(rootDir))
	err = app.Start(ctx)
	if err != nil {
		panic(err)
	}
}
