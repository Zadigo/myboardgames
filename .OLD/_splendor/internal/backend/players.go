package backend

import (
	"github.com/google/uuid"
)

// PlayerWallet represents the resources a
// player has, including the tokens he was able to
// take and bonuses from cards.
type PlayerWallet struct {
	CardResources
	MarvelCardResources
	MarvelSpecialResouces
}

func (w *PlayerWallet) CanAfford(card CardInterface) bool {
	return false
}

func (w *PlayerWallet) Pay(card CardInterface) {
	// Implement the logic for paying the card's cost
}

func (w *PlayerWallet) AvailableResources(card CardInterface) {
	// Implement the logic for calculating available resources
}

type Player struct {
	PlayerWallet
	Uuid          string          `json:"uuid"`
	Username      string          `json:"username"`
	Points        int             `json:"points"`
	ReservedCards []CardInterface `json:"reservedCards"`
}

type PlayerInterface interface {
	// GetPlayer() *Player
	BuyCard(card CardInterface) error
	ReserveCard(card CardInterface) error
	TakeTokens(emerald int, diamond int, sapphire int, onyx int, ruby int) (int, error)
	ReturnTokens(emerald int, diamond int, sapphire int, onyx int, ruby int) error
	CanBuyCard(card CardInterface) bool
	CanReserveCard() bool
}

// func (p *Player) GetPlayer() *Player {
// 	return p
// }

func (p *Player) BuyCard(card CardInterface) error {
	return nil
}

func (p *Player) ReserveCard(card CardInterface) error {
	return nil
}

func (p *Player) TakeTokens(emerald int, diamond int, sapphire int, onyx int, ruby int) (int, error) {
	return 0, nil
}

func (p *Player) ReturnTokens(emerald int, diamond int, sapphire int, onyx int, ruby int) error {
	return nil
}

func (p *Player) CanBuyCard(card CardInterface) bool {
	return false
}

func (p *Player) CanReserveCard() bool {
	return false
}

// func (p *Player) NumberOfTokens() int {
// 	return p.Emerald + p.Diamond + p.Sapphire + p.Onyx + p.Ruby + p.GoldJoker
// }

// func (p *Player) BuyCard(card *Card) {
// 	if p.CanBuyCard(card) {
// 		p.Points += card.Points
// 		p.Emerald -= card.Emerald
// 		p.Diamond -= card.Diamond
// 		p.Sapphire -= card.Sapphire
// 		p.Onyx -= card.Onyx
// 		p.Ruby -= card.Ruby
// 	}
// }

// func (p *Player) ReserveCard(card *Card) {}

// // Return the number of tokens that need to be returned to the table. If the player has more than 10 tokens,
// // return the number of tokens that need to be returned and true. Otherwise, return 0 and false.
// func (p *Player) TakeTokens(emerald int, diamond int, sapphire int, onyx int, ruby int) (int, bool) {
// 	p.Emerald += emerald
// 	p.Diamond += diamond
// 	p.Sapphire += sapphire
// 	p.Onyx += onyx
// 	p.Ruby += ruby

// 	if p.NumberOfTokens() > 10 {
// 		return p.NumberOfTokens() - 10, true
// 	} else {
// 		return 0, false
// 	}
// }

// func (p *Player) ReturnTokens(emerald int, diamond int, sapphire int, onyx int, ruby int) {
// 	p.Emerald -= emerald
// 	p.Diamond -= diamond
// 	p.Sapphire -= sapphire
// 	p.Onyx -= onyx
// 	p.Ruby -= ruby
// }

// func (p *Player) CanBuyCard(card *Card) bool {
// 	if p.Emerald >= card.Emerald && p.Diamond >= card.Diamond && p.Sapphire >= card.Sapphire && p.Onyx >= card.Onyx && p.Ruby >= card.Ruby {
// 		return true
// 	}
// 	return false
// }

// func (p *Player) CanReserveCard() bool {
// 	if p.NumberOfReservedCards < 3 {
// 		return true
// 	}
// 	return false
// }

func NewPlayer(username string, isNormalGame bool) *Player {
	return &Player{
		Username: username,
		PlayerWallet: PlayerWallet{
			CardResources: CardResources{
				Emerald:  0,
				Diamond:  0,
				Sapphire: 0,
				Onyx:     0,
				Ruby:     0,
			},
			MarvelCardResources: MarvelCardResources{
				Mind:    0,
				Power:   0,
				Reality: 0,
				Soul:    0,
			},
		},
		Uuid:   uuid.NewString(),
		Points: 0,
	}
}
