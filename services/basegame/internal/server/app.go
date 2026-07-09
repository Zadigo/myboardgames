package server

import (
	"context"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Zadigo/basegame/internal/app"
	"github.com/Zadigo/basegame/internal/game"
	"github.com/Zadigo/basegame/internal/game/games"
	"github.com/Zadigo/basegame/internal/models"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

// ServerApp is the main server application that
// synchronizes all the servers (chi HTTP server, game server)
// and serves as mediator between them. It is responsible for initializing
// and managing the lifecycle of the servers, handling requests, and coordinating
// communication between different components of the application.
type ServerApp struct {
	// Parent context
	ctx         context.Context
	rootDir     string
	redisClient *redis.Client
	GameApp     *game.GameApp
	httpApp     models.AppInterface
}

func (s *ServerApp) Start() error {
	rootDir := s.ctx.Value("baseDir").(string)
	debug := s.ctx.Value("debug").(bool)

	// Once the base directory is validated, we can walk through it to find the config.yaml file
	filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
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
		httpApp := app.NewApp(s.ctx, models.AppOptions{
			RedisClient: s.redisClient,
			ServerApp:   s,
		})
		s.httpApp = httpApp
		appErrors <- httpApp.Start()
	}()

	// Start the game server application
	gameErrors := make(chan error)

	go func() {
		log.Printf("🔵 Starting %s standard game application...", os.Getenv("SERVICE_NAME"))

		ctx := context.WithValue(s.ctx, "gameType", games.STANDARD)
		cancelCtx, cancel := context.WithCancel(ctx)
		defer cancel()

		gameApp := game.NewGameApp(cancelCtx, models.GameAppOptions{
			GameType: games.STANDARD,
			Debug:    debug,
			AppOptions: models.AppOptions{
				ServerApp:   s,
				RedisClient: s.redisClient,
			},
		})
		gameErrors <- gameApp.Start()
	}()

	go func() {
		log.Printf("🔵 Starting %s extension game application...", os.Getenv("SERVICE_NAME"))

		ctx := context.WithValue(s.ctx, "gameType", games.EXTENSION)
		cancelCtx, cancel := context.WithCancel(ctx)
		defer cancel()

		gameApp := game.NewGameApp(cancelCtx, models.GameAppOptions{
			GameType: games.EXTENSION,
			Debug:    debug,
			AppOptions: models.AppOptions{
				ServerApp:   s,
				RedisClient: s.redisClient,
			},
		})
		gameErrors <- gameApp.Start()
	}()

	select {
	case gameErr := <-gameErrors:
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

func (s *ServerApp) JoinGame(conn *websocket.Conn, gameUUID string) error {
	return s.GameApp.AddPlayer(games.NewWebsocketClient(conn))
}

func (s *ServerApp) CreateGame(gameType string) error {
	s.GameApp.Create(models.GameAppOptions{
		GameType: gameType,
		AppOptions: models.AppOptions{
			ServerApp:   s,
			RedisClient: s.redisClient,
		},
	})
	return nil
}

func NewServerApp(ctx context.Context) *ServerApp {
	return &ServerApp{
		ctx: ctx,
	}
}
