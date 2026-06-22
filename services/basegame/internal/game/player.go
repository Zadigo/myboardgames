package game

import "github.com/Zadigo/basegame/internal/models"

type Player struct {
	Username string
}

func (p *Player) GetCards() []models.BaseCardInterface {
	return nil
}
