package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

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

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
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
		log.Panic("❌ Failed to upgrade request")
		return
	}

	mutex.Lock()
	clients[connection] = true
	mutex.RUnlock()

	defer func() {
		mutex.Lock()
		delete(clients, connection)
		mutex.RUnlock()
		connection.Close()
	}()

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

		case "start_game":

		default:
		}
	}
}

func main() {
	http.HandleFunc("/ws/flip-seven", Cors(liveGameHandler))
	http.HandleFunc("/v1/flip-seven/create", Cors(createTableHandler))

	err := http.ListenAndServe(":9000", nil)

	if errors.Is(err, http.ErrServerClosed) {
		log.Print("❌ Server closed")
	} else {
		log.Fatal("❌ Could not start server")
		os.Exit(1)
	}
}
