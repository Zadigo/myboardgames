package models

import (
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type ExtensionPlayer struct {
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
	// HasFreezeCard bool `json:"hasFreezeCard"`
	// Indicates that the player has flipped a second chance card
	HasSecondChanceCard bool `json:"hasSecondChanceCard"`
	// Indicates that a player has flipped seven cards
	// (only cards with numbers) which is the end of the
	// current round
	HasSevenCards bool `json:"hasSevenCards"`
	// Indicates that the player has flipped a
	// second chance card and can continue flipping cards
	// HasSecondChance bool `json:"hasSecondChance"`
	// The number of rounds played by the player
	// NumberOfRounds int
	// The cards that the player has in hands
	Cards []Card `json:"cards"`
	// The player's current score
	Score int `json:"score"`
	// Indicates that the player is the initiator of the game
	IsInitiator bool `json:"isInitiator"`
	// The player's websocket connection, not serialized to JSON
	conn *websocket.Conn `json:"-"`
	// Mutex to protect access to the player's data, not serialized to JSON
	mu sync.Mutex `json:"-"`
}

func (p *ExtensionPlayer) AddCard(card Card) {
}

func (p *ExtensionPlayer) CalculatePoints(card Card) {
}

func (p *ExtensionPlayer) RemoveCard(conn *websocket.Conn) {
}

// Returns a new player with the given username and table UUID, and initializes
// the player's attributes to their default values.
func CreateExtensionPlayer(username string, tableUuid string, isInitiator bool) *ExtensionPlayer {
	return &ExtensionPlayer{
		Username:            username,
		Uuid:                uuid.NewString(),
		TableUuid:           tableUuid,
		NumberOfCards:       0,
		IsFreezed:           false,
		HasSecondChanceCard: false,
		HasSevenCards:       false,
		Cards:               []Card{},
		Score:               0,
		IsInitiator:         isInitiator,
	}
}
