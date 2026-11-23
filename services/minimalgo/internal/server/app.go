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
	"os"
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

func (s *ServerApp) GetGrpcConnection(id string) (*grpc.ClientConn, error) {
	conn, ok := s.grpcConnections[id]
	if !ok {
		return nil, RpcConnectionNotFoundError
	}
	return conn, nil
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

	server := &ServerApp{
		options:         options,
		grpcRouter:      options.Router,
		grpcConnections: make(map[string]*grpc.ClientConn),
	}

	for _, r := range options.Router.Routes {
		var conn *grpc.ClientConn
		var err error

		if utils.IsProduction() {
			conn, err = grpcclient.NewSecuredGrpcClient(r.Address(), &tls.Config{}, priv, "key-1")
		} else {
			conn, err = grpcclient.NewInsecureGrpcClient(r.Address())
			slog.Warn("Using RPC server in insecure mode", "mode", os.Getenv("GRPC_ENV"))
		}

		if err != nil {
			slog.Error("Failed to create rpc connection", "game", r.NameEn, "err", err)
		} else {
			slog.Info("Created RPC connection", "game", r.NameEn, "address", r.Address())
		}

		server.grpcConnections[r.Hash] = conn
	}

	return server, nil
}
