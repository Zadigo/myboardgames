package logic

import (
	"context"
	"maps"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

const maxNumberOfPlayers int = 12

func (table *PlayersTable) StartGame(connection *websocket.Conn) {
	// Do something
}

func (table *PlayersTable) EndGame(connection *websocket.Conn) {
	// Do something
}

// Pops x number of cards from the deck
func (table *PlayersTable) FlipCard(playerId string, k int) *Card {
	if k <= 0 {
		return nil
	}

	if k > 3 {
		k = 3
	}

	if table.GetNumberOfCards() == 0 {
		return nil
	}

	card := &table.CurrentDeck[k]
	player := table.GetPlayer(playerId)

	if len(player.Details.Cards) == 7 {
		player.Details.HasSevenCards = true
		return nil
	}

	player.Details.Cards = append(player.Details.Cards, *card)
	return card
}

func (table *PlayersTable) FreezePlayer(player string) {
	// Do something
}

func (table *PlayersTable) NextRound() {
	// Do something
}

func (table *PlayersTable) GetTableId() string {
	return table.TableId
}

func (table *PlayersTable) GetPlayer(playerId string) *PlayerLayer {
	// Do something
	return table.Clients[playerId]
}

func (table *PlayersTable) HasPlayer(PlayerId string) bool {
	_, exists := table.Clients[PlayerId]
	return exists
}

// Checks if the table exists in Redis by checking the existence of a specific
// key (e.g., "openForJoin") associated with the table ID.
func (table *PlayersTable) TableExists(redisConn *redis.Client, ctx context.Context) bool {
	result := redisConn.HExists(context.Background(), table.TableId, "openForJoin").Val()
	return result
}

// Calculates the points for all the players on the table by calling
// the TotalPoints method for each player
func (table *PlayersTable) CalculateAllPoints() int {
	totalGains := 0

	for _, player := range table.Clients {
		totalGains += player.Details.TotalPoints()
	}

	return totalGains
}

// Returns the number of cards in the current deck
func (table *PlayersTable) GetNumberOfCards() int {
	return len(table.CurrentDeck)
}

// Returns the number of players on the table
func (table *PlayersTable) GetNumberOfPlayers() int {
	keys := maps.Keys(table.Clients)

	count := 0
	for range keys {
		count++
	}

	return count
}

func (table *PlayersTable) AddPlayer(username string, connection *websocket.Conn) {
	if username == "" {
		// Do something
		return
	}

	table.Clients[username] = &PlayerLayer{
		Conn: connection,
		Details: Player{
			Username:    username,
			Uuid:        uuid.NewString(),
			IsInitiator: false,
		},
	}
}

func (table *PlayersTable) CanAcceptNewPlayers() bool {
	return !table.GameStarted && table.NumberOfPlayers < maxNumberOfPlayers
}

func (table *PlayersTable) IsStarted() bool {
	return table.GameStarted
}

func (table *PlayersTable) GetTable() PlayersTable {
	return *table
}

// Create a new table layer interface
func CreateNewTableLayer() *TableLayer {
	return &TableLayer{
		Layer: &PlayersTable{
			TableId: uuid.NewString(),
			Clients: make(map[string]*PlayerLayer),
		},
	}
}
