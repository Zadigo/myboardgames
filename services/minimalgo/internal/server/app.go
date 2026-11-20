package server

import (
	"context"
	"errors"
	"fmt"
	"myboardgames/minimalgo/internal/backend"
	"myboardgames/minimalgo/internal/games"
	httpapp "myboardgames/minimalgo/internal/http"
	"myboardgames/minimalgo/internal/models"
	"time"

	"github.com/redis/go-redis/v9"
)

type ServerApp struct {
	models.BaseServerApp
	_redis *redis.Client
}

func (s *ServerApp) GetRedis() *redis.Client {
	return s._redis
}

func (s *ServerApp) Start(ctx context.Context) {
	chError := make(chan error)

	serverCtx := context.WithValue(ctx, "state", true)

	fmt.Println("Initializing main server")

	httpApp := httpapp.NewHttpApp(ctx, s)
	go func(ctx context.Context) {
		httpApp.Start(ctx)
	}(ctx)

	gameApp := games.NewGameServerApp(ctx)
	go func(ctx context.Context) {
		gameApp.Start(ctx)
	}(ctx)

	go checkDjangoBackend(serverCtx, chError)

	select {
	case err := <-chError:
		fmt.Println(err)
	case <-serverCtx.Done():
		timeoutCtx, cancel := context.WithTimeout(ctx, 10000*time.Second)
		defer cancel()

		fmt.Printf("server stopped %v", timeoutCtx.Err())
	}
}

func NewServerApp(ctx context.Context, options *models.ServerAppOptions) models.ServerAppInterface {
	instance, err := backend.CreateRedisConnection()
	if err != nil {
		panic(errors.Join(FailedToLoadRedisError, err))
	}
	return &ServerApp{_redis: instance}
}
