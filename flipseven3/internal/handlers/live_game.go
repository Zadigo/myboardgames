package handlers

import (
	"log"

	"github.com/Zadigo/flipseven3/internal/backend/logic"
)

func getGame(message logic.WebsocketMessage, serverRegistry *logic.ServerRegistry) (*logic.GameRegistry, error) {
	game, err := serverRegistry.GetGame(message.TableUuid)
	if err != nil {
		log.Printf("❌ Error retrieving game: %s", message.TableUuid)

		err := game.PublishToRoom(nil, logic.WebsocketMessage{
			Action:    "error",
			TableUuid: message.TableUuid,
		})

		if err != nil {
			log.Printf("❌ Error publishing flip_card error message: %v", err)
		}

		return nil, err
	}
	return game, nil
}

func GameHandler(message logic.WebsocketMessage, client *logic.WebsocketClient, serverRegistry *logic.ServerRegistry) {
	switch message.Action {
	case "create_game":
		_, state := serverRegistry.CreateGame(client)
		if state {
			client.SendJsonMessage(logic.WebsocketMessage{
				Action:     "game_created",
				PlayerUuid: client.Uuid,
				Message:    "Game created successfully! Waiting for players to join.",
			})
		}

	case "start_game":
		game, err := getGame(message, serverRegistry)
		if err != nil {
			return
		}

		err = game.StartGame()
		if err != nil {
			log.Printf("❌ Error starting game: %v", err)
		}

	case "end_game":
		game, err := getGame(message, serverRegistry)
		if err != nil {
			return
		}
		game.EndGame()

	case "flip_card":
		game, err := getGame(message, serverRegistry)
		if err != nil {
			return
		}

		player, state := serverRegistry.GetClient(client.Uuid)
		if state {
			card := game.CardQueue.FlipCard(player)
			if card.IsSpecial {
				switch card.Category {
				case "Freeze":
					// Handle Freeze card effect
					game.PublishToRoom(nil, logic.WebsocketMessage{
						Action:     "must_choose_option",
						Message:    "Player must decide what to do with the Freeze card",
						PlayerUuid: client.Uuid,
					})

				case "Flip 3":
					// Handle Flip 3 card effect
					game.PublishToRoom(nil, logic.WebsocketMessage{
						Action:     "must_choose_option",
						Message:    "Player must decide what to do with the Flip 3 card",
						PlayerUuid: client.Uuid,
					})

				case "Second Chance":
					player.SetFlipCard(game.CardQueue, card)
					game.PublishToRoom(nil, logic.WebsocketMessage{
						Action:      "second_chance_used",
						Message:     "Player used Second Chance card",
						PlayerUuid:  client.Uuid,
						CardsInPlay: []*logic.Card{card},
					})

				default:
					log.Printf("❌ Unrecognized special card category: %s", card.Category)
				}
			} else {

			}
		}

	case "give_card":
		game, err := getGame(message, serverRegistry)
		if err != nil {
			return
		}

		otherClient := serverRegistry.GetClientByUuid(message.GiveCardToPlayer)
		if otherClient != nil {
			card, err := game.GetCardByUuid(message.GiveCardUuid)
			if err != nil {
				log.Printf("❌ Error retrieving card by UUID: %v", err)
				return
			}

			otherClient.SetFlipCard(game.CardQueue, card)

			if card.IsSpecial {
				if card.Category == "Flip 3" {
					game.PublishToRoom(nil, logic.WebsocketMessage{
						Action:     "must_flip_three",
						Message: "Player must flip three cards consecutively",
						PlayerUuid: otherClient.Uuid,
					})
				}
			}
		}

	case "start_turn":
		// Handle the logic for starting a player's turn

	case "end_turn":
		// Handle the logic for ending the current player's turn

	default:
		// Handle unrecognized actions
	}
}
