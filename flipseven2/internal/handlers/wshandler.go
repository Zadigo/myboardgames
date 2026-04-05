package handlers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/Zadigo/flipseven2/internal/backend/broadcasting"
	"github.com/Zadigo/flipseven2/internal/models"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

func GameEngine(response http.ResponseWriter, request *http.Request, redisClient *redis.Client, broadcastRegistry *broadcasting.BroadcasterRegistry, baseRegistry *models.BaseRegistry) {
	log.Print("⚡️ Processing new connection...")

	connection, _ := RequestUpgrader.Upgrade(response, request, nil)
	defaultContext := request.Context()

	err := connection.WriteJSON(models.WebsocketMessage{
		Action:  "initial_connection",
		Message: "Connection successful!",
	})

	if err != nil {
		log.Printf("⚠️ Error sending initial connection message: %v", err)
		return
	}

	var once sync.Once
	closeConnection := func() {
		once.Do(func() {
			connection.Close()
		})
	}

	go func() {
		<-defaultContext.Done()
		log.Print("⚠️ Context cancelled → closing websocket")
		closeConnection()
	}()

	defer func() {
		baseRegistry.Mu.Lock()
		delete(baseRegistry.Clients, connection)
		baseRegistry.Mu.Unlock()

		// Send proper closing handshake before closing
		err := connection.WriteMessage(
			websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
		)

		if err != nil {
			log.Printf("⚠️ Error sending close message: %v", err)
		}

		log.Printf("⚡️ Connection closed")
		closeConnection()
	}()

	for {
		message := &models.WebsocketMessage{}
		err := connection.ReadJSON(message)

		if err != nil {
			log.Printf("⚠️ Error reading JSON message: %v", err)
			return
		}

		switch message.Action {
		case "initiate_table":
			tableLayer := models.CreatePlayersTable()
			player := tableLayer.Layer.AddPlayer(connection, message.Username, true)

			baseRegistry.Set(tableLayer)
			redisClient.HSet(defaultContext, tableLayer.Layer.GetUuid(), []string{"Owner", message.Username, "CreatedAt", time.Now().String()})

			b := broadcastRegistry.GetOrCreate(tableLayer.Layer.GetUuid())
			b.AddClient(connection)

			log.Printf("🟢 Table created: %s", tableLayer.Layer.GetUuid())
			player.WriteJson(models.WebsocketMessage{
				Action:       "table_initiated",
				TableId:      tableLayer.Layer.GetUuid(),
				TableDetails: tableLayer.Layer.PrintDetails(),
			})

		case "accept_player":
			tableLayer, state := baseRegistry.Get(message.TableId)
			if !state {
				log.Printf("⚠️ Cannot accept player. Table not found: %s", message.TableId)
				return
			}

			hasPlayer := tableLayer.Layer.HasPlayer(connection)
			if hasPlayer {
				log.Print("⚠️ Player already present on the table")
				continue
			}

			tableLayer.Layer.AddPlayer(connection, message.Username, false)

			b := broadcastRegistry.GetOrCreate(message.TableId)
			b.AddClient(connection)
			err := b.Publish("accept_player", broadcasting.Message{
				Action:  "accept_player",
				TableId: message.TableId,
				// Payload: map[string]any{
				// 	"Something": "a",
				// },
			})

			if err != nil {
				log.Printf("⚠️ Error publishing accept_player message: %v", err)
			}

			player := tableLayer.Layer.GetPlayer(connection)

			log.Printf("🟢 Player accepted: %s (%s) [%d players]", player.Username, player.Uuid, tableLayer.Layer.GetNumberOfPlayers())
			player.WriteJson(models.WebsocketMessage{
				Action:       "player_accepted",
				TableId:      message.TableId,
				TableDetails: tableLayer.Layer.PrintDetails(),
			})

		case "start_game":
			tableLayer, state := baseRegistry.Get(message.TableId)

			if !state {
				log.Printf("⚠️ Cannot start game. Table not found: %s", message.TableId)
				return
			}

			tableLayer.Layer.GetDeck()

			b, state := broadcastRegistry.Get(message.TableId)
			if !state {
				log.Printf("⚠️ Cannot start game. Broadcaster not found for table: %s", message.TableId)
				return
			}

			b.Publish("deck_loaded", broadcasting.Message{
				Action:  "deck_loaded",
				TableId: message.TableId,
				Payload: map[string]any{
					"TableDetails": tableLayer.Layer.PrintDetails(),
				},
			})

		case "reconnect":
			tableLayer, state := baseRegistry.Get(message.TableId)

			if !state {
				log.Printf("❌ Table with ID %s not found in registry for reconnection", message.TableId)
				return
			}

			player := tableLayer.Layer.GetPlayer(connection)
			log.Printf("🔄 Reconnecting player %s to table %s", message.Username, message.TableId)

			player.WriteJson(models.WebsocketMessage{
				Action:       "reconnected",
				TableDetails: tableLayer.Layer.PrintDetails(),
			})

		// case "flip_card":
		// 	tableId := ""
		// 	tableLayer, state := baseRegistry.Get(tableId)

		// 	if !state {
		// 		// TODO:
		// 	}

		// 	player := tableLayer.Layer.GetPlayer(connection)
		// 	if player == nil {
		// 		// TODO:
		// 	}

		// 	action, card, err := tableLayer.FlipCard(player)
		// 	if err != nil {
		// 		// TODO:
		// 	}

		// 	player.SetPlayerCards(tableLayer)
		// 	player.CalculatePoints(card)

		// 	if action == "next_round" {
		// 		tableLayer.ResetPlayers()

		// 		b, state := broadcastRegistry.Get(tableId)
		// 		if !state {
		// 			// TODO:
		// 		}

		// 		b.Publish("flip_card", broadcasting.Message{
		// 			Action:  "flip_card",
		// 			TableId: tableId,
		// 			Payload: map[string]string{},
		// 		})

		// 		return
		// 	}

		// 	if action == "continue" {
		// 		// TODO:
		// 	}

		// 	b, state := broadcastRegistry.Get(tableId)
		// 	if !state {
		// 		// TODO:
		// 	}

		// 	b.Publish("flip_card", broadcasting.Message{
		// 		Action:  "flip_card",
		// 		TableId: tableId,
		// 		Payload: map[string]string{},
		// 	})

		case "give_card":
			tableLayer, state := baseRegistry.Get(message.TableId)

			if !state {
				log.Printf("⚠️ Cannot give card. Table not found: %s", message.TableId)
				return
			}

			card, err := tableLayer.Layer.GetCurrentCard()
			if err != nil {
				log.Printf("⚠️ Cannot give card. Error getting current card: %v", err)
				return
			}

			var receivingPlayer *models.Player
			currentPlayer := tableLayer.Layer.GetCurrentPlayer()

			if message.GivesCardTo != "" {
				receivingPlayer = tableLayer.Layer.GetPlayerByUuid(message.GivesCardTo)
			}

			// Function to apply the freeze effect to a player
			// and move to the next player
			applyFreeze := func(applyOn *models.Player) {
				tableLayer.Layer.FreezePlayer(applyOn)
				tableLayer.Layer.NextPlayer(broadcastRegistry)
			}

			if receivingPlayer != nil {
				card.SetOwner(receivingPlayer)

				if card.IsSpecial {
					if card.Category == "Freeze" {
						applyFreeze(receivingPlayer)
					}

					if card.Category == "Flip 3" {
						return
					}
				}
			} else {
				card.SetOwner(currentPlayer)

				if card.IsSpecial {
					if card.Category == "Freeze" {
						applyFreeze(currentPlayer)
						return
					}

					if card.Category == "Flip 3" {
						return
					}
				}
			}

		default:
			log.Printf("⚠️ Unrecognized action: %s", message.Action)
		}
	}
}
