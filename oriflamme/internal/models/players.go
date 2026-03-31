package models

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Player struct {
	Username string
	Conn     *websocket.Conn
	mu       sync.Mutex
	Tokens   int
}

func (player *Player) Prepare(message string) {
	// Each player starts with 1 influence token
	player.Tokens = 1
}

func (player *Player) PlaceCard(cardId string) {

}

// Place a token on a card in the influence queue. 
// The player can only place a token on a card that they own.
func (player *Player) PlaceToken() {

}

// Increase the player's influence tokens by k. This can be used when a player wins tokens from a card.
func (player *Player) IncreaseTokens(k int) {
	player.Tokens += k
}

// Decrease the player's influence tokens by k. This can be used when a player loses tokens from a card.
func (player *Player) DecreaseTokens(k int) {
	if player.Tokens == 0 {
		return
	} else {
		player.Tokens -= k
	}

}

func (player *Player) SendJsonMessage(message string) {}

func (player *Player) ReceiveJsonMessage() string {
	return ""
}
