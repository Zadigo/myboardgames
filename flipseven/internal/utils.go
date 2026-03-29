package internal

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func CustomPanic(err error, message string) {
	if err != nil {
		log.Panicf("%v: %v", message, err)
	}
}

// Read an incomming websocket message
func ReadWsMessage(connection *websocket.Conn) WebsocketMessage {
	var message WebsocketMessage
	err := connection.ReadJSON(&message)

	if err != nil {
		WriteWsError(connection, err)
		return WebsocketMessage{}
	}

	return message
}

// Return a message to the client. Returns true if the message was
// sent successfully, false otherwise
func WriteWsMessage(connection *websocket.Conn, message WebsocketMessage) bool {
	err := connection.WriteJSON(message)

	if err != nil {
		log.Fatalf("❌ Could not send message: %v", err.Error())
		return false
	}
	return true
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
