package game

import (
	"context"
	"log"
	"sync/atomic"

	"github.com/Zadigo/basegame/internal/game/games"
	"github.com/Zadigo/basegame/internal/models"
)

type GamesMiniServer struct {
	ctx       context.Context                `json:"-"`
	Games     map[string]games.GameInterface `json:"game"`
	serverApp models.ServerAppInterface      `json:"-"`
	IsRunning atomic.Bool                    `json:"-"`
}

func (app *GamesMiniServer) SetContext(ctx context.Context) {
	app.ctx = ctx
}

// Starts the GamesMiniServer, allowing it to manage multiple game instances concurrently.
func (app *GamesMiniServer) Start() error {
	app.IsRunning.Store(true)

	gameType := app.ctx.Value("gameType").(string)
	log.Printf("🔵 Starting %s game mini-server...", gameType)

	<-app.ctx.Done()

	app.IsRunning.Store(false)
	log.Printf("🔴 Game mini-server for %s is stopping...", gameType)
	return nil
}

// Stops the GamesMiniServer, terminating all managed game instances and cleaning up resources.
func (app *GamesMiniServer) Stop() error {
	app.IsRunning.Store(false)
	return nil
}

// CreateGame creates a new game instance based on the provided options and adds it to the GamesMiniServer's managed games.
func (app *GamesMiniServer) CreateGame(options models.GameAppOptions) games.GameInterface {
	game := games.CreateGame(app.ctx, options)
	app.Games[game.GetUuid()] = game
	return game
}

// GetGame retrieves a game instance by its unique identifier (gameId) from the GamesMiniServer's managed games.
func (app *GamesMiniServer) GetGame(gameId string) (games.GameInterface, error) {
	handler := games.StandardGameError{}

	game, exists := app.Games[gameId]
	if !exists {
		return nil, handler.GameNotFound(gameId)
	}

	return game, nil
}

// GamesMiniServer is a lightweight server that manages multiple game instances for a given type of game (e.g., standard or extension).
// It allows for the creation and management of multiple game instances, each with its own state and players. The server can start and
// stop the game instances as needed, and it provides a centralized point for managing the lifecycle of the games.
func NewGamesMiniServer(ctx context.Context, options models.GameAppOptions) *GamesMiniServer {
	miniServerCtx := context.WithValue(ctx, "gameType", options.GameType)

	return &GamesMiniServer{
		ctx:       miniServerCtx,
		IsRunning: atomic.Bool{},
		serverApp: options.ServerApp,
		Games:     map[string]games.GameInterface{},
	}
}
