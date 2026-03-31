package handlers

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

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
func LiveGameHandler(response http.ResponseWriter, request *http.Request, ctx context.Context, redisConn *redis.Client, broadcaster *backend.BroadcasterRegistry, baseRegistry *BaseRegistry) {
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

	for {
		message, err := ReadWsMessage(connection)
		if err != nil {
			return
		}

		tableId := request.URL.Query().Get("table")

		switch message.Action {
		case "initiate_table":
			// TODO:
			tableLayer := baseRegistry.GetOrCreate(tableId, func() *logic.TableLayer {
				layer := logic.CreateNewTableLayer()
				// TODO: Username
				tableId = layer.Layer.GetTable().TableId

				result := redisConn.HSet(context.Background(), tableId, []string{"initiator", "Alice Test", "createdAt", time.Now().Format(time.RFC3339), "openForJoin", "true"})
				if result.Err() != nil {
					log.Printf("❌ Failed to set table details in Redis for table ID: %s, error: %s", tableId, result.Err().Error())
					return nil
				} else {
					log.Printf("✅ Created new table with ID: %s", tableId)
				}

				return layer
			})

			if tableLayer.Layer == nil {
				log.Printf("❌ Could not initialize PlayersTable for table ID: %s", tableId)
				return
			}

			WriteWsMessage(connection, WebsocketMessage{
				Action:  "table_initiated",
				TableId: tableId,
			})

		case "reconnect":
			tableLayer, state := baseRegistry.Get(message.TableId)
			if !state {
				log.Printf("❌ Table with ID %s not found in registry for reconnection", message.TableId)
			}

			log.Printf("🔄 Reconnecting player %s to table %s", message.Username, message.TableId)

			WriteWsMessage(connection, WebsocketMessage{
				Action:       "reconnected",
				TableDetails: tableLayer.Layer.GetTable(),
			})

		case "waiting_lobby":
			tableLayer, state := baseRegistry.Get(message.TableId)
			if !state {
				log.Printf("❌ Table with ID %s not found in registry", message.TableId)
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
			tableLayer, state := baseRegistry.Get(message.TableId)
			if !state {
				WriteWsMessage(connection, WebsocketMessage{
					Action:  "error",
					Message: "Table not found",
				})
				return
			}

			deck := tableLayer.Layer.GetDeck()

			WriteWsMessage(connection, WebsocketMessage{
				Action: "deck_created",
				Deck:   deck,
			})

		case "flip_card":
			tableLayer, state := baseRegistry.Get(message.TableId)
			if !state {
				WriteWsMessage(connection, WebsocketMessage{
					Action:  "error",
					Message: "Table not found",
				})
				return
			}

			_, action := tableLayer.Layer.FlipCard(message.PlayerId, 1)

			if action == "next_round" {
				tableLayer.Layer.CalculateAllPoints()

				for _, player := range tableLayer.Layer.GetTable().Clients {
					player.ResetPlayerState()
				}

				tableLayer.Layer.NextRound()

				WriteWsMessage(connection, WebsocketMessage{
					Action:       "new_round",
					TableDetails: tableLayer.Layer.GetTable(),
				})
			}

		default:
			WriteWsMessage(connection, WebsocketMessage{
				Action:  "error",
				Message: "Unknown action",
			})
		}

		// select {
		// case <-ctx.Done():
		// 	log.Print("⚠️ Context cancelled, closing handler")
		// 	return
		// default:
		// }
	}
}
