package handlers

import (
	"context"
	"net/http"

	"github.com/Zadigo/flipseven2/internal/backend"
	"github.com/Zadigo/flipseven2/internal/models"
	"github.com/redis/go-redis/v9"
)

func GameEngine(response http.ResponseWriter, request *http.Request, ctx context.Context, redisConn *redis.Client, broadcaster *backend.BroadcasterRegistry, baseRegistry *models.BaseRegistry) {
	connection, _ := RequestUpgrader.Upgrade(response, request, nil)

	for {
		message := ""

		switch message {
		case "initiate_table":
			tableLayer := models.CreatePlayersTable()
			tableLayer.Layer.AddPlayer(connection, "Some Username", true)
			redisConn.HSet(tableLayer.Layer)

		case "accept_player":
			tableId := ""
			tableLayer, state := baseRegistry.Get(tableId)
			if !state {
				// TODO:
			}
			if tableLayer.HasPlayer(connection) {
				// TODO:
			}

			tableLayer.AddPlayer(connection, "Another Username", false)

		case "flip_card":
			tableId := ""
			tableLayer, state := baseRegistry.Get(tableId)

			if !state {
				// TODO:
			}

			player := tableLayer.GetPlayer(connection)
			if player == nil {
				// TODO:
			}

			action, card, err := tableLayer.FlipCard(player)
			if err != nil {
				// TODO:
			}

			player.SetPlayerCards(tableLayer)
			player.CalculatePoints(card)

			if action == "next_round" {
				tableLayer.ResetPlayers()
			}

			if action == "continue" {
				// TODO:
			}

		case "give_card":
			tableId := ""
			tableLayer, state := baseRegistry.Get(tableId)

			if !state {
			}

			card, err := tableLayer.GetCurrentCard()
			if err != nil {
			}

			if tableLayer.CurrentPlayer != nil {
				card.SetOwner(tableLayer.CurrentPlayer)

				if card.IsSpecial {
					if card.Category == "Freeze" {
						tableLayer.FreezePlayer(tableLayer.CurrentPlayer)
						tableLayer.NextPlayer(broadcaster)
						return
					}

					if card.Category == "Flip 3" {
						return
					}
				}
			}
		}
	}
}
