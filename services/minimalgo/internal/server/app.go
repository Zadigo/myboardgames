package server

import (
	"context"
	"log/slog"
	"myboardgames/minimalgo/internal/games"
	httpapp "myboardgames/minimalgo/internal/http"
	"myboardgames/minimalgo/internal/models"
	"sync/atomic"

	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/errgroup"
)

type ServerApp struct {
	models.BaseServerApp
	options *models.ServerAppOptions
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
	return s.options.RedisClient
}

func (s *ServerApp) Start(ctx context.Context) error {
	// chError := make(chan error)
	// serverCtx := context.WithValue(ctx, "state", true)

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

	// var wg sync.WaitGroup

	// wg.Add(2)
	// wg.Go(func() {
	// 	defer wg.Done()
	// 	httpApp := http.NewHttpApp(ctx, s)
	// 	httpApp.Start(ctx)
	// })

	// wg.Go(func() {
	// 	defer wg.Done()
	// 	gameApp := games.NewGameServerApp(ctx, s)
	// 	gameApp.Start(ctx)
	// })

	// wg.Go(func() {
	// 	defer wg.Done()
	// 	checkDjangoBackend(serverCtx, chError)
	// })

	// <-serverCtx.Done()

	// wg.Wait()

	g, gCtx := errgroup.WithContext(ctx)

	httpApp := httpapp.NewHttpApp(ctx, s)
	gameApp := games.NewGameServerApp(ctx, s)

	g.Go(func() error {
		return httpApp.Start(gCtx)
	})

	g.Go(func() error {
		return gameApp.Start(gCtx)
	})

	s.isActive.Store(true)
	defer s.isActive.Store(false)
	defer func() {
		for _, c := range s.Clients() {
			c.GetConn().Close()
		}
	}()

	return g.Wait()
}

func NewServerApp(ctx context.Context, options *models.ServerAppOptions) models.ServerAppInterface {
	return &ServerApp{
		options:    options,
		ChInternal: make(chan models.InternalMessage, 64),
		isActive:   atomic.Bool{},
	}
}
