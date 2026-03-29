package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func CustomPanic(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func ReadWebsocketMessage(connection *websocket.Conn) WebsocketMessage {
	var message WebsocketMessage
	err := connection.ReadJSON(WebsocketMessage{})

	if err != nil {
		connection.WriteJSON(WebsocketMessage{
			Action:  "error",
			Message: fmt.Sprintf("Could not read JSON message: %s", err.Error()),
		})
	}

	return message
}
