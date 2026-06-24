package handlers

import (
	"context"

	"github.com/Zadigo/basegame/internal/models"
	"github.com/gorilla/websocket"
)

type ProcessorOptions struct {
	Ctx     context.Context
	App     models.AppInterface
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
	gameId := options.Ctx.Value("gameId")

	switch options.Message.Action {
	case models.JOIN:
		// Handle join game logic here
		options.App.GetServerApp().JoinGame(options.Conn, gameId.(string))
	case models.START_GAME:
		// Handle start game logic here
		// options.App.GetServerApp().NotifyAll("game-uuid")
	case models.END_GAME:
		// Handle end game logic here
		// options.App.GetServerApp().EndGame("game-uuid")
	case models.START_TURN:
		// Handle start turn logic here
		// options.App.GetServerApp().NotifyAll("game-uuid")
	case models.END_TURN:
		// Handle end turn logic here
		// result := options.App.GetServerApp().NotifyAll("game-uuid")
	case models.FLIP_CARD:
		// Handle flip card logic here
		// Two options: card is special, indicate to client to resolve the card, else, just flip the card and notify all players
		// result := options.App.GetServerApp().GetGameApp().DrawCard("game-uuid")
	case models.RESOLVE_CARD:
		// Handle resolve card logic here
		// options.App.GetServerApp().NotifyAll("game-uuid")
		// if result.MustResolve {
		// 	result := options.App.GetServerApp().ResolveEvent(game.EventResolutionOptions{})
		// }
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
