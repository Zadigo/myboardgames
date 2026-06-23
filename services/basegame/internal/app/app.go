package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Zadigo/basegame/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
)

type App struct {
	ctx         context.Context
	baseDir     string
	redisClient *redis.Client
	router      *chi.Mux
	errorCh     chan error
}

func (app *App) Start() error {
	log.Printf("🔵 Starting %s HTTP server...\n", os.Getenv("SERVICE_NAME"))

	if app.redisClient == nil {
		return fmt.Errorf("Redis client is not initialized")
	} else {
		log.Println("✅ Connected to Redis successfully")
	}

	if app.ctx == nil {
		return fmt.Errorf("Context is not initialized")
	}

	port, err := strconv.ParseUint(os.Getenv("GO_PORT"), 10, 16)
	if err != nil {
		port = 9000
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: app.router,
	}

	go func() {
		log.Printf("✅ %s HTTP server is running on port %d", os.Getenv("SERVICE_NAME"), port)
		app.errorCh <- server.ListenAndServe()
	}()

	select {
	case err := <-app.errorCh:
		return fmt.Errorf("🔴 %s HTTP server error: %v", os.Getenv("SERVICE_NAME"), err)
	case <-app.ctx.Done():
		timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		app.redisClient.Close()

		log.Printf("🔴 Shutting down %s HTTP server...", os.Getenv("SERVICE_NAME"))
		return server.Shutdown(timeoutCtx)
	}
}

func (app *App) GetBaseDir() string {
	return app.baseDir
}

func (app *App) GetContext() context.Context {
	return app.ctx
}

func (app *App) GetRedisClient() *redis.Client {
	return app.redisClient
}

// NewApp initializes and returns a new instance of the App struct
// with the provided context and base directory. It also sets up the Redis client
// and service registry.
func NewApp(ctx context.Context, baseDir string, options models.AppOptions) models.AppInterface {
	app := &App{
		ctx:     ctx,
		baseDir: baseDir,
		router:  nil,
		errorCh: make(chan error),
	}

	app.redisClient = options.RedisClient
	app.loadRouter()

	return app
}
