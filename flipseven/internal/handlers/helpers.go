package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Zadigo/flipseven/internal/logic"
	"github.com/gorilla/websocket"
)

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
func GetTableLayer(tableId string) (*logic.TableLayer, error) {
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

func CheckTableId(tableId string) error {
	if tableId == "" {
		return fmt.Errorf("Table ID is required")
	}
	return nil
}

// Read an incomming websocket message
func ReadWsMessage(connection *websocket.Conn) (WebsocketMessage, error) {
	var message WebsocketMessage
	err := connection.ReadJSON(&message)

	if err != nil {
		WriteWsError(connection, err)
		return WebsocketMessage{}, err
	}

	return message, nil
}

// Return a message to the client. Returns true if the message was
// sent successfully, false otherwise
func WriteWsMessage(connection *websocket.Conn, message WebsocketMessage) error {
	err := connection.WriteJSON(message)

	if err != nil {
		log.Fatalf("❌ Could not send message: %v", err.Error())
		connection.Close()
		return err
	}
	return nil
}

// Return an error message to the client
func WriteWsError(connection *websocket.Conn, err error) {
	formattedError := fmt.Sprintf("❌ An error occurred: %v", err.Error())

	WriteWsMessage(connection, WebsocketMessage{
		Action:  "error",
		Message: formattedError,
	})

	log.Fatal(formattedError)
}
