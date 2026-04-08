package handlers

import (
	"net/http"

	"github.com/Zadigo/oriflamme/internal/backend"
)

func LiveGameHandler(w http.ResponseWriter, r *http.Request, serverRegistry *backend.ServerRegistry) {
	conn, err := CustomRequestUpgrader.Upgrade(w, r, nil)

	if err != nil {
		return
	}

	defer func() {
		conn.Close()

		// connection.WriteMessage(
		// 	websocket.CloseMessage,
		// 	websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
		// )
	}()

	for {
		var message backend.WebsocketMessage
		err = conn.ReadJSON(&message)

		if err != nil {
			// conn.WriteJSON(backend.WebsocketMessage{})
			continue
		}

		// Add the new client to the server registry
		clientUuid, client := serverRegistry.AddClient(message.Username, conn)

		switch message.Action {
		case "create_game":
			gameRegistry, state := serverRegistry.CreateGame(clientUuid)
			if !state {
				// Handle the error case where the game could not be created (e.g., invalid player UUID)
			}

			gameRegistry.JoinTable(client)
			client.SendJsonMessage(backend.WebsocketMessage{Action: "create_game", Message: "Simple message"})
		case "start_game":
			// 1. Join the player to the game table using the provided table UUID
			// 1. Create all the redis dependencies for the game
			// 2. Send a message to all players that the game has started
			client.SendJsonMessage(backend.WebsocketMessage{})
		case "play_card":
			// Do something to play a card
			switch message.CardAction {
			case "reveal":
				// Do something to reveal a card
			case "place_token":
				// Do something to place a token on a card
			case "stack_card":
				// Do something to stack a card on top of another card
			default:
				client.SendJsonMessage(backend.WebsocketMessage{})
			}
		case "end_game":
			// Do something to end the game
		default:
			client.SendJsonMessage(backend.WebsocketMessage{})
		}
	}
}
