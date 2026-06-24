package games

import "github.com/Zadigo/basegame/internal/game/cards"

type Player struct {
	Username string
}

func (p *Player) GetCards() []cards.CardInterface {
	return nil
}
