package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/Zadigo/basegame/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	server := server.NewServerApp(ctx)
	err := server.Start(".")

	if err != nil {
		panic(err)
	}
}
