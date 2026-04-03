package models

import (
	"slices"
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
	Cards []*Card `json:"cards"`
	// The player's current score
	Score int `json:"score"`
	// Indicates that the player is the initiator of the game
	IsInitiator bool `json:"isInitiator"`
	// The player's websocket connection, not serialized to JSON
	conn *websocket.Conn `json:"-"`
	// Mutex to protect access to the player's data, not serialized to JSON
	mu sync.Mutex `json:"-"`
}

// Calculates the player's score based on the cards they have in their hand. It
// first sums up the values of all number and bonus cards, and then applies any multipliers
// to the total score. The final score is stored in the player's Score attribute.
func (p *Player) CalculatePoints(card *Card) {
	intermediateScore := 0

	for _, card := range p.Cards {
		if card.IsNumber || card.IsBonus {
			intermediateScore += card.Value
		}
	}

	for _, card := range p.Cards {
		if card.IsMultiplier {
			intermediateScore *= card.Value
		}
	}

	p.Score = intermediateScore
}

// Set the cards that are owned by the player in the deck
// to the players hand. This is used to update the player's
// hand after flipping a card.
func (p *Player) SetPlayerCards(t *TableLayer) {
	for _, card := range t.Deck {
		if card.Owner != nil && card.Owner.Uuid == p.Uuid {
			if slices.Contains(p.Cards, card) {
				continue
			} else {
				p.Cards = append(p.Cards, card)
			}
		}
	}
}

// Returns a new player with the given username and table UUID, and initializes
// the player's attributes to their default values.
func CreatePlayer(username string, tableUuid string, isInitiator bool) *Player {
	return &Player{
		Username:            username,
		Uuid:                uuid.NewString(),
		TableUuid:           tableUuid,
		NumberOfCards:       0,
		IsFreezed:           false,
		HasSecondChanceCard: false,
		HasSevenCards:       false,
		Cards:               []*Card{},
		Score:               0,
		IsInitiator:         isInitiator,
	}
}
