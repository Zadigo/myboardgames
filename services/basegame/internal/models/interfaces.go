package models

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type AppConfigInterface interface {
	LoadConfig(app AppInterface) error
}

type AppExtraInterface interface {
	GetRedisClient() *redis.Client
	GetBaseDir() string
}

type AppInterface interface {
	AppExtraInterface
	Start() error
	GetContext() context.Context
	GetConfig() AppConfigInterface
	// GetGameApp() *game.GameApp
}

type GameInterface interface {
	// Start initializes and starts the game server application.
	// It checks if the game instance is properly initialized, prepares the game state,
	// and then starts the game logic.
	Start() error
	// CreateTable sets up a new game table, initializing the necessary data structures
	// and state for a new game session. This method is responsible for ensuring that the
	// game environment is ready for players to join and interact with the game.
	CreateTable() error
	// CreateCards initializes the card deck for the game,
	// setting up the necessary cards and their properties. This method is
	// responsible for ensuring that the game has a complete set of cards ready for play,
	// and it may involve shuffling or organizing the cards as needed.
	CreateCards() error
	// Prepare sets up the game state, initializes necessary components,
	// and ensures that the game is ready to start. This method is called before
	// the game begins and can include tasks such as shuffling cards, assigning
	// players to tables, and setting initial game parameters.
	Prepare() error
}

type BaseCardInterface interface {
	// Resolve defines the behavior of a card when it is played or activated in the game.
	// Each card type will implement its own logic for this method.
	Resolve()
	AttribuRteCard(player any)
}
