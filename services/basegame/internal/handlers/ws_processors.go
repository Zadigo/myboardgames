package handlers

import (
	"github.com/Zadigo/basegame/internal/models"
	"github.com/gorilla/websocket"
)

// AuthMessageProcessor processes messages that require identification
func AuthMessageProcessor(conn *websocket.Conn, message models.WebsocketMessage) {
	switch message.Action {
	case models.MUST_IDENTIFY:
		// Handle identification logic here
	default:
		// Handle other actions or log unrecognized actions
	}
}

func GameMessageProcessor(conn *websocket.Conn, message models.WebsocketMessage) {
	switch message.Action {
	case models.JOIN:
		// Handle join game logic here
	case models.START_GAME:
		// Handle start game logic here
	case models.END_GAME:
		// Handle end game logic here
	case models.START_TURN:
		// Handle start turn logic here
	case models.END_TURN:
		// Handle end turn logic here
	case models.FLIP_CARD:
		// Handle flip card logic here
	case models.GIVE_CARD:
		// Handle give card logic here
	default:
		// Handle other actions or log unrecognized actions
	}
}

func ObserveMessageProcessor(conn *websocket.Conn, message models.WebsocketMessage) {
	switch message.Action {
	case models.JOIN:
		// Handle observe game logic here
	case models.OBSERVE:
		// Handle observe game logic here
	default:
		// Handle other actions or log unrecognized actions
	}
}
