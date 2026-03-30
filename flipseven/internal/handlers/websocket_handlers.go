package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/Zadigo/flipseven/internal"
	"github.com/Zadigo/flipseven/internal/backend"
	"github.com/Zadigo/flipseven/internal/cards"
	"github.com/Zadigo/flipseven/internal/logic"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

var mutex = sync.RWMutex{}

// In-memory storage for game tables and connected clients.
// This is used for managing the state of the game and
// the connected clients.
var Tables = make(map[string]*internal.PlayersTable)

// In-memory storage for connected clients. This is used for
// broadcasting messages to all connected clients.
var clients = make(map[*websocket.Conn]bool)

const maxNumberOfPlayers int = 12

// Handler for the live game websocket connection. This is used for real-time
// communication between the server and the clients during the game.
func LiveGameHandler(response http.ResponseWriter, request *http.Request, ctx context.Context, redisConn *redis.Client, broadcaster *backend.BroadcasterRegistry) {
	connection, err := RequestUpgrader.Upgrade(response, request, nil)

	if err != nil {
		log.Println("❌ Failed to upgrade request", err.Error())
		return
	}

	// Reference any connections regardless of
	// whether they are associated to a table
	// or not
	mutex.Lock()
	clients[connection] = true
	log.Printf("⚡️ New connection from client")
	mutex.Unlock()

	defer func() {
		mutex.Lock()
		delete(clients, connection)
		mutex.Unlock()

		// Send proper closing handshake before closing
		connection.WriteMessage(
			websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
		)

		log.Printf("⚡️ Connection closed")
		connection.Close()
	}()

	WriteWsMessage(connection, WebsocketMessage{
		Action:  "initial_connection",
		Message: "Connection successful",
	})

	var currentRound int = 1

	for {
		var message WebsocketMessage
		err := connection.ReadJSON(&message)
		if err != nil {
			WriteWsMessage(connection, WebsocketMessage{
				Action:  "error",
				Message: fmt.Sprintf("❌ Failed to send JSON message: %v", err.Error()),
			})
			return
		}

		tableId := request.URL.Query().Get("table")

		// select {
		// case <-ctx.Done():
		// 	log.Println("⚠️ Context cancelled, closing handler")
		// 	return
		// default:
		// }

		switch message.Action {
		case "waiting_lobby":
			if err := CheckTableId(tableId); err != nil {
				WriteWsMessage(connection, WebsocketMessage{
					Action:  "error",
					Message: err.Error(),
				})
				return
			}

			connectedPlayer := internal.ConnectedPlayer{
				Conn: connection,
				Details: internal.Player{
					Username:      message.Username,
					Uuid:          uuid.NewString(),
					TableUuid:     tableId,
					NumberOfCards: 0,
				},
			}

			// If we have an initiator (aka a table on Redis), we have
			// a table... we need to recreate it locally on the Tables
			result := redisConn.HExists(ctx, tableId, "openForJoin").Val()
			if !result {
				WriteWsMessage(connection, WebsocketMessage{
					Action:  "error",
					Message: "Table not found or not open for join",
				})
				return
			}

			// Connect any new players to the table
			mutex.Lock()
			table := Tables[tableId]

			if table == nil {
				table = &internal.PlayersTable{
					NumberOfPlayers: 0,
					GameStarted:     false,
					Clients:         []*internal.ConnectedPlayer{},
				}
				Tables[tableId] = table
			}

			if table.GameStarted {
				WriteWsMessage(connection, WebsocketMessage{
					Action:  "error",
					Message: "Game has already started",
				})
				mutex.Unlock()
				return
			}

			table.NumberOfPlayers += 1
			table.Clients = append(table.Clients, &connectedPlayer)
			mutex.Unlock()

			// Register with the broadcaster so Redis messages reach this client
			broadcaster := broadcaster.GetOrCreate(tableId)
			broadcaster.AddClient(connection)

			if table.NumberOfPlayers == maxNumberOfPlayers {
				WriteWsMessage(connection, WebsocketMessage{
					Action: "start_game",
				})

				return
			}

			WriteWsMessage(connection, WebsocketMessage{
				Action:       "update_waiting_lobby",
				TableDetails: *table,
			})

		case "get_deck":
			if err := CheckTableId(tableId); err != nil {
				WriteWsMessage(connection, WebsocketMessage{
					Action:  "error",
					Message: err.Error(),
				})
				return
			}

			// 1. Return the deck of cards
			baseDeck := cards.GetDeck()

			// table := Tables[message.TableId]
			// table.CurrentDeck = baseDeck
			table, err := GetTable(tableId)
			if err != nil {
				WriteWsMessage(connection, WebsocketMessage{
					Action:  "error",
					Message: err.Error(),
				})
				return
			}

			table.CurrentDeck = baseDeck

			WriteWsMessage(connection, WebsocketMessage{
				Action: "deck_created",
				Deck:   baseDeck,
			})

		case "flip_card":
			if err := CheckTableId(tableId); err != nil {
				WriteWsMessage(connection, WebsocketMessage{
					Action:  "error",
					Message: err.Error(),
				})
				return
			}

			table, err := GetTable(tableId)
			if err != nil {
				WriteWsMessage(connection, WebsocketMessage{
					Action:  "error",
					Message: err.Error(),
				})
				return
			}

			for _, player := range table.Clients {
				if player.Details.IsFreezed {
					continue
				}

				if player.Conn == connection {
					card := cards.FlipCard(table.CurrentDeck, 1)
					cards.AttributeCardToPlayer(card, player)
				}
			}

			goToNextRound := false

			// If a player has flipped seven cards, move
			// to the next round
			for _, player := range table.Clients {
				if player.Details.HasSevenCards {
					goToNextRound = true
					break
				}
			}

			// If all the players are frozen, calculate the
			// scores and move to a next round
			var frozenPlayers int = 0
			for _, player := range table.Clients {
				if player.Details.IsFreezed {
					frozenPlayers += 1
				}
			}

			if frozenPlayers == maxNumberOfPlayers {
				goToNextRound = true
			}

			if goToNextRound {
				for _, player := range table.Clients {
					logic.CalcualtePlayerScores(player)
				}

				WriteWsMessage(connection, WebsocketMessage{
					Action: "new_round",
				})

				currentRound += 1
			}

		default:
		}
	}
}
