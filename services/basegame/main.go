package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/Zadigo/basegame/internal/server"
	"github.com/Zadigo/basegame/internal/utils"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	mainCtx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	debug := os.Getenv("DEBUG") == "true"

	baseDir, err := utils.GetAbsolutePath(".")
	if err != nil {
		panic(err)
	}

	mainCtx = context.WithValue(mainCtx, "baseDir", baseDir)
	mainCtx = context.WithValue(mainCtx, "debug", debug)

	server := server.NewServerApp(mainCtx)
	err = server.Start()

	if err != nil {
		panic(err)
	}
}
