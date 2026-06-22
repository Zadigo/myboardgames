package handlers

import (
	"github.com/Zadigo/basegame/internal/models"
	"github.com/gorilla/websocket"
)

type ProcessorOptions struct {
	Conn    *websocket.Conn
	Message models.WebsocketMessage
	Errors  []string
}

// AuthMessageProcessor processes messages that require identification
func AuthMessageProcessor(options ProcessorOptions) {
	switch options.Message.Action {
	case models.MUST_IDENTIFY:
		// Handle identification logic here
	default:
		options.Errors = append(options.Errors, "Unrecognized action for AuthMessageProcessor")
	}
}

func GameMessageProcessor(options ProcessorOptions) {
	switch options.Message.Action {
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
	case models.RESOLVE_CARD:
		// Handle resolve card logic here
	default:
		options.Errors = append(options.Errors, "Unrecognized action for GameMessageProcessor")
	}
}

func ObserveMessageProcessor(options ProcessorOptions) {
	switch options.Message.Action {
	case models.JOIN:
		// Handle observe game logic here
	case models.OBSERVE:
		// Handle observe game logic here
	default:
		options.Errors = append(options.Errors, "Unrecognized action for ObserveMessageProcessor")
	}
}

func ErrorProcessor(conn *websocket.Conn, errors []string) {
	if len(errors) > 0 {
		for _, err := range errors {
			conn.WriteJSON(map[string]string{"error": err})
		}
	}
}
