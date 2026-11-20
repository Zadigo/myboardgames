package server

import (
	"context"
	"errors"
	"log/slog"
	"myboardgames/minimalgo/internal/backend"
	"myboardgames/minimalgo/internal/games"
	"myboardgames/minimalgo/internal/http"
	"myboardgames/minimalgo/internal/models"
	"sync"
	"sync/atomic"

	"github.com/redis/go-redis/v9"
)

type ServerApp struct {
	models.BaseServerApp
	_redis *redis.Client
	// Channel that allows the HTTP app to communicate
	// with gameApp and inversely
	ChInternal chan models.InternalMessage
	isActive   atomic.Bool
}

func (s *ServerApp) GetState() bool {
	return s.isActive.Load()
}

func (s *ServerApp) WriteInternalMessage(ctx context.Context, msg models.InternalMessage) error {
	select {
	case s.ChInternal <- msg:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (s *ServerApp) GetInternalChannel() <-chan models.InternalMessage {
	return s.ChInternal
}

func (s *ServerApp) GetRedis() *redis.Client {
	return s._redis
}

func (s *ServerApp) Start(ctx context.Context) {
	// chError := make(chan error)

	serverCtx := context.WithValue(ctx, "state", true)

	slog.Info("Initializing main server")

	// httpApp := httpapp.NewHttpApp(ctx, s)
	// go func(ctx context.Context) {
	// 	httpApp.Start(ctx)
	// }(ctx)

	// gameApp := games.NewGameServerApp(ctx, s)
	// go func(ctx context.Context) {
	// 	gameApp.Start(ctx)
	// }(ctx)

	// go checkDjangoBackend(serverCtx, chError)

	// select {
	// case err := <-chError:
	// 	fmt.Println(err)
	// case <-serverCtx.Done():
	// 	timeoutCtx, cancel := context.WithTimeout(ctx, 10000*time.Second)
	// 	defer cancel()

	// 	fmt.Printf("server stopped %v", timeoutCtx.Err())
	// }

	var wg sync.WaitGroup

	wg.Add(2)
	wg.Go(func() {
		defer wg.Done()
		httpApp := http.NewHttpApp(ctx, s)
		httpApp.Start(ctx)
	})

	wg.Go(func() {
		defer wg.Done()
		gameApp := games.NewGameServerApp(ctx, s)
		gameApp.Start(ctx)
	})

	// wg.Go(func() {
	// 	defer wg.Done()
	// 	checkDjangoBackend(serverCtx, chError)
	// })

	s.isActive.Store(true)

	<-serverCtx.Done()

	wg.Wait()
}

func NewServerApp(ctx context.Context, options *models.ServerAppOptions) models.ServerAppInterface {
	instance, err := backend.CreateRedisConnection()
	if err != nil {
		panic(errors.Join(FailedToLoadRedisError, err))
	}
	return &ServerApp{
		_redis:     instance,
		ChInternal: make(chan models.InternalMessage, 64),
		isActive:   atomic.Bool{},
	}
}
