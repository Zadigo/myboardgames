package game

import (
	"context"
	"time"

	"github.com/Zadigo/basegame/internal/models"
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
	return nil
}

func (b *BaseGame) Prepare() error {
	return nil
}

type StandardGame struct {
	BaseGame
}

// NewStandardGame creates a standard game with standard rules and configurations.
// It initializes the game state, sets up the necessary components, and prepares the game
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

// CreateGame is a factory function that creates a new game instance based on the specified game type.
// It returns an instance of GameInterface, which can be either a StandardGame or an ExtensionGame,
// depending on the provided gameType. If the gameType is not recognized, it returns nil.
func CreateGame(gameType string) models.GameInterface {
	switch gameType {
	case STANDARD:
		return NewStandardGame()
	case EXTENSION:
		return NewExtensionGame()
	default:
		return nil
	}
}
