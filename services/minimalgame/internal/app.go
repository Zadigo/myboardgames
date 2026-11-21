package internal

import (
	"context"
	"fmt"
	"log/slog"
	"myboardgames/minimalgame/internal/models"
	minimalgame "myboardgames/minimalgame/minimalgameproto"
	"net"

	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
)

type App struct {
	options *models.ServerOptions
	router  *chi.Mux
	errCh   chan error
}

func (a *App) Start(ctx context.Context) error {
	slog.Info("Starting server...")

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", a.options.Port))
	if err != nil {
		return fmt.Errorf("Could not start TCP server %s", err)
	}

	defer listener.Close()

	grpcServer := grpc.NewServer()
	minimalgame.RegisterTableServiceServer(grpcServer, &Server{})

	slog.Info("Server registered")
	a.errCh <- grpcServer.Serve(listener)

	select {
	case err := <-a.errCh:
		return fmt.Errorf("Error running server %s", err)
	case <-ctx.Done():
		grpcServer.GracefulStop()
		close(a.errCh)

		slog.Info("Server stopped")
		return nil
	}
}

func NewApp(ctx context.Context, options *models.ServerOptions) *App {
	return &App{
		options: options,
		errCh:   make(chan error, 32),
	}
}
