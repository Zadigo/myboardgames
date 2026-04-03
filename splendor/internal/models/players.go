package models

import (
	"github.com/gorilla/websocket"
)

type Player struct {
	Uuid                  string
	Points                int
	Emerald               int
	Diamond               int
	Sapphire              int
	Onyx                  int
	Ruby                  int
	GoldJoker             int
	NumberOfReservedCards int
	conn                  websocket.Conn
}

func (p *Player) NumberOfTokens() int {
	return p.Emerald + p.Diamond + p.Sapphire + p.Onyx + p.Ruby + p.GoldJoker
}

func (p *Player) BuyCard(card *Card) {
	if p.CanBuyCard(card) {
		p.Points += card.Points
		p.Emerald -= card.Emerald
		p.Diamond -= card.Diamond
		p.Sapphire -= card.Sapphire
		p.Onyx -= card.Onyx
		p.Ruby -= card.Ruby
	}
}

func (p *Player) ReserveCard(card *Card) {}

// Return the number of tokens that need to be returned to the table. If the player has more than 10 tokens,
// return the number of tokens that need to be returned and true. Otherwise, return 0 and false.
func (p *Player) TakeTokens(emerald int, diamond int, sapphire int, onyx int, ruby int) (int, bool) {
	p.Emerald += emerald
	p.Diamond += diamond
	p.Sapphire += sapphire
	p.Onyx += onyx
	p.Ruby += ruby

	if p.NumberOfTokens() > 10 {
		return p.NumberOfTokens() - 10, true
	} else {
		return 0, false
	}
}

func (p *Player) ReturnTokens(emerald int, diamond int, sapphire int, onyx int, ruby int) {
	p.Emerald -= emerald
	p.Diamond -= diamond
	p.Sapphire -= sapphire
	p.Onyx -= onyx
	p.Ruby -= ruby
}

func (p *Player) CanBuyCard(card *Card) bool {
	if p.Emerald >= card.Emerald && p.Diamond >= card.Diamond && p.Sapphire >= card.Sapphire && p.Onyx >= card.Onyx && p.Ruby >= card.Ruby {
		return true
	}
	return false
}

func (p *Player) CanReserveCard() bool {
	if p.NumberOfReservedCards < 3 {
		return true
	}
	return false
}
