package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

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
var Tables = make(map[string]*logic.PlayersTable)

// In-memory storage for connected clients. This is used for
// broadcasting messages to all connected clients.
var clients = make(map[*websocket.Conn]bool)

const maxNumberOfPlayers int = 12

var allowedOrigins = map[string]bool{
	"http://localhost:3000": true,
}

var RequestUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(request *http.Request) bool {
		origin := request.Header.Get("Origin")

		_, ok := allowedOrigins[origin]
		if !ok {
			return false
		}

		return allowedOrigins[origin]
	},
}

// CORS middleware to handle cross-origin requests
func Cors(next http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Access-Control-Allow-Origin", "*")
		response.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		response.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if request.Method == "OPTIONS" {
			response.WriteHeader(http.StatusOK)
			return
		}

		next(response, request)
	}
}

// Helper function to get a game table by its ID. This is used in various handlers
// to retrieve the game table associated with a specific table ID.
func getTable(tableId string) (*logic.PlayersTable, error) {
	if tableId == "" {
		return nil, fmt.Errorf("Table ID is required")
	}

	mutex.RLock()
	table, exists := Tables[tableId]
	mutex.RUnlock()

	if !exists {
		return nil, fmt.Errorf("Table with ID %s not found", tableId)
	}

	return table, nil
}

func checkTableId(tableId string) error {
	if tableId == "" {
		return fmt.Errorf("Table ID is required")
	}
	return nil
}

// Handler for creating a new game table. This is used when the user creates
// a new game and needs a unique table ID that needs to be shared with other players.
func CreateTableHandler(response http.ResponseWriter, request *http.Request, redisClient *redis.Client) {
	if request.Method != http.MethodPost {
		http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if request.Header.Get("Content-Type") != "application/json" {
		http.Error(response, "Unsupported media type", http.StatusUnsupportedMediaType)
		return
	}

	tableId := uuid.NewString()

	var message struct{ Username string }
	err := json.NewDecoder(request.Body).Decode(&message)

	if message.Username == "" {
		http.Error(response, "Username is required", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(response, fmt.Sprintf("Failed to parse data: %v", err.Error()), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	if message.Username == "" {
		http.Error(response, "Username is required", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	redisClient.HSet(ctx, tableId, []string{"initiator", message.Username, "createdAt", time.Now().Format(time.RFC3339), "openForJoin", "true"})

	err = json.NewEncoder(response).Encode(PostDataMessage{
		TableId: tableId,
	})

	if err != nil {
		http.Error(response, fmt.Sprintf("Failed to encode response: %v", err.Error()), http.StatusInternalServerError)
		return
	}

	log.Printf("✅ Created new table with ID: %s", tableId)
}

// Handler for the live game websocket connection. This is used for real-time
// communication between the server and the clients during the game.
func LiveGameHandler(response http.ResponseWriter, request *http.Request, ctx context.Context, redisConn *redis.Client) {
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
			if err := checkTableId(tableId); err != nil {
				WriteWsMessage(connection, WebsocketMessage{
					Action:  "error",
					Message: err.Error(),
				})
				return
			}

			connectedPlayer := logic.ConnectedPlayer{
				Conn: connection,
				Details: logic.Player{
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
				connection.WriteJSON(WebsocketMessage{
					Action:  "error",
					Message: "Table not found or not open for join",
				})
				return
			}

			// Connect any new players to the table
			mutex.Lock()
			table := Tables[tableId]

			if table == nil {
				table = &logic.PlayersTable{
					NumberOfPlayers: 0,
					GameStarted:     false,
					Clients:         []*logic.ConnectedPlayer{},
				}
				Tables[tableId] = table
			}

			if table.GameStarted {
				connection.WriteJSON(WebsocketMessage{
					Action:  "error",
					Message: "Game has already started",
				})
				mutex.Unlock()
				return
			}

			table.NumberOfPlayers += 1
			table.Clients = append(table.Clients, &connectedPlayer)
			mutex.Unlock()

			if table.NumberOfPlayers == maxNumberOfPlayers {
				err := connection.WriteJSON(WebsocketMessage{
					Action: "start_game",
				})

				if err != nil {
					log.Fatalf("❌ Failed to send JSON message: start game")
				}

				return
			}

			connection.WriteJSON(WebsocketMessage{
				Action:       "update_waiting_lobby",
				TableDetails: *table,
			})

		case "get_deck":
			if err := checkTableId(tableId); err != nil {
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
			table, err := getTable(tableId)
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
			if err := checkTableId(tableId); err != nil {
				WriteWsMessage(connection, WebsocketMessage{
					Action:  "error",
					Message: err.Error(),
				})
				return
			}

			table, err := getTable(tableId)
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
					card := logic.FlipCard(table.CurrentDeck, 1)
					logic.AttributeCardToPlayer(card, player)
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

				err := connection.WriteJSON(WebsocketMessage{
					Action: "new_round",
				})

				if err != nil {
					log.Fatalf("❌ Failed to send JSON message: go to next round")
					return
				}

				currentRound += 1
			}

		default:
		}
	}
}
