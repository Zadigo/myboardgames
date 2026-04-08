package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/Zadigo/oriflamme/internal/backend"
)

func LiveGameHandler(w http.ResponseWriter, r *http.Request, serverRegistry *backend.ServerRegistry) {
	conn, err := CustomRequestUpgrader.Upgrade(w, r, nil)

	if err != nil {
		return
	}

	// Setup connection safety measures (e.g., origin check, authentication) here if needed
	conn.SetReadLimit(1024)
	conn.SetReadDeadline(time.Now().Add(60 * time.Second))

	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	defer func() {
		conn.Close()

		// connection.WriteMessage(
		// 	websocket.CloseMessage,
		// 	websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
		// )
	}()

	// Add the new client to the server registry
	clientUuid, client := serverRegistry.AddClient(conn)
	client.SendJsonMessage(backend.WebsocketMessage{
		Action:     "must_identify",
		PlayerUuid: clientUuid,
	})

	for {
		var message backend.WebsocketMessage
		err = conn.ReadJSON(&message)

		if err != nil {
			log.Println("❌ Read error:", err)
			break
		}

		switch message.Action {
		case "identify":
			// Update the client's username in the server registry
			client, exists := serverRegistry.GetClient(message.PlayerUuid)

			if exists {
				client.Username = message.Username
			} else {
				log.Printf("❌ Client with UUID %s not found for identification", message.PlayerUuid)
				client.SendJsonMessage(backend.WebsocketMessage{Action: "error", Message: "Client was not found"})
			}
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
			// client.SendJsonMessage(backend.WebsocketMessage{})
		case "play_card":
			// Do something to play a card
			switch message.CardAction {
			case "place_card":
				// Do something to place a card on the board
				// This could involve updating the game state in Redis and sending a message to all players about the new card placement
			case "reveal":
				// Do something to reveal a card
			case "place_token":
				// Do something to place a token on a card
			case "stack_card":
				// Do something to stack a card on top of another card
			default:
				client.SendJsonMessage(backend.WebsocketMessage{})
			}
		case "resolve_queue":
			// Do something to resolve the influence queue
		case "end_game":
			// Do something to end the game
		default:
			// Handle unrecognized actions
		}
	}
}
