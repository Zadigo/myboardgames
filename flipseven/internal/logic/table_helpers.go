package logic

import (
	"context"
	"maps"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

// Pops x number of cards from the deck
func (table *PlayersTable) FlipCard(playerId string, k int) (*Card, string) {
	if k <= 0 {
		return nil, ""
	}

	if k > 3 {
		k = 3
	}

	if table.GetNumberOfCards() == 0 {
		return nil, ""
	}

	player := table.GetPlayer(playerId)
	if player.Details.HasSevenCards {
		return nil, ""
	}

	if player.Details.IsFreezed {
		return nil, ""
	}

	card := &table.CurrentDeck[k]
	card.Owner = playerId

	if len(player.Details.Cards) == 7 {
		player.Details.HasSevenCards = true
		return nil, ""
	}

	player.Details.Cards = append(player.Details.Cards, *card)

	// continue or next_round
	var action string = "continue"

	// Once the player has flipped a card, we need to
	// check again if they reached seven cards and if so,
	// we need to freeze them and all the other players
	if player.Details.HasSevenCards {
		for _, player := range table.Clients {
			player.Details.IsFreezed = true
		}
		action = "next_round"
	}

	return card, action
}

// Prevents a player from flipping cards for the rest of the round
func (table *PlayersTable) FreezePlayer(playerId string) {
	player := table.GetPlayer(playerId)
	if player != nil {
		player.Details.IsFreezed = true
	}
}

func (table *PlayersTable) NextRound() {
	table.CurrentRound++
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
	return !table.GameStarted && table.NumberOfPlayers <= table.maxNumberOfPlayers
}

func (table *PlayersTable) IsStarted() bool {
	return table.GameStarted
}

// Returns a copy of the PlayersTable struct to prevent 
// external modification of the original struct
func (table *PlayersTable) GetTable() *PlayersTable {
	return table
}

// Create a new table layer interface
func CreateNewTableLayer() *TableLayer {
	return &TableLayer{
		Layer: &PlayersTable{
			TableId:            uuid.NewString(),
			Clients:            make(map[string]*PlayerLayer),
			maxNumberOfPlayers: 12,
		},
	}
}
