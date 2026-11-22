package main

import (
	"context"
	"myboardgames/minimalgo/internal/backend"
	"myboardgames/minimalgo/internal/models"
	"myboardgames/minimalgo/internal/server"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/errgroup"
)

func preStartSetup(ctx context.Context) (*redis.Client, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	if client, err := backend.CreateRedisConnection(ctx); err != nil {
		return nil, err
	} else {
		g := errgroup.Group{}
		g.Go(func() error {
			if router, err := models.NewRouter(ctx, client); err != nil {
				return err
			} else {
				if err := router.Save(ctx, router); err != nil {
					return err
				}
			}

			return nil
		})
		return client, nil
	}
}

func main() {
	ctx := context.Background()
	rootDir, _ := filepath.Abs(".")
	ctx = context.WithValue(ctx, "rootDir", rootDir)
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer cancel()

	// Redis client
	if client, err := preStartSetup(ctx); err != nil {
		panic(err)
	} else {
		// Main server
		server, err := server.NewServerApp(ctx, &models.ServerAppOptions{
			RedisClient: client,
		})

		if err != nil {
			panic(err)
		}

		err = server.Start(ctx)
		if err != nil {
			panic(err)
		}
	}
}
