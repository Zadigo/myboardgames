package main

import (
	"context"
	"myboardgames/minimalgame/internal"
	"myboardgames/minimalgame/internal/backend"
	"myboardgames/minimalgame/internal/models"
	"myboardgames/minimalgame/internal/utils"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/errgroup"
)

func preStartSetup(ctx context.Context) (*redis.Client, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	// Redis
	client, err := backend.CreateRedisConnection(ctx)
	if err != nil {
		return nil, err
	}

	g, _ := errgroup.WithContext(ctx)
	
	g.Go(func() error {
		rules := &utils.Rule{}
		if err := utils.LoadRules(".", rules); err != nil {
			return err
		}
		
		rules.SetRedis(client)
		if err := rules.Save(ctx, rules); err != nil {
			return err
		}
		
		return nil
	})

	return client, g.Wait()
}

func main() {
	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	client, err := preStartSetup(ctx)
	if err != nil {
		panic(err)
	}

	app := internal.NewApp(ctx, client, &models.ServerOptions{
		Host: "localhost",
		Port: 9001,
	})
	app.Start(ctx)
}
