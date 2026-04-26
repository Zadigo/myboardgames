package logic

import (
	"sync"

	"github.com/google/uuid"
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
	// Cumulative score across all rounds
	// played by the player
	TotalScore int `json:"totalScore"`
	// Indicates that the player is the initiator of the game
	IsInitiator bool `json:"isInitiator"`
	// The player's websocket connection, not serialized to JSON
	Conn *websocket.Conn `json:"-"`
	// Mutex to protect access to the player's data, not serialized to JSON
	Mu sync.Mutex `json:"-"`
}

func (p *Player) GetCards(queue *CardQueue) []*Card {
	cards := []*Card{}
	for _, card := range queue.Queue {
		if card.Owner != nil && card.Owner.Uuid == p.Uuid {
			cards = append(cards, card)
		}
	}
	return cards
}

func (p *Player) GetCurrentScore(queue *CardQueue) int {
	currentHand := p.GetCards(queue)

	score := 0
	for _, card := range currentHand {
		score += card.Value
	}
	return score
}

func (p *Player) HasSevenCards() bool {
	return p.NumberOfCards == 7
}

func (p *Player) SetFlipCard(selectedCard *Card) {
	selectedCard.Owner = p
}

// Returns a new player with the given username and table UUID, and initializes
// the player's attributes to their default values.
func CreatePlayer(connection *websocket.Conn, username string, tableUuid string, isInitiator bool) *Player {
	return &Player{
		Username:      username,
		Uuid:          uuid.NewString(),
		TableUuid:     tableUuid,
		NumberOfCards: 0,
		IsFreezed:     false,
		TotalScore:    0,
		IsInitiator:   isInitiator,
		Conn:          connection,
		Mu:            sync.Mutex{},
	}
}
