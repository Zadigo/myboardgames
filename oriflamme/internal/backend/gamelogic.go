package backend

import "log"

func AuthenticationLogic(message WebsocketMessage, client *WebsocketClient, serverRegistry *ServerRegistry) {
	switch message.Action {
	case "identify":
		// Update the client's username in the server registry
		client, exists := serverRegistry.GetClient(message.PlayerUuid)

		if exists {
			client.Username = message.Username
			client.SendJsonMessage(WebsocketMessage{
				Action:  "identify",
				Message: "Identification successful",
			})
		} else {
			log.Printf("❌ Client with UUID %s not found for identification", message.PlayerUuid)
			client.SendJsonMessage(WebsocketMessage{Action: "error", Message: "Client was not found"})
		}
	default:
		// Handle unrecognized actions
	}
}

func GameLogic(message WebsocketMessage, client *WebsocketClient, serverRegistry *ServerRegistry) {
	switch message.Action {
	case "create_game":
		gameRegistry, state := serverRegistry.CreateGame(client)
		if !state {
			// Handle the error case where the game could not be created (e.g., invalid player UUID)
		}
		gameRegistry.JoinTable(client)
		gameRegistry.broadcast <- WebsocketMessage{
			Action:  "create_game",
			Message: "Game created successfully",
		}
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
			client.SendJsonMessage(WebsocketMessage{})
		}
	case "resolve_queue":
		// Do something to resolve the influence queue
	case "end_game":
		// Do something to end the game
	default:
		// Handle unrecognized actions
	}
}
