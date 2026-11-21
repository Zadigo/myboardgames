package internal

import (
	"context"
	"fmt"
	"log/slog"
	"myboardgames/minimalgame/internal/backend"
	"myboardgames/minimalgame/internal/logic"
	"myboardgames/minimalgame/internal/models"
	minimalgame "myboardgames/minimalgame/minimalgameproto"
	"net"
	"time"

	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
)

type App struct {
	options *models.ServerOptions
	router  *chi.Mux
}

func (a *App) Start(ctx context.Context) error {
	slog.Info("Starting server...")

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", a.options.Port))
	if err != nil {
		return fmt.Errorf("Could not start TCP server %s", err)
	}

	defer listener.Close()

	// Game Registry
	registry := logic.NewGameRegistry()

	// Redis
	client, err := backend.CreateRedisConnection(ctx)
	if err != nil {
		return err
	}

	defer client.Close()

	// GrPC server
	grpcServer := grpc.NewServer(
		grpc.MaxRecvMsgSize(1 << 20),
		// grpc.WithChainUnaryInterceptor()
	)
	minimalgame.RegisterTableServiceServer(grpcServer, &GameServices{
		gameRegistry: registry,
		redisClient:  client,
	})

	slog.Info("Server registered")

	errCh := make(chan error, 1)
	go func() {
		errCh <- grpcServer.Serve(listener)
	}()

	select {
	case err := <-errCh:
		return fmt.Errorf("Error running server %s", err)
	case <-ctx.Done():
		done := make(chan struct{})
		go func() { grpcServer.GracefulStop(); close(done) }()

		select {
		case <-done:
		case <-time.After(10 * time.Second):
			grpcServer.Stop()
		}

		slog.Info("server stopped")
		return nil
	}
}

func NewApp(ctx context.Context, options *models.ServerOptions) *App {
	return &App{
		options: options,
	}
}
