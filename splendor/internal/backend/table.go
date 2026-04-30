package backend

import (
	"fmt"
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
	Uuid          string           `json:"uuid"`
	CurrentRound  int              `json:"currentRound"`
	CurrentPlayer *WebsocketClient `json:"currentPlayer"`
	IsStarted     bool             `json:"isStarted"`
	StartedAt     time.Time        `json:"startedAt"`
}

type PlayingTableRules struct {
	MaxPlayers     int `json:"-"`
	MaxCardsPerRow int `json:"-"`
}

// PlayingTable represents the state of the game table,
// including the decks and the cards currently on the table.
type PlayingTable struct {
	PlayingTableDetails
	PlayingTableRules
	PlayingTableTokens

	DeckLevelOne   []CardInterface `json:"deckLevelOne"`
	DeckLevelTwo   []CardInterface `json:"deckLevelTwo"`
	DeckLevelThree []CardInterface `json:"deckLevelThree"`

	CardsLevelThree []CardInterface `json:"cardsLevelThree"`
	CardsLevelTwo   []CardInterface `json:"cardsLevelTwo"`
	CardsLevelOne   []CardInterface `json:"cardsLevelOne"`

	IsNormalGame bool `json:"isNormalGame"`

	// Clients connected to this specific table,
	// mapped by player UUIDs
	Clients map[string]WebsocketClientInterface[WebsocketMessage] `json:"-"`

	// Channel used to broadcast messages to all clients at the table
	broadcast chan WebsocketMessage `json:"-"`
	// Channel used to broadcast messages to all clients at the table
	// when a player registers on the table
	register chan WebsocketClientInterface[WebsocketMessage] `json:"-"`
	// Channel used to broadcast messages to all clients at the table
	// when a player unregisters from the table
	unregister chan WebsocketClientInterface[WebsocketMessage] `json:"-"`
	mu         sync.RWMutex                                    `json:"-"`
}

// Add a player to the table
func (t *PlayingTable) AddPlayer(player WebsocketClientInterface[WebsocketMessage]) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if len(t.Clients) >= t.MaxPlayers {
		return fmt.Errorf("Table is full. Max players: %d", t.MaxPlayers)
	}

	playerUuid := player.GetPlayer().Uuid
	if _, exists := t.Clients[playerUuid]; exists {
		return fmt.Errorf("Player with UUID %s is already at the table", playerUuid)
	}

	t.Clients[playerUuid] = player
	return nil
}

// Start the game and initialize the table.
func (t *PlayingTable) StartGame() error {
	// Depending on the number of players,
	// we need to adjust the number of tokens
	// available on the table.
	switch t.NumberOfPlayers() {
	// case 1:
	// 	return fmt.Errorf("At least 2 players are required to start the game")
	case 2:
		t.Emerald = 4
		t.Diamond = 4
		t.Sapphire = 4
		t.Onyx = 4
		t.Ruby = 4
	case 3:
		t.Emerald = 5
		t.Diamond = 5
		t.Sapphire = 5
		t.Onyx = 5
		t.Ruby = 5
	case 4:
		t.Emerald = 7
		t.Diamond = 7
		t.Sapphire = 7
		t.Onyx = 7
		t.Ruby = 7
	}

	// Populate the decks and the cards on the table.
	t.NewDeck()
	t.CardsLevelOne = t.DeckLevelOne[:t.MaxCardsPerRow]
	t.CardsLevelTwo = t.DeckLevelTwo[:t.MaxCardsPerRow]
	t.CardsLevelThree = t.DeckLevelThree[:t.MaxCardsPerRow]

	t.CurrentRound = 1
	t.IsStarted = true
	t.StartedAt = time.Now()

	t.BroadcastMessage(WebsocketMessage{
		Action:       "game_started",
		PlayingTable: t.GetTableDetails(),
	})

	// if t.IsStarted && t.NumberOfPlayers() > 2 {
	// }
	return nil
}

// End the game and reset the table.
func (t *PlayingTable) EndGame() error {
	return nil
}

