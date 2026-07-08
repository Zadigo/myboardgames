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

	mainCtx = context.WithValue(mainCtx, "baseDir", ".")

	server := server.NewServerApp(mainCtx)
	err := server.Start()

	if err != nil {
		panic(err)
	}
}
