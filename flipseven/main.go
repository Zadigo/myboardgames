package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/Zadigo/flipseven/backend"
	"github.com/Zadigo/flipseven/cards"
	"github.com/Zadigo/flipseven/logic"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

var mutex = sync.RWMutex{}

var tables = make(map[string]*logic.PlayersTable)
var clients = make(map[*websocket.Conn]bool)

const maxNumberOfPlayers int = 12

var allowedOrigins = map[string]bool{
	"http://localhost:3000": true,
}

var upgrader = websocket.Upgrader{
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

// Handler for creating a new game table. This is used when the user creates
// a new game and needs a unique table ID that needs to be shared with other players.
func createTableHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if request.Header.Get("Content-Type") != "application/json" {
		http.Error(response, "Unsupported media type", http.StatusInternalServerError)
		return
	}

	_, err := io.ReadAll(request.Body)

	if err != nil {
		http.Error(response, "Failed to parse data", http.StatusInternalServerError)
		return
	}

	request.Header.Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	tableId := uuid.New()

	var responseData = struct {
		TableId string
	}{
		TableId: tableId.String(),
	}

	json.NewEncoder(response).Encode(responseData)
}

// Handler for the live game websocket connection. This is used for real-time
// communication between the server and the clients during the game.
func liveGameHandler(response http.ResponseWriter, request *http.Request) {
	connection, err := upgrader.Upgrade(response, request, nil)

	if err != nil {
		log.Println("❌ Failed to upgrade request", err.Error())
		return
	}

	// Reference any connections regardless of
	// whether they are associated to a table
	// or not
	mutex.Lock()
	clients[connection] = true
	log.Println("⚡️ New connection from client: 1.1.1.1")
	mutex.Unlock()

	defer func() {
		mutex.Lock()
		delete(clients, connection)
		mutex.Unlock()
		connection.Close()
	}()

	connection.WriteJSON(WebsocketMessage{
		Action:  "initial_connection",
		Message: "Connection successful!",
	})

	var currentRound int = 1

	for {
		var message WebsocketMessage
		err := connection.ReadJSON(&message)

		if err != nil {
			connection.WriteJSON(WebsocketMessage{
				Action:  "error",
				Message: err.Error(),
			})
			return
		}

		switch message.Action {
		case "waiting_lobby":
			tableId := message.TableId

			if tableId == "" {
				connection.WriteJSON(WebsocketMessage{
					Action:  "error",
					Message: "Table ID is required",
				})
				return
			}

			connectedPlayer := logic.ConnectedPlayer{
				Conn: connection,
				Details: logic.Player{
					Username:      "Some username",
					Uuid:          "Some UUID",
					TableUuid:     tableId,
					NumberOfCards: 0,
				},
			}

			// TODO: Store the ID in Redis

			// Connect any new players to the table
			mutex.Lock()
			table := tables[tableId]

			if table.GameStarted {
				mutex.Unlock()
				return
			}

			table.NumberOfPlayers += 1
			table.Clients = append(table.Clients, &connectedPlayer)
			mutex.Unlock()

			if table.NumberOfPlayers == maxNumberOfPlayers {
				connection.WriteJSON(WebsocketMessage{
					Action: "start_game",
				})
				return
			}

		case "get_deck":
			message := ReadWebsocketMessage(connection)

			// 1. Return the deck of cards
			baseDeck := cards.GetDeck()

			table := tables[message.TableId]
			table.CurrentDeck = baseDeck

			connection.WriteJSON(WebsocketMessage{
				Action: "deck_created",
				Deck:   baseDeck,
			})

		case "flip_card":
			message = ReadWebsocketMessage(connection)
			table := tables[message.TableId]

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

				connection.WriteJSON(WebsocketMessage{
					Action: "new_round",
				})

				currentRound += 1
			}

		default:
		}
	}
}

func beforeStart() *redis.Client {
	baseConfig := backend.ServerConfig{
		Config: backend.ServerBaseConfig{
			Backends: &backend.ServerBackendsConfig{
				Redis: &backend.ServerBackendConfig{
					Url: "redis://:@localhost:6379/0",
				},
			},
		},
	}

	client, err := backend.CreateRedisClient(baseConfig.Config.Backends.Redis)

	if err != nil {
		log.Panicln("❌ Failed to create Redis client", err.Error())
	}

	return client
}

func main() {
	log.Println("🚀 Starting Flip 7 Webserver...")
	log.Println("✅ Server started on 127.0.0.1:9000")

	http.HandleFunc("/ws/flip-seven", Cors(liveGameHandler))
	http.HandleFunc("/v1/flip-seven/create", Cors(createTableHandler))

	redisClient := beforeStart()
	defer redisClient.Close()

	err := http.ListenAndServe(":9000", nil)

	if errors.Is(err, http.ErrServerClosed) {
		log.Println("❌ Server closed")
	} else {
		log.Fatalln("❌ Could not start server")
	}
}
