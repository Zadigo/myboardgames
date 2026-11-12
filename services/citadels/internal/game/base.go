package game

import (
	"fmt"
	"uuid"

	"github.com/Zadigo/citadels/internal/models"
)

type Game struct {
	models.BaseGame
}

func (g *Game) GetPlayers() []models.PlayerInterface {
	return g.BaseGame.GetPlayers()
}

func (g *Game) Start() error {
	if len(g.GetPlayers()) == 0 {
		return fmt.Errorf("no players")
	}

	if len(g.GetPlayers()) == 2 {
		// TODO: Special rules for 2 players
	}

	if len(g.GetPlayers()) == 3 {
		// TODO: Special rules for 3 players
	}

	if len(g.GetPlayers()) >= 4 && len(g.GetPlayers()) <= 8 {
		// TODO: Normal rules
	}

	return fmt.Errorf("Could not start game")
}

func NewGame() models.GameInterface {
	table := &models.BaseTable{
		ResolutionIndicator: make([]models.BaseCharacterIndicator, 8),
		Crown:               nil,
		Characters:          make([]string, 8),
		Players:             make([]models.PlayerInterface, 0, 8),
		Stage:               models.WaitingForPlayers,
		DistrictPile:        nil,
		Phase:               models.WaitingForPlayers,
	}

	game := &Game{
		ID:    uuid.New().String(),
		Table: table,
	}

	game.Started.Store(true)
	return game
}
