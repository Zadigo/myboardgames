package models

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type GameOptions struct{}

// GameInterface represents a game logic that a given
// table should host. The interface implements basic
// game logic that most games need.
type GameLogicInterface interface {
	// Entrypoint option used by the Table to execute
	// game logic actions for the given game
	Execute(options GameOptions)
}

type TableInterface interface {
	GetID() string
	// Run blocks the table until the context is
	// terminated and waits for instructions
	Run(ctx context.Context, optoins *TableOptions)
	// Join adds a player to the table
	Join(client WebsocketClientInterface)
	// Leave removes a player from the table
	Leave(client WebsocketClientInterface)
	// DistributeCards is instruction that
	// can be used to distribute cards to
	// the players around the table
	DistributeCards()
	// NextRound moves to the whole game
	// to the next round
	NextRound()
	// NextTurn from a player's turn to
	// another player
	NextTurn()
	// ShuffleCards shuffles the deck
	ShuffleCards()
	// Reset clears the table state
	Reset()
	// SetGame sets the game logic for the table
	Setup(GameLogicInterface)
}

// Common table events that occur in most board games
const (
	TableEventIdle    string = "table_event_idle"    // The table is waiting for players to join
	TableEventWaiting string = "table_event_waiting" // The table is waiting for player(s) to take a decision
	TableEventCreated string = "table_event_created"
	TableEventUpdated string = "table_event_updated"

	TableEventAddPlayer    string = "table_event_add_player"
	TableEventRemovePlayer string = "table_event_remove_player"
	TableEventBanPlayer    string = "table_event_ban_player"

	TableEventCreateDeck  string = "table_event_create_deck"
	TableEventShuffleDeck string = "table_event_shuffle_cards"

	TableEventDistributeCards         string = "table_event_distribute_cards"
	TableEventExchangeCardsWithPlayer string = "table_event_exchange_cards_with_player"
	TableEventExchangeCardsWithDeck   string = "table_event_exchange_cards_with_deck"
	TableEventStealCard               string = "table_event_steal_card"
	TableEventDrawCard                string = "table_event_draw_card"

	TableEventNextTurn    string = "table_event_next_turn"
	TableEventReverseTurn string = "table_event_reverse_turn"
)

type TableEventDetails[T any] struct {
	EventName string `json:"event_name"`
	// The single or group of players that
	// triggered the event
	From []WebsocketClientInterface `json:"from"`
	// The single or group of players affected
	// by the event (if applicable)
	To []WebsocketClientInterface `json:"to"`
	// The card that is responsible for the event
	Card T `json:"card"`
}

// TableSnapshot holds the current state of the
// table which can then be used for restoring state
// and emergency recovery
type TableSnapshot[D any] struct {
	Deck []D `json:"deck"`
}

// TableEvents represents events
// happening on the table. These events
// are also registered on Redis to recover
// for a example a previous move
type TableEvent struct {
	ID        string                   `json:"id"`
	GameID    string                   `json:"game_id"`
	Details   TableEventDetails[any]   `json:"details"`
	Snapshot  TableSnapshot[any]       `json:"snapshot"`
	CreatedAt time.Time                `json:"created_at"`
	Client    WebsocketClientInterface `json:"client"`
}

// Table represents a game table with players.
type Table struct {
	TableInterface
	ID            string                     `json:"id"`
	TimeLimit     int                        `json:"time_limit"`
	Duration      time.Duration              `json:"duration"`
	CurrentRound  int                        `json:"current_index"`
	CurrentPlayer int                        `json:"current_player"`
	UpdatedAt     time.Time                  `json:"updated_at"`
	CreatedAt     time.Time                  `json:"created_at"`
	inbox         chan TableEvent            `json:"-"`
	game          GameLogicInterface         `json:"-"`
	state         bool                       `json:"-"`
	players       []WebsocketClientInterface `json:"-"`
}

func (t *Table) GetID() string {
	return t.ID
}

func (t *Table) Setup(g GameLogicInterface) {
	t.game = g
	t.inbox = make(chan TableEvent, 32)
	t.players = make([]WebsocketClientInterface, 20)
}

func (t *Table) GetInbox() <-chan TableEvent {
	return t.inbox
}

type TableOptions struct {
	Redis      *redis.Client
	MainServer ServerAppInterface
	Game       GameLogicInterface
}
