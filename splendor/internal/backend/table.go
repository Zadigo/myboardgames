package backend

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

// PlayingTableTokens represents the tokens available on the table,
// including the number of each type of token and the gold jokers.
type PlayingTableTokens struct {
	CardResources
	MarvelCardResources
}

type PlayingTableDetails struct {
	Uuid          string `json:"uuid"`
	CurrentRound  int
	CurrentPlayer *WebsocketClient `json:"currentPlayer"`
	IsStarted     bool             `json:"isStarted"`
	StartedAt     time.Time        `json:"startedAt"`
}

type PlayingTableRules struct {
	MaxPlayers     int
	MaxCardsPerRow int
}

// PlayingTable represents the state of the game table,
// including the decks and the cards currently on the table.
type PlayingTable struct {
	PlayingTableDetails
	PlayingTableRules

	DeckLevelOne   []CardInterface `json:"deckLevelOne"`
	DeckLevelTwo   []CardInterface `json:"deckLevelTwo"`
	DeckLevelThree []CardInterface `json:"deckLevelThree"`

	CardsLevelThree []CardInterface `json:"cardsLevelThree"`
	CardsLevelTwo   []CardInterface `json:"cardsLevelTwo"`
	CardsLevelOne   []CardInterface `json:"cardsLevelOne"`

	IsNormalGame bool `json:"isNormalGame"`

	// Clients connected to this specific table,
	// mapped by player UUIDs
	Clients []map[string]*WebsocketClient `json:"clients"`

	broadcast  chan WebsocketMessage `json:"-"`
	register   chan *WebsocketClient `json:"-"`
	unregister chan *WebsocketClient `json:"-"`
	mu         sync.RWMutex          `json:"-"`
}

// Add a player to the table
func (t *PlayingTable) AddPlayer(player *Player) error {
	return nil
}

// Start the game and initialize the table.
func (t *PlayingTable) StartGame() {}

// End the game and reset the table.
func (t *PlayingTable) EndGame() {}

func (t *PlayingTable) NewDeck() {}

// Return the number of players currently at the table.
func (t *PlayingTable) NumberOfPlayers() int {
	return 0
}

func NewPlayingTable(isNormalGame bool) *PlayingTable {
	return &PlayingTable{
		PlayingTableDetails: PlayingTableDetails{
			Uuid:          uuid.NewString(),
			CurrentRound:  0,
			CurrentPlayer: nil,
			IsStarted:     false,
			StartedAt:     time.Time{},
		},
		PlayingTableRules: PlayingTableRules{
			MaxPlayers:     4,
			MaxCardsPerRow: 4,
		},
		CardsLevelThree: []CardInterface{},
		CardsLevelTwo:   []CardInterface{},
		CardsLevelOne:   []CardInterface{},
		IsNormalGame:    isNormalGame,
		broadcast:       make(chan WebsocketMessage),
		register:        make(chan *WebsocketClient),
		unregister:      make(chan *WebsocketClient),
	}
}

// // Start the game and initialize the table.
// func (t *PlayingTable) StartGame() {
// 	if !t.GameStarted && t.NumberOfPlayers() > 2 {
// 		t.GameStarted = true

// 		if t.NumberOfPlayers() == 2 {
// 			t.Emerald = 4
// 			t.Diamond = 4
// 			t.Sapphire = 4
// 			t.Onyx = 4
// 			t.Ruby = 4
// 		}

// 		if t.NumberOfPlayers() == 3 {
// 			t.Emerald = 5
// 			t.Diamond = 5
// 			t.Sapphire = 5
// 			t.Onyx = 5
// 			t.Ruby = 5
// 		}

// 		if t.NumberOfPlayers() == 4 {
// 			t.Emerald = 7
// 			t.Diamond = 7
// 			t.Sapphire = 7
// 			t.Onyx = 7
// 			t.Ruby = 7
// 		}
// 	}
// }

// // End the game and reset the table.
// func (t *PlayingTable) EndGame() {
// 	if t.GameStarted {
// 		t.GameStarted = false
// 		t.Players = make(map[string]*Player)

// 		t.CardsLevelThree = []*Card{}
// 		t.CardsLevelTwo = []*Card{}
// 		t.CardsLevelOne = []*Card{}

// 		t.Emerald = 7
// 		t.Diamond = 7
// 		t.Sapphire = 7
// 		t.Onyx = 7
// 		t.Ruby = 7
// 		t.GoldJoker = 5
// 	}
// }

// func (t *PlayingTable) AddCard(card *Card, level int) {

// }

// func (t *PlayingTable) NewDeck() {
// 	cardsLevelOne := CardsLevelOne()
// 	cardsLevelTwo := CardsLevelTwo()
// 	cardsLevelThree := CardsLevelThree()

// 	t.DeckLevelOne = append(t.DeckLevelOne, cardsLevelOne...)
// 	t.DeckLevelTwo = append(t.DeckLevelTwo, cardsLevelTwo...)
// 	t.DeckLevelThree = append(t.DeckLevelThree, cardsLevelThree...)
// }

// // Return the number of players currently at the table.
// func (t *PlayingTable) NumberOfPlayers() int {
// 	return len(t.Players)
// }

// // Create the initial influnce queue with 5 cards of
// // each color for each player
// func CreateTableLayer() *TableLayer {
// 	return &TableLayer{
// 		CurrentRound:   0,
// 		CurrentPlayer:  nil,
// 		MaxPlayers:     4,
// 		MaxCardsPerRow: 4,
// 		Layer: &PlayingTable{
// 			CardsLevelThree: []*Card{},
// 			CardsLevelTwo:   []*Card{},
// 			CardsLevelOne:   []*Card{},
// 		},
// 	}
// }
