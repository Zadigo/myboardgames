package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/Zadigo/flipseven/cards"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

const NumberOfPlayers int = 12

// const numberOfCards int = 179

// var tables = make(map[string]logic.ConnectedPlayer)
var clients = make(map[*websocket.Conn]bool)
var mutex = sync.RWMutex{}

// func startGame() func(int) {
// 	var currentPlayer int
// 	var players []cards.Player
// 	var gameStarted bool
// 	var remainingCards int = numberOfCards

// 	return func(numberOfPlayers int) {
// 		for {
// 			// Do something
// 		}
// 	}
// }

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

func liveGameHandler(response http.ResponseWriter, request *http.Request) {
	connection, err := upgrader.Upgrade(response, request, nil)

	if err != nil {
		log.Println("❌ Failed to upgrade request", err.Error())
		return
	}

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

	for {
		// _, content, err := connection.ReadMessage()

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
		case "initial_connection":

		case "distribute_cards":
			baseDeck := cards.GetDeck()

			connection.WriteJSON(WebsocketMessage{
				Action: "distribute_cards",
				Deck:   baseDeck,
			})

		case "start_game":

		default:
		}
	}
}

func main() {
	log.Println("🚀 Starting Flip 7 Webserver...")
	log.Println("✅ Server started on 127.0.0.1:9000")

	http.HandleFunc("/ws/flip-seven", Cors(liveGameHandler))
	http.HandleFunc("/v1/flip-seven/create", Cors(createTableHandler))

	err := http.ListenAndServe(":9000", nil)

	if errors.Is(err, http.ErrServerClosed) {
		log.Println("❌ Server closed")
	} else {
		log.Fatalln("❌ Could not start server")
	}
}
