package games

import "github.com/Zadigo/basegame/internal/game/cards"

type GameSpecific struct {
	// Indicates that the player was stopped by the freeze card
	// or cannot continue flipping cards because he has
	// two similar card numbers
	IsFrozen bool `json:"isFreezed"`
	// The score of the player in each round played by the player
	// Scores []int `json:"scores"`
	Cards []cards.CardInterface `json:"cards"`
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

func (p *Player) GetCards() []cards.CardInterface {
	return p.Cards
}

func (p *Player) GetTotalScore() int {
	total := 0
	for _, card := range p.Cards {
		total += card.Resolve(total, cards.CardResolveOptions{})
	}
	return total
}
