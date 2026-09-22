package server

import (
	"context"
	"fmt"
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
	httpApp     models.AppInterface
	miniServers map[string]*game.GamesMiniServer
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

	appErrs := make(chan error)

	// Start the HTTP server application
	go func() {
		httpApp := app.NewApp(s.ctx, models.AppOptions{
			RedisClient: s.redisClient,
			ServerApp:   s,
		})
		s.httpApp = httpApp
		appErrs <- httpApp.Start()
	}()

	// Start the game server application
	miniServerErrs := make(chan error)

	// Create two mini-servers: one for the standard game and 
	// one for the extension game. These mini-servers encapsulate
	// all the games of a particular type, allowing for isolated 
	// management and execution of each game variant.

	go func() {
		ctx := context.WithValue(s.ctx, "gameType", games.STANDARD)
		gameCtx, cancel := context.WithCancel(ctx)
		defer cancel()

		gameApp := game.NewGamesMiniServer(
			gameCtx,
			models.GameAppOptions{
				GameType: games.STANDARD,
				Debug:    debug,
				AppOptions: models.AppOptions{
					ServerApp:   s,
					RedisClient: s.redisClient,
				},
			},
		)
		s.miniServers[games.STANDARD] = gameApp
		miniServerErrs <- gameApp.Start()
	}()

	go func() {
		ctx := context.WithValue(s.ctx, "gameType", games.EXTENSION)
		gameCtx, cancel := context.WithCancel(ctx)
		defer cancel()

		gameApp := game.NewGamesMiniServer(
			gameCtx, 
			models.GameAppOptions{
				GameType: games.EXTENSION,
				Debug:    debug,
				AppOptions: models.AppOptions{
					ServerApp:   s,
					RedisClient: s.redisClient,
				},
			},
		)
		s.miniServers[games.EXTENSION] = gameApp
		miniServerErrs <- gameApp.Start()
	}()

	go func() {
		select {
		case miniServerErr, ok := <-miniServerErrs:
			if ok {
				log.Printf("🔴 %s game mini-server error: %v", os.Getenv("SERVICE_NAME"), miniServerErr)
				return
			}
			close(miniServerErrs)
		case appErr, ok := <-appErrs:
			if ok {
				log.Printf("🔴 %s HTTP application error: %v", os.Getenv("SERVICE_NAME"), appErr)
				return
			}
			close(appErrs)
		case <-s.ctx.Done():
			close(appErrs)
			s.redisClient.Close()

			log.Printf("🔴 Shutting down %s server: %v", os.Getenv("SERVICE_NAME"), s.ctx.Err())
			return
		}
	}()

	<-s.ctx.Done()

	log.Printf("🔴 Shutting down %s server: %v", os.Getenv("SERVICE_NAME"), s.ctx.Err())
	return nil
}

func (s *ServerApp) JoinGame(conn *websocket.Conn, gameType string, gameId string) error {
	miniServer := s.miniServers[gameType]
	game, err := miniServer.GetGame(gameId)

	if err != nil {
		return fmt.Errorf("Failed to get game server for gameType %s and gameId %s. %s", gameType, gameId, err)
	}

	err = game.AddPlayer(games.NewWebsocketClient(conn))
	if err != nil {
		return fmt.Errorf("Failed to add player to game %s: %s", gameId, err)
	}

	return nil
}

func (s *ServerApp) CreateGame(gameType string)  (gameApp models.GameInterface, err error) {
	miniServer := s.miniServers[gameType]
	debug := s.ctx.Value("debug").(bool)

	newGame := miniServer.CreateGame(models.GameAppOptions{
		GameType: gameType,
		Debug:    debug,
		AppOptions: models.AppOptions{
			ServerApp:   s,
			RedisClient: s.redisClient,
		},
	})
	return newGame, nil
}

func (s *ServerApp) GetMiniServer(gameType string) *game.GamesMiniServer {
	return s.miniServers[gameType]
}

func (s *ServerApp) GetAllMiniServers() map[string]*game.GamesMiniServer {
	return s.miniServers
}

func NewServerApp(ctx context.Context) *ServerApp {
	return &ServerApp{
		ctx:        ctx,
		miniServers: make(map[string]*game.GamesMiniServer),
	}
}
