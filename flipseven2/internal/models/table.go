package models

import (
	"errors"

	"github.com/Zadigo/flipseven2/internal/backend/broadcasting"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type TableLayerInterface interface {
	FlipCard(player *Player) (string, *Card, error)
	StartGame()
	EndGame()
	FreezePlayer(player *Player)
	AddPlayer(connection *websocket.Conn, username string, isInitiator bool) *Player
	ResetPlayers()
	NextPlayer(broadcastingRegistry *broadcasting.BroadcasterRegistry)
	GetUuid() string
	GetCurrentCard() (*Card, error)
	GetDeck() []*Card
	NumberOfCards() int
	HasPlayer(connection *websocket.Conn) bool
	GetPlayer(connection *websocket.Conn) *Player
}

type TableLayer struct {
	// Unique identifier for the table
	Uuid string `json:"uuid"`
	// The deck of cards for the current game, which is a slice of pointers to Card structs. It is initialized as an empty slice and will be populated with the appropriate cards when the game starts.
	Deck []*Card `json:"deck"`
	// The index of the current card in the deck. It is initialized to -1 to indicate that no card has been flipped yet.
	DeckIndex int `json:"deckIndex"`
	// A map of player UUIDs to Player structs, representing the players currently at the table
	Players map[string]*Player `json:"players"`
	// The current player who is flipping the card, nil if no player is flipping a card
	CurrentPlayer *Player `json:"currentPlayer"`
	// Indicates whether the game has started or not
	IsStarted bool `json:"isStarted"`
}

type PlayersTable struct {
	Layer TableLayerInterface `json:"-"`
}

func (t *TableLayer) GetUuid() string {
	return t.Uuid
}

// Returns the current card based on the DeckIndex. If the DeckIndex is not
// initialized (i.e., -1), it returns an error.
func (t *TableLayer) GetCurrentCard() (*Card, error) {
	if t.DeckIndex > 0 {
		return t.Deck[t.DeckIndex], nil
	} else {
		return nil, errors.New("Deck index is not initialized")
	}

}

// Flips a card for the given player. It increments the DeckIndex and checks if
// it has reached the end of the deck. If it has, it resets the DeckIndex and
// returns "next_round". Otherwise, it returns "continue".
func (t *TableLayer) FlipCard(player *Player) (string, *Card, error) {
	t.DeckIndex++

	if t.DeckIndex >= len(t.Deck) {
		t.DeckIndex = 0

		card, err := t.GetCurrentCard()
		if err != nil {
			return "", nil, err
		}

		card.Owner = player
		player.SetPlayerCards(t)

		return "next_round", nil, nil
	}
	card, err := t.GetCurrentCard()
	if err != nil {
		return "", nil, err
	}
	return "continue", card, nil
}

func (t *TableLayer) StartGame() {
	t.IsStarted = true
	t.DeckIndex = -1
}

func (t *TableLayer) EndGame() {
	t.IsStarted = false
}

func (t *TableLayer) FreezePlayer(player *Player) {
}

func (t *TableLayer) ResetPlayers() {
	for _, player := range t.Players {
		player.NumberOfCards = 0
		player.IsFreezed = false
		player.HasSecondChanceCard = false
		player.HasSevenCards = false
		player.Cards = []*Card{}
	}
}

// Returns the new deck for the current play
func (t *TableLayer) GetDeck() []*Card {
	numberCards := GetNumberCards()
	specialCards := GetSpecialCards()

	deck := append(t.Deck, numberCards...)
	deck = append(deck, specialCards...)

	return deck
}

// Returns the number of cards in the current deck
func (t *TableLayer) NumberOfCards() int {
	return len(t.Deck)
}

// Adds a player to the table with the given websocket connection, username and initiator status. It creates a new player using the CreatePlayer function, sets the player's websocket connection, and adds the player to the Players map of the table layer. Finally, it returns the newly created player.
func (t *TableLayer) AddPlayer(connection *websocket.Conn, username string, isInitiator bool) *Player {
	player := CreatePlayer(connection, username, t.Uuid, isInitiator)
	t.Players[player.Uuid] = player
	return player
}

func (t *TableLayer) HasPlayer(connection *websocket.Conn) bool {
	for _, player := range t.Players {
		if player.Conn == connection {
			return true
		}
	}
	return false
}

func (t *TableLayer) GetPlayer(connection *websocket.Conn) *Player {
	for _, player := range t.Players {
		if player.Conn == connection {
			return player
		}
	}
	return nil
}

func (t *TableLayer) NextPlayer(broadcaster *broadcasting.BroadcasterRegistry) {

}

// Returns a new table layer with a unique UUID, an empty deck,
// an uninitialized DeckIndex, an empty Players map, and IsStarted set to false.
func CreatePlayersTable() *PlayersTable {
	return &PlayersTable{
		Layer: &TableLayer{
			Uuid:      uuid.NewString(),
			Players:   make(map[string]*Player),
			Deck:      []*Card{},
			DeckIndex: -1,
		},
	}
}
