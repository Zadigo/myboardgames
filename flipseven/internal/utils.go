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
	err := connection.ReadJSON(WebsocketMessage{})

	if err != nil {
		WriteWsError(connection, err)
	}

	return message
}

// Return a message to the client
func WriteWsMessage(connection *websocket.Conn, message WebsocketMessage) {
	err := connection.WriteJSON(message)

	if err != nil {
		log.Fatalf("❌ Could not send message: %v", err.Error())
		return
	}
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
