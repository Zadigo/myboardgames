package game

import "github.com/Zadigo/citadels/internal/models"

type Game struct {
	models.BaseGame
}

func (g *Game) GetPlayers() []models.PlayerInterface {
	return g.BaseGame.GetPlayers()
}

func NewGame() models.GameInterface {
	return &Game{}
}
