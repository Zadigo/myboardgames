package server

import (
	"context"
	"log/slog"
	httpapp "myboardgames/minimalgo/internal/http"
	"myboardgames/minimalgo/internal/models"
	"sync/atomic"

	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/errgroup"
)

type ServerApp struct {
	models.BaseServerApp
	options *models.ServerAppOptions
	isActive              atomic.Bool
}

func (s *ServerApp) GetState() bool {
	return s.isActive.Load()
}

func (s *ServerApp) GetRedis() *redis.Client {
	return s.options.RedisClient
}

func (s *ServerApp) GetOptions() *models.ServerAppOptions {
	return s.options
}

func (s *ServerApp) Start(ctx context.Context) error {
	slog.Info("Initializing main server")

	g, gCtx := errgroup.WithContext(ctx)

	httpApp := httpapp.NewHttpApp(ctx, s)

	g.Go(func() error {
		return httpApp.Start(gCtx)
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
		options:               options,
		isActive:              atomic.Bool{},
	}
}
