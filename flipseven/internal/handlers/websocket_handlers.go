package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/Zadigo/flipseven/internal/backend"
	"github.com/Zadigo/flipseven/internal/logic"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

var mutex = sync.RWMutex{}

// In-memory storage for game tables and connected clients.
// This is used for managing the state of the game and
// the connected clients.
var Tables = make(map[string]*logic.TableLayer)

// In-memory storage for connected clients. This is used for
// broadcasting messages to all connected clients.
var clients = make(map[*websocket.Conn]bool)

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

			tableLayer := Tables[tableId]

			if tableLayer == nil {
				WriteWsMessage(connection, WebsocketMessage{
					Action:  "error",
					Message: "Table does not exist",
				})
				return
			}

			tableLayer = logic.CreateNewTableLayer()
			tableLayer.Layer.AddPlayer(message.Username, connection)

			// If we have an initiator (aka a table on Redis), we have
			// a table... we need to recreate it locally on the Tables
			// result := redisConn.HExists(ctx, tableId, "openForJoin").Val()
			result := tableLayer.Layer.TableExists(redisConn, ctx)
			if !result {
				WriteWsMessage(connection, WebsocketMessage{
					Action:  "error",
					Message: "Table not found or not open for join",
				})
				return
			}

			// Connect any new players to the table
			mutex.Lock()
			if !tableLayer.Layer.CanAcceptNewPlayers() {
				WriteWsMessage(connection, WebsocketMessage{
					Action:  "error",
					Message: "Game has already started or table is full",
				})
				mutex.Unlock()
				return
			}

			tableLayer.Layer.AddPlayer("Test Player", connection)
			Tables[tableId] = tableLayer

			if tableLayer.Layer.IsStarted() {
				WriteWsMessage(connection, WebsocketMessage{
					Action: "start_game",
				})
				mutex.Unlock()
				return
			}
			mutex.Unlock()

			// Register with the broadcaster so Redis messages reach this client
			// broadcaster := broadcaster.GetOrCreate(tableId)
			// broadcaster.AddClient(connection)

			WriteWsMessage(connection, WebsocketMessage{
				Action:       "update_waiting_lobby",
				TableDetails: tableLayer.Layer.GetTable(),
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
			// baseDeck := logic.GetDeck()

			// table := Tables[message.TableId]
			// table.CurrentDeck = baseDeck
			tableLayer, err := GetTableLayer(tableId)
			if err != nil {
				WriteWsMessage(connection, WebsocketMessage{
					Action:  "error",
					Message: err.Error(),
				})
				return
			}

			deck := tableLayer.Layer.GetDeck()

			WriteWsMessage(connection, WebsocketMessage{
				Action: "deck_created",
				Deck:   deck,
			})

		case "flip_card":
			if err := CheckTableId(tableId); err != nil {
				WriteWsMessage(connection, WebsocketMessage{
					Action:  "error",
					Message: err.Error(),
				})
				return
			}

			tableLayer, err := GetTableLayer(tableId)
			if err != nil {
				WriteWsMessage(connection, WebsocketMessage{
					Action:  "error",
					Message: err.Error(),
				})
				return
			}

			tableLayer.Layer.FlipCard(message.PlayerId, 1)

			goToNextRound := false

			player := tableLayer.Layer.GetPlayer(message.PlayerId)
			if player.Details.HasSevenCards {
				for _, player := range tableLayer.Layer.GetTable().Clients {
					player.Details.IsFreezed = true
					WriteWsMessage(player.Conn, WebsocketMessage{
						Action:  "player_frozen",
						Message: "Players are frozen and cannot flip more cards because one of the players has flipped seven cards",
					})
				}
			}

			// If all the players are frozen, calculate the
			// scores and move to a next round
			// var frozenPlayers int = 0
			// for _, player := range tableLayer.Layer.GetTable().Clients {
			// 	if player.Details.IsFreezed {
			// 		frozenPlayers += 1
			// 	}
			// }

			// if frozenPlayers == 12 {
			// 	goToNextRound = true
			// }

			if goToNextRound {
				tableLayer.Layer.CalculateAllPoints()

				for _, player := range tableLayer.Layer.GetTable().Clients {
					player.ResetPlayerState()
				}

				WriteWsMessage(connection, WebsocketMessage{
					Action:       "new_round",
					TableDetails: tableLayer.Layer.GetTable(),
				})

				currentRound += 1
			}

		default:
			WriteWsMessage(connection, WebsocketMessage{
				Action:  "error",
				Message: "Unknown action",
			})
		}
	}
}
