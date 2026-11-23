package main

import (
	"context"
	"crypto/sha256"
	"fmt"
	"myboardgames/minimalgo/internal/backend"
	"myboardgames/minimalgo/internal/models"
	"myboardgames/minimalgo/internal/server"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func preStartSetup(ctx context.Context) (*redis.Client, *models.Router, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, nil, err
	}

	var router *models.Router

	if client, err := backend.CreateRedisConnection(ctx); err != nil {
		return nil, nil, err
	} else {
		if router, err = models.NewRouter(ctx, client); err != nil {
			return nil, nil, err
		} else {
			if err := router.Save(ctx, router); err != nil {
				return nil, nil, err
			}
		}

		// Hash each routes using the name in
		// in english for uniqueness and for
		// quick lookup
		for _, r := range router.Routes {
			r.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(r.NameEn)))
		}

		return client, router, nil
	}
}

func main() {
	ctx := context.Background()
	rootDir, _ := filepath.Abs(".")
	ctx = context.WithValue(ctx, "rootDir", rootDir)
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer cancel()

	if client, router, err := preStartSetup(ctx); err != nil {
		panic(err)
	} else {
		options := &models.ServerAppOptions{
			RedisClient: client,
			Router:      router,
		}

		// Main server
		server, err := server.NewServerApp(ctx, options)

		if err != nil {
			panic(err)
		}

		err = server.Start(ctx)
		if err != nil {
			panic(err)
		}
	}
}
