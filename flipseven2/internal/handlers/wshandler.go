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
	log.Print("🚀 Starting Flip 7 engine...")

	connection, _ := RequestUpgrader.Upgrade(response, request, nil)
	defaultContext := request.Context()
	// defaultContext := context.Background()

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

			redisClient.HSet(defaultContext, tableLayer.Layer.GetUuid(), []string{"Owner", message.Username, "CreatedAt", time.Now().String()})

			b := broadcastRegistry.GetOrCreate(tableLayer.Layer.GetUuid())
			b.AddClient(connection)

			log.Printf("🟢 Table created: %s", tableLayer.Layer.GetUuid())
			player.WriteJson(models.WebsocketMessage{
				Action:       "table_initiated",
				TableDetails: tableLayer,
			})

		// case "accept_player":
		// 	tableId := ""
		// 	tableLayer, state := baseRegistry.Get(tableId)
		// 	if !state {
		// 		// TODO:
		// 	}
		// 	if tableLayer.HasPlayer(connection) {
		// 		// TODO:
		// 	}

		// 	tableLayer.AddPlayer(connection, "Another Username", false)
		// 	b := broadcastRegistry.GetOrCreate(tableId)
		// 	b.AddClient(connection)
		// 	b.Publish("accept_player", broadcasting.Message{
		// 		Action:  "accept_player",
		// 		TableId: tableId,
		// 		// Payload: map[string]any{
		// 		// 	"Something": "a",
		// 		// },
		// 	})

		// case "flip_card":
		// 	tableId := ""
		// 	tableLayer, state := baseRegistry.Get(tableId)

		// 	if !state {
		// 		// TODO:
		// 	}

		// 	player := tableLayer.GetPlayer(connection)
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

		// case "give_card":
		// 	tableId := ""
		// 	tableLayer, state := baseRegistry.Get(tableId)

		// 	if !state {
		// 	}

		// 	card, err := tableLayer.GetCurrentCard()
		// 	if err != nil {
		// 	}

		// 	if tableLayer.CurrentPlayer != nil {
		// 		card.SetOwner(tableLayer.CurrentPlayer)

		// 		if card.IsSpecial {
		// 			if card.Category == "Freeze" {
		// 				tableLayer.FreezePlayer(tableLayer.CurrentPlayer)
		// 				tableLayer.NextPlayer(broadcastRegistry)
		// 				return
		// 			}

		// 			if card.Category == "Flip 3" {
		// 				return
		// 			}
		// 		}
		// 	}

		default:
			log.Printf("⚠️ Unrecognized action: %s", message.Action)
		}
	}
}