// Shuffle the cards of each level and reveals four cards
// on each level to be available for purchase.
func (t *PlayingTable) NewDeck() {
	if t.IsNormalGame {
		t.DeckLevelOne = NormalCardsLevelOne()
		t.DeckLevelTwo = NormalCardsLevelTwo()
		t.DeckLevelThree = NormalCardsLevelThree()
	} else {
		t.DeckLevelOne = MarvelCardsLevelOne()
		t.DeckLevelTwo = MarvelCardsLevelTwo()
		t.DeckLevelThree = MarvelCardsLevelThree()
	}
}

// Return the number of players currently at the table.
func (t *PlayingTable) NumberOfPlayers() int {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return len(t.Clients)
}

// Starts the broadcaster for the table, which listens for messages
// to broadcast to all clients at the table.
func (t *PlayingTable) StartBroadcaster() {
	go func() {
		for {
			select {
			case client := <-t.register:
				t.Clients[client.GetPlayer().Uuid] = client
			case client := <-t.unregister:
				delete(t.Clients, client.GetPlayer().Uuid)
			case message := <-t.broadcast:
				for _, client := range t.Clients {
					err := client.SendJsonMessage(message)
					if err != nil {
						t.unregister <- client
					}
				}
			}
		}
	}()
}

func (t *PlayingTable) BroadcastMessage(message WebsocketMessage) {
	t.broadcast <- message
}

// GetTableDetails returns a map containing the details of the table,
// including the current round, current player, and available resources.
func (t *PlayingTable) GetTableDetails() map[string]any {
	details := map[string]any{
		"decklevelone":    t.DeckLevelOne,
		"deckleveltwo":    t.DeckLevelTwo,
		"decklevelthree":  t.DeckLevelThree,
		"cardslevelthree": t.CardsLevelThree,
		"cardsleveltwo":   t.CardsLevelTwo,
		"cardslevelone":   t.CardsLevelOne,
		"isNormalGame":    t.IsNormalGame,
		"clients":         []map[string]string{},
		"details": map[string]any{
			"uuid":          t.Uuid,
			"currentRound":  t.CurrentRound,
			"currentPlayer": t.CurrentPlayer,
			"isStarted":     t.IsStarted,
			"startedAt":     t.StartedAt,
		},
		"tokens": map[string]any{
			"normalTokens": map[string]int{
				"emerald":  t.Emerald,
				"diamond":  t.Diamond,
				"sapphire": t.Sapphire,
				"onyx":     t.Onyx,
				"ruby":     t.Ruby,
			},
			"marvelTokens": map[string]int{
				"mind":    t.Mind,
				"space":   t.Space,
				"soul":    t.Soul,
				"power":   t.Power,
				"reality": t.Reality,
				"time":    t.Time,
				"shield":  t.Shield,
			},
		},
	}

	for key, client := range t.Clients {
		details["clients"] = append(details["clients"].([]map[string]string), map[string]string{
			"username": client.GetPlayer().Username,
			"uuid":     key,
		})
	}

	return details
}

func NewPlayingTable(isNormalGame bool) *PlayingTable {
	return &PlayingTable{
		PlayingTableDetails: PlayingTableDetails{
			Uuid:          uuid.NewString(),
			CurrentRound:  0,
			CurrentPlayer: nil,
			IsStarted:     false,
			StartedAt:     time.Now(),
		},
		PlayingTableRules: PlayingTableRules{
			MaxPlayers:     4,
			MaxCardsPerRow: 4,
		},
		PlayingTableTokens: PlayingTableTokens{
			CardResources: CardResources{
				Emerald:  7,
				Diamond:  7,
				Sapphire: 7,
				Onyx:     7,
				Ruby:     7,
			},
			MarvelCardResources: MarvelCardResources{
				Mind:    7,
				Space:   7,
				Soul:    7,
				Power:   7,
				Reality: 7,
				Time:    7,
				Shield:  7,
			},
		},
		DeckLevelOne:    []CardInterface{},
		DeckLevelTwo:    []CardInterface{},
		DeckLevelThree:  []CardInterface{},
		CardsLevelThree: []CardInterface{},
		CardsLevelTwo:   []CardInterface{},
		CardsLevelOne:   []CardInterface{},
		IsNormalGame:    isNormalGame,
		Clients:         make(map[string]WebsocketClientInterface[WebsocketMessage]),
		broadcast:       make(chan WebsocketMessage),
		register:        make(chan WebsocketClientInterface[WebsocketMessage]),
		unregister:      make(chan WebsocketClientInterface[WebsocketMessage]),
	}
}

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
