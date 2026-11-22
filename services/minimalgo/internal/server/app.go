package server

import (
	"context"
	"crypto/tls"
	"log/slog"
	"myboardgames/minimalgo/internal/backend"
	httpapp "myboardgames/minimalgo/internal/http"
	"myboardgames/minimalgo/internal/models"
	"sync/atomic"
	"time"

	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

type ServerApp struct {
	models.BaseServerApp
	options         *models.ServerAppOptions
	isActive        atomic.Bool
	grpcRouter      *backend.Router
	grpcConnections map[string]*grpc.ClientConn
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
	router := &backend.Router{}
	router.SetRedis(options.RedisClient)

	if err := router.LoadFromCache(ctx, router); err != nil {
		return nil, err
	} else {
		server := &ServerApp{
			options:         options,
			grpcRouter:      router,
			grpcConnections: make(map[string]*grpc.ClientConn),
		}

		for _, r := range router.Routes {
			var _credentials credentials.TransportCredentials
			if r.Environment == "prod" {
				_credentials = credentials.NewTLS(&tls.Config{})
			} else {
				_credentials = insecure.NewCredentials()
			}

			conn, err := grpc.NewClient(
				r.Address(),
				grpc.WithTransportCredentials(_credentials),
				// grpc.WithPerRPCCredentials(_credentials),
				grpc.WithKeepaliveParams(keepalive.ClientParameters{
					Time:                time.Minute, // must be >= the server's MinTime (30s)
					Timeout:             20 * time.Second,
					PermitWithoutStream: true,
				}),
			)
			if err != nil {
				return nil, err
			}

			server.grpcConnections[r.Hash] = conn
			state := conn.GetState().String()
			
			switch state {
			case "IDLE":
				slog.Error("Could not connect to rpc server", "address", r.Address())
			case "READY":
				go func() {
					err := conn.Close()
					slog.Error("Failed to close gRPC connection", "error", err)
				}()
				slog.Info("Connected to gRPC", "address", r.Address())
			}
		}

		return server, nil
	}
}
