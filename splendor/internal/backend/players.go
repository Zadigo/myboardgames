package backend

import (
	"sync"

	"github.com/gorilla/websocket"
)

// PlayerWallet represents the resources a
// player has, including the tokens and bonuses
// from cards.
type PlayerWallet struct {
	// IsNormalGame indicates whether the game is a normal
	// game or a Marvel expansion game.
	IsNormalGame bool
	CardResources
	MarvelCardResources
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
	Uuid          string
	Points        int
	ReservedCards []CardInterface

	// Emerald               int
	// Diamond               int
	// Sapphire              int
	// Onyx                  int
	// Ruby                  int
	// GoldJoker             int
}

type PlayerInterface interface {
	BuyCard(card CardInterface)
	ReserveCard(card CardInterface)
	TakeTokens(emerald int, diamond int, sapphire int, onyx int, ruby int) (int, bool)
	ReturnTokens(emerald int, diamond int, sapphire int, onyx int, ruby int)
	CanBuyCard(card CardInterface) bool
	CanReserveCard() bool
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

// WebsocketClient represents a client connected 
// to the game via websocket. It wraps the player's 
// information and the websocket connection.
type WebsocketClient struct {
	Player *Player
	Conn   *websocket.Conn
	Mu     *sync.RWMutex
}

func NewWebsocketClient(player *Player, conn *websocket.Conn) *WebsocketClient {
	return &WebsocketClient{
		Player: player,
		Conn:   conn,
		Mu:     &sync.RWMutex{},
	}
}
