package server

import (
	"context"
	"crypto/ed25519"
	"crypto/tls"
	"log/slog"
	"myboardgames/minimalgo/internal/grpcclient"
	httpapp "myboardgames/minimalgo/internal/http"
	"myboardgames/minimalgo/internal/models"
	"myboardgames/minimalgo/internal/utils"
	"sync/atomic"

	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type ServerApp struct {
	models.BaseServerApp
	options         *models.ServerAppOptions
	isActive        atomic.Bool
	grpcRouter      *models.Router
	grpcConnections map[string]*grpc.ClientConn
}

func (s *ServerApp) GetRouter() *models.Router {
	return s.grpcRouter
}

func (s *ServerApp) GetGrpcConnections() map[string]*grpc.ClientConn {
	return s.grpcConnections
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

func NewServerApp(ctx context.Context, options *models.ServerAppOptions) (models.ServerAppInterface, error) {
	// priv, err := grpcclient.LoadPrivateKey(os.Getenv("GRPC_SIGNING_KEY_PATH"))
	// if err != nil && utils.IsProduction() {
	// 	return nil, err
	// } else {
	// 	slog.Error("Failed to load signing key", "err", err)
	// }
	var priv ed25519.PrivateKey

	if router, err := models.NewRouter(ctx, options.RedisClient); err != nil {
		return nil, err
	} else {
		server := &ServerApp{
			options:         options,
			grpcRouter:      router,
			grpcConnections: make(map[string]*grpc.ClientConn),
		}

		for _, r := range router.Routes {
			var conn *grpc.ClientConn
			var err error

			if utils.IsProduction() {
				conn, err = grpcclient.NewSecuredGrpcClient(r.Address(), &tls.Config{}, priv, "key-1")
			} else {
				conn, err = grpcclient.NewInsecureGrpcClient(r.Address())
			}

			server.grpcConnections[r.Hash] = conn
			if err != nil {
				slog.Error("Failed to create rpc connection", "game", r.NameEn, "err", err)
			} else {
				slog.Info("Creating rpc connection", "game", r.NameEn, "address", r.Address())
			}
		}

		return server, nil
	}
}
