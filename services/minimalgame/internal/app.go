package internal

import (
	"context"
	"fmt"
	"log/slog"
	pb "myboardgames/gen/minimalgame/v1"
	"myboardgames/minimalgame/internal/logic"
	"myboardgames/minimalgame/internal/models"
	"net"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type App struct {
	options *models.ServerOptions
	router  *chi.Mux
	redisDb *redis.Client
}

func (a *App) Start(ctx context.Context) error {
	slog.Info("Starting server...")

	port, err := strconv.ParseUint(a.options.Port, 10, 16)
	if err != nil {
		return err
	}

	if port == 0 {
		return fmt.Errorf("Invalid port")
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("Could not start TCP server %s", err)
	}

	defer listener.Close()

	// Game Registry
	registry := logic.NewGameRegistry()

	// GrPC server
	grpcServer := grpc.NewServer(
		// Reject messages above 1 MiB (the default is 4 MiB)
		grpc.MaxRecvMsgSize(1<<20),
		grpc.MaxSendMsgSize(1<<20),

		// Cap concurrent streams per connection
		grpc.MaxConcurrentStreams(100),

		// Server-side keepalive: detect dead connections and recycle old ones
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle:     5 * time.Minute,  // close connections idle this long
			MaxConnectionAge:      30 * time.Minute, // force periodic reconnects
			MaxConnectionAgeGrace: 10 * time.Second, // let in-flight RPCs finish
			Time:                  1 * time.Minute,  // ping the client if inactive this long
			Timeout:               20 * time.Second, // close if no ping ack in this time
		}),

		// What the server tolerates from clients
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             30 * time.Second, // clients may not ping more often than this
			PermitWithoutStream: true,             // allow pings with no active RPC
		}),

		// Order matters: recovery goes first so it wraps everything else
		grpc.ChainUnaryInterceptor(
			recoveryUnary, // 1. catch panics from everything below
			loggingUnary,  // 2. log every outcome, including rejections
			// authUnary(&backend.Verifier{}), // 3. who are you?
			// RateLimitUnary(limiter), // 4. needs the caller from auth; cheap, drops floods early
			// ValidationUnary,         // 5. only well-formed requests reach the handler
		),
		grpc.ChainStreamInterceptor(recoveryStream),
	)

	pb.RegisterTableServiceServer(grpcServer, &GameServices{
		gameRegistry: registry,
		redisDb:      a.redisDb,
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

func NewApp(ctx context.Context, redisDb *redis.Client, options *models.ServerOptions) *App {
	return &App{
		redisDb: redisDb,
		options: options,
	}
}
