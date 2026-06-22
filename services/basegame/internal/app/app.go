package app

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/Zadigo/basegame/internal/game"
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
	appConfig   models.AppConfigInterface
	gameApp     *game.GameApp
}

func (app *App) Start() error {
	log.Printf("🚀 Starting %s service...\n", os.Getenv("SERVICE_NAME"))

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
		log.Printf("✅ %s service is running on port %d", os.Getenv("SERVICE_NAME"), port)
		app.errorCh <- server.ListenAndServe()
	}()

	gameChError := make(chan error)

	go func() {
		log.Printf("🔵 Starting %s game application...", os.Getenv("SERVICE_NAME"))
		gameApp := game.NewGameApp(game.STANDARD)
		app.gameApp = gameApp
		gameChError <- app.gameApp.Start()
	}()

	select {
	case gameErr := <-gameChError:
		log.Printf("🔴 %s game application error: %v", os.Getenv("SERVICE_NAME"), gameErr)
		return gameErr
	case err := <-app.errorCh:
		return fmt.Errorf("🔴 %s service error: %v", os.Getenv("SERVICE_NAME"), err)
	case <-app.ctx.Done():
		timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		app.redisClient.Close()

		log.Printf("🔴 Shutting down %s service...", os.Getenv("SERVICE_NAME"))
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
	// Get the absolute path of the base directory
	absPath, err := filepath.Abs(baseDir)
	if err != nil {
		log.Printf("❌ Failed to get absolute path: %v", err)
		return nil
	}

	result := path.Ext(absPath)
	if result != "" {
		log.Printf("❌ Base directory should be a directory, got a file: %s", absPath)
		return nil
	}

	app := &App{
		ctx:       ctx,
		baseDir:   absPath,
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

	app.loadRouter()

	// Once the base directory is validated, we can walk through it to find the config.yaml file
	filepath.WalkDir(absPath, func(path string, d fs.DirEntry, err error) error {
		if strings.Contains(path, ".yaml") {
			if strings.Contains(path, "config.yaml") {
				// TODO: Load the config.yaml file and initialize the appConfig
				return nil
			}
		}

		return nil
	})

	return app
}
