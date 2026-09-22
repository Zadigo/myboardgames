package cards

import "github.com/Zadigo/basegame/internal/models"

// CardPile are the collection of cards that are available to be drawn
// by the player. The pile is typically shuffled at the start of the game,
// and cards are drawn from the top of the pile. When a card is drawn, it is
// removed from the pile and added to the player's hand or played directly,
// depending on the game rules.
type CardPile struct {
	Cards []models.CardInterface `json:"cards"`
}

func (p *CardPile) AddCards(card ...models.CardInterface) {
	p.Cards = append(p.Cards, card...)
}

func (p *CardPile) DrawCard() models.CardInterface {
	if len(p.Cards) == 0 {
		return nil
	}

	card := p.Cards[0]
	p.Cards = p.Cards[1:]

	return card
}

func (p *CardPile) RemainingCards() int {
	return len(p.Cards)
}

func (p *CardPile) Shuffle() {}

func (p *CardPile) Clear() {
	p.Cards = []models.CardInterface{}
}

func NewCardPile() *CardPile {
	return &CardPile{
		Cards: []models.CardInterface{},
	}
}
