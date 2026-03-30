package logic

import (
	"github.com/Zadigo/flipseven/internal/cards"
	"github.com/gorilla/websocket"
)

type Player struct {
	// Player's username
	Username string `json:"username"`
	// Unique identifier for the player
	Uuid string `json:"uuid"`
	// The table asociated with the player (used for
	// broadcasting messages to the right clients)
	TableUuid string `json:"tableUuid"`
	// The current number of cards the player has flipped
	NumberOfCards int `json:"numberOfCards"`
	// Whether the player was stopped by the freeze card
	// or cannot continue flipping cards because he has
	// two similar card numbers
	IsFreezed bool `json:"isFreezed"`
	// Indicates that the player has flipped a freeze card
	HasFreezeCard bool `json:"hasFreezeCard"`
	// Indicates that the player has flipped a second chance card
	HasSecondChanceCard bool `json:"hasSecondChanceCard"`
	// Indicates that a player has flipped seven cards
	// (only cards with numbers) which is the end of the
	// current round
	HasSevenCards bool `json:"hasSevenCards"`
	// Indicates that the player has flipped a
	// second chance card and can continue flipping cards
	HasSecondChance bool `json:"hasSecondChance"`
	// The number of rounds played by the player
	// NumberOfRounds int
	// The cards that the player has hands
	Cards []cards.Card `json:"cards"`
	// The player's current score
	Score int `json:"score"`
}

// Stores the connection and the detailed
// information of player
type ConnectedPlayer struct {
	Conn    *websocket.Conn `json:"-"`
	Details Player          `json:"details"`
}

type PlayersTable struct {
	Clients         []*ConnectedPlayer `json:"clients"`
	CurrentDeck     []cards.Card       `json:"currentDeck"`
	NumberOfPlayers int                `json:"numberOfPlayers"`
	GameStarted     bool               `json:"gameStarted"`
}
