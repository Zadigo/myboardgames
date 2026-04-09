package backend

import (
	"log"
	"slices"
)

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
			return
		}

		log.Printf("🟢 New table created: %s", gameRegistry.Uuid)

		gameRegistry.JoinTable(client)
		log.Printf("🟢 Client %s joined table %s", client.Uuid, gameRegistry.Uuid)

		gameRegistry.broadcast <- WebsocketMessage{
			Action:    "create_game",
			TableUuid: gameRegistry.Uuid,
			Message:   "Game created successfully",
		}

		// TEST:
		// err := gameRegistry.PublishToRoom(serverRegistry.redisClient, WebsocketMessage{
		// 	Action:  "test_redis_publication",
		// 	Message: "Gaga",
		// })
		// err = gameRegistry.PublishToRoom(serverRegistry.redisClient, WebsocketMessage{
		// 	Action:  "test_redis_publication_two",
		// 	Message: "Lady",
		// })
		// if err != nil {
		// 	log.Println("❌ Failed to publish message to Redis channel:", err)
		// }
	case "start_game":
		gameRegistry, err := serverRegistry.GetGame(message.TableUuid)
		if err != nil {
			client.SendJsonMessage(WebsocketMessage{Action: "error", Message: "Game not found"})
			log.Print("🔴 Game registry does not exist")
			return
		}

		// 1. Start the game
		err = gameRegistry.StartGame(message.TableUuid)
		if err != nil {
			client.SendJsonMessage(WebsocketMessage{Action: "error", Message: err.Error()})
			log.Print("🔴 Could not start non-existent game")
			return
		}

		log.Printf("⚡️ Game started: %s", gameRegistry.Uuid)

		// 2. Create all the redis dependencies for the game

		// 3. Send a message to all players that the game has started
	case "select_cards":
		gameRegistry, err := serverRegistry.GetGame(message.TableUuid)
		if err != nil {
			client.SendJsonMessage(WebsocketMessage{Action: "error", Message: "Game not found"})
			log.Print("🔴 Game registry does not exist")
			return
		}

		for _, card := range gameRegistry.CardsInPlay {
			if slices.Contains(message.SelectedCards, card.Uuid) {
				card.IsSelected = true
			} else {
				card.IsSelected = false
			}
		}

	case "play_card":
		gameRegistry, err := serverRegistry.GetGame(message.TableUuid)
		if err != nil {
			client.SendJsonMessage(WebsocketMessage{Action: "error", Message: "Game not found"})
			log.Print("🔴 Game registry does not exist")
			return
		}

		switch message.CardAction {
		case "place_card_left":
			card, err := gameRegistry.GetCardByUuid(message.CardUuid)
			if err != nil {
				client.SendJsonMessage(WebsocketMessage{Action: "error", Message: err.Error()})
				log.Printf("🟠 Card with UUID %s not found (table: %s)", message.CardUuid, gameRegistry.Uuid)
				return
			}

			gameRegistry.InfluenceQueueLayer.AddCardLeft(card)
			gameRegistry.broadcast <- WebsocketMessage{
				Action: "place_card",
				Queue:  gameRegistry.InfluenceQueueLayer,
			}
		case "place_card_right":
			card, err := gameRegistry.GetCardByUuid(message.CardUuid)
			if err != nil {
				client.SendJsonMessage(WebsocketMessage{Action: "error", Message: err.Error()})
				log.Printf("🟠 Card with UUID %s not found (table: %s)", message.CardUuid, gameRegistry.Uuid)
				return
			}

			gameRegistry.InfluenceQueueLayer.AddCardRight(card)
			gameRegistry.broadcast <- WebsocketMessage{
				Action: "place_card",
				Queue:  gameRegistry.InfluenceQueueLayer,
			}
		case "reveal":
			// Do something to reveal a card
		case "place_token":
			// Do something to place a token on a card
		case "stack_card":
			// Do something to stack a card on top of another card
		default:
			client.SendJsonMessage(WebsocketMessage{Action: "error", Message: "Invalid card action"})
		}
	case "resolve_queue":
		// Do something to resolve the influence queue
	case "end_game":
		// Do something to end the game
	default:
		client.SendJsonMessage(WebsocketMessage{Action: "error", Message: "Action not recognized"})
	}
}
