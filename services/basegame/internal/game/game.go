package game

import (
	"context"
	"time"

	"github.com/Zadigo/basegame/internal/game/redis"
	"github.com/go-co-op/gocron"
)

type BaseGame struct {
	ctx     context.Context
	table   map[string]*Table
	players map[string]*Player
}

func (b *BaseGame) Start() error {
	scheduler := gocron.NewScheduler(time.UTC)

	// Start a taks that will run every X seconds in order
	// to check if the game is still active, that the players
	// are still connected, etc.
	_, err := scheduler.Every(60).Second().Do(func() {})

	if err != nil {
		// Handle error
		return err
	}

	scheduler.StartBlocking()
	return nil
}

func (b *BaseGame) CreateTable() error {
	return nil
}

func (b *BaseGame) CreateCards() error {
	redisHandler := redis.NewGameRedis(b.ctx, nil)
	redisHandler.GetCards()
	return nil
}

func (b *BaseGame) Prepare() error {
	return nil
}

type StandardGame struct {
	BaseGame
}

// NewStandardGame creates a standard game with standard rules and configurations.
//  It initializes the game state, sets up the necessary components, and prepares the game 
// for players to join and interact with.
func NewStandardGame() *StandardGame {
	return &StandardGame{
		BaseGame: BaseGame{
			table:   make(map[string]*Table),
			players: make(map[string]*Player),
		},
	}
}

type ExtensionGame struct {
	BaseGame
}

// NewExtensionGame creates an extension game with additional rules and configurations.
// It initializes the game state, sets up the necessary components, and prepares the game 
// for players to join and interact with, including any special features or mechanics.
func NewExtensionGame() *ExtensionGame {
	return &ExtensionGame{
		BaseGame: BaseGame{
			table:   make(map[string]*Table),
			players: make(map[string]*Player),
		},
	}
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

// CreateGame is a factory function that creates a new game instance based on the specified game type.
// It returns an instance of GameInterface, which can be either a StandardGame or an ExtensionGame,
// depending on the provided gameType. If the gameType is not recognized, it returns nil.
func CreateGame(gameType string) GameInterface {
	switch gameType {
	case STANDARD:
		return NewStandardGame()
	case EXTENSION:
		return NewExtensionGame()
	default:
		return nil
	}
}
