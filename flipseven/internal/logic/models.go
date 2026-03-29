package logic

import (
	"github.com/Zadigo/flipseven/internal/cards"
	"github.com/gorilla/websocket"
)

type Player struct {
	// Player's username
	Username string
	// Unique identifier for the player
	Uuid string
	// The table asociated with the player (used for 
	// broadcasting messages to the right clients)
	TableUuid string
	// The current number of cards the player has flipped
	NumberOfCards int
	// Whether the player was stopped by the freeze card
	// or cannot continue flipping cards because he has
	// two similar card numbers
	IsFreezed bool
	// Indicates that a player has flipped seven cards 
	// (only cards with numbers) which is the end of the 
	// current round
	HasSevenCards bool
	// Indicates that the player has flipped a
	// second chance card and can continue flipping cards
	HasSecondChance bool
	// The number of rounds played by the player
	// NumberOfRounds int
	// The cards that the player has hands
	Cards []cards.Card
	// The player's current score
	Score int
}

// Stores the connection and the detailed
// information of player
type ConnectedPlayer struct {
	Conn    *websocket.Conn
	Details Player
}

type PlayersTable struct {
	Clients         []*ConnectedPlayer
	CurrentDeck     []cards.Card
	NumberOfPlayers int
	GameStarted     bool
}
