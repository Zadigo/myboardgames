package internal

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func CustomPanic(err error, msg string) {
	if err != nil {
		log.Panicf("%v: %v", msg, err)
	}
}

func ReadWsMessage(connection *websocket.Conn) WebsocketMessage {
	var message WebsocketMessage
	err := connection.ReadJSON(WebsocketMessage{})

	if err != nil {
		connection.WriteJSON(WebsocketMessage{
			Action:  "error",
			Message: fmt.Sprintf("Could not read JSON message: %v", err.Error()),
		})
	}

	return message
}

func WriteWsMessage(connection *websocket.Conn, message WebsocketMessage) {
	err := connection.WriteJSON(message)

	if err != nil {
		log.Fatalf("❌ Could not send message: %v", err.Error())
		return
	}
}
