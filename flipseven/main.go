package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

// import (
// 	"github.com/Zadigo/flipseven/cards"
// )

// const NumberOfPlayers int = 12

// const numberOfCards int = 179

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

type WebsocketMessage interface {
}

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

func liveGameHandler(response http.ResponseWriter, request *http.Request) {
	connection, err := upgrader.Upgrade(response, request, nil)

	if err != nil {
		log.Panic("❌ Failed to upgrade request")
		return
	}

	for {
		messageType, content, err := connection.ReadMessage()

		if err != nil {
			connection.WriteJSON(WebsocketMessage{})
			return
		}

		
	}
}

func main() {
	http.HandleFunc("/ws/flip-seven", Cors(liveGameHandler))

	err := http.ListenAndServe(":9000", nil)

	if errors.Is(err, http.ErrServerClosed) {
		log.Print("❌ Server closed")
	} else {
		log.Fatal("❌ Could not start server")
		os.Exit(1)
	}
}
