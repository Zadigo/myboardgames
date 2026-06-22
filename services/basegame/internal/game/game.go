package game

import (
	"time"

	"github.com/go-co-op/gocron"
)

type BaseGame struct {
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

func NewExtensionGame() *ExtensionGame {
	return &ExtensionGame{
		BaseGame: BaseGame{
			table:   make(map[string]*Table),
			players: make(map[string]*Player),
		},
	}
}

type GameInterface interface {
	Start() error
	CreateTable() error
	CreateCards() error
	Prepare() error
}

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
