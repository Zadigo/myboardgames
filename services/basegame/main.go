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

	mainCtx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	debug := os.Getenv("DEBUG") == "true"

	mainCtx = context.WithValue(mainCtx, "baseDir", ".")
	mainCtx = context.WithValue(mainCtx, "debug", debug)

	server := server.NewServerApp(mainCtx)
	err := server.Start()

	if err != nil {
		panic(err)
	}
}
