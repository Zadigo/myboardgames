package game

import "github.com/Zadigo/citadels/internal/models"

type Game struct {
	models.BaseGame
}

func (g *Game) GetPlayers() []models.Player {
	return nil
}

func NewGame() models.GameInterface {
	return &Game{}
}
