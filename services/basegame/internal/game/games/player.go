package games

import (
	"github.com/Zadigo/basegame/internal/models"
)

type GameSpecific struct {
	// Indicates that the player was stopped by the freeze card
	// or cannot continue flipping cards because he has
	// two similar card numbers
	IsFrozen bool `json:"isFreezed"`
	// The score of the player in each round played by the player
	// Scores []int `json:"scores"`
	Cards []models.CardInterface `json:"cards"`
}

func (g *GameSpecific) GetNumberOfCards() int {
	return len(g.Cards)
}

type Player struct {
	GameSpecific
	// Username of the player
	Username string `json:"username"`
	// Indicates that the player is the initiator of the game
	IsInitiator bool `json:"isInitiator"`
}

func (p *Player) GetCards() []models.CardInterface {
	return p.Cards
}

func (p *Player) SetCards(cards []models.CardInterface) {
	p.Cards = cards
}

func (p *Player) GetTotalScore() int {
	total := 0
	for _, card := range p.Cards {
		total += card.Resolve(total, models.CardResolveOptions{})
	}
	return total
}

func (p *Player) GetIsFrozen() bool {
	return p.IsFrozen
}

func (p *Player) SetIsFrozen(frozen bool) {
	p.IsFrozen = frozen
}
