package server

import (
	"context"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/Zadigo/basegame/internal/app"
	"github.com/Zadigo/basegame/internal/game"
	"github.com/Zadigo/basegame/internal/models"
	"github.com/redis/go-redis/v9"
)

// ServerApp is the main server application that
// synchronizes all the servers (chi HTTP server, game server)
// and serves as mediator between them. It is responsible for initializing
// and managing the lifecycle of the servers, handling requests, and coordinating
// communication between different components of the application.
type ServerApp struct {
	ctx         context.Context
	rootDir     string
	redisClient *redis.Client
}

func (s *ServerApp) Start(rootDir string) error {
	absPath, err := filepath.Abs(rootDir)
	if err != nil {
		log.Printf("❌ Failed to get absolute path: %v", err)
		return nil
	}

	result := path.Ext(absPath)
	if result != "" {
		log.Printf("❌ Base directory should be a directory, got a file: %s", absPath)
		return nil
	}

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

	// Setup Redis for the whole server
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	s.redisClient = redisClient

	cmd := s.redisClient.Ping(s.ctx)
	if cmd.Err() != nil {
		panic(cmd.Err())
	}

	appErrors := make(chan error)

	// Start the HTTP server application
	go func() {
		httpApp := app.NewApp(s.ctx, absPath, models.AppOptions{
			RedisClient: s.redisClient,
			ServerApp:   s,
		})
		appErrors <- httpApp.Start()
	}()

	// Start the game server application
	gameChError := make(chan error)

	go func() {
		log.Printf("🔵 Starting %s game application...", os.Getenv("SERVICE_NAME"))
		gameApp := game.NewGameApp(s.ctx, absPath, models.GameAppOptions{
			GameType: game.STANDARD,
			AppOptions: models.AppOptions{
				ServerApp: s,
			},
		})
		gameChError <- gameApp.Start()
	}()

	select {
	case gameErr := <-gameChError:
		log.Printf("🔴 %s game application error: %v", os.Getenv("SERVICE_NAME"), gameErr)
		return gameErr
	case appErr := <-appErrors:
		log.Printf("🔴 %s HTTP application error: %v", os.Getenv("SERVICE_NAME"), appErr)
		return appErr
	case <-s.ctx.Done():
		s.redisClient.Close()

		log.Printf("🔴 Shutting down %s server...", os.Getenv("SERVICE_NAME"))
		return nil
	}
}

func NewServerApp(ctx context.Context) *ServerApp {
	return &ServerApp{
		ctx: ctx,
	}
}
