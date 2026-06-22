package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Zadigo/flipseven/internal/game"
	"github.com/Zadigo/flipseven/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
)

type App struct {
	ctx         context.Context
	baseDir     string
	redisClient *redis.Client
	router      *chi.Mux
	errorCh     chan error
	appConfig   models.AppConfigInterface
	gameApp     *game.GameApp
}

func (app *App) Start() error {
	log.Print("🚀 Starting Goemailer service...")

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
		return fmt.Errorf("🔴 Invalid port: %w", err)
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: app.router,
	}

	go func() {
		log.Printf("✅ Goemailer service is running on port %d", port)
		app.errorCh <- server.ListenAndServe()
	}()

	select {
	case err := <-app.errorCh:
		return fmt.Errorf("HTTP server error: %v", err)
	case <-app.ctx.Done():
		timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		app.redisClient.Close()

		log.Println("🔴 Shutting down Goemailer service...")
		return server.Shutdown(timeoutCtx)
	}
}

func (app *App) GetBaseDir() string {
	return app.baseDir
}

func (app *App) GetContext() context.Context {
	return app.ctx
}

func (app *App) GetConfig() models.AppConfigInterface {
	return app.appConfig
}

func (app *App) GetRedisClient() *redis.Client {
	return app.redisClient
}

// NewApp initializes and returns a new instance of the App struct
// with the provided context and base directory. It also sets up the Redis client
// and service registry.
func NewApp(ctx context.Context, baseDir string) models.AppInterface {
	app := &App{
		ctx:       ctx,
		baseDir:   baseDir,
		router:    nil,
		errorCh:   make(chan error),
		appConfig: &AppConfig{},
		gameApp:   nil,
	}

	app.redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	cmd := app.redisClient.Ping(ctx)
	if cmd.Err() != nil {
		panic(cmd.Err())
	}

	go func() {
		log.Println("🔵 Starting game service...")
		gameApp := game.NewGameApp(game.STANDARD)
		app.gameApp = gameApp
		app.errorCh <- gameApp.Start()
	}()

	app.loadRouter()

	return app
}
