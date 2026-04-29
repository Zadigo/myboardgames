package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/Zadigo/splendor/internal/backend"
	"github.com/gorilla/websocket"
)

// authenticationHandler handles messages related to client authentication, such as identifying the player.
func authenticationHandler(message backend.WebsocketMessage, client backend.WebsocketClientInterface[backend.WebsocketMessage], serverRegistry backend.ServerRegistryInterface) {
	switch message.Action {
	case "identify":
		username := message.Username
		if username == "" {
			client.SendJsonMessage(backend.WebsocketMessage{
				Action:  "error",
				Message: "Username cannot be empty.",
			})
			return
		}

		player := client.GetPlayer()
		player.Username = username
	}
}

// tableHandler handles messages related to table management, such as creating or joining a table.
func tableHandler(message backend.WebsocketMessage, client backend.WebsocketClientInterface[backend.WebsocketMessage], serverRegistry backend.ServerRegistryInterface) {
	switch message.Action {
	case "create_table":
		table := backend.NewPlayingTable(message.IsNormalGame)
		serverRegistry.AddPlayingTable(table)

		err := table.AddPlayer(client)
		if err != nil {
			client.SendJsonMessage(backend.WebsocketMessage{
				Action:  "error",
				Message: err.Error(),
			})
		}

		// Once the table is created, we can start broadcasting
		// messages to all potential clients at the table.
		table.StartBroadcaster()

		table.BroadcastMessage(backend.WebsocketMessage{
			Action:       "table_created",
			Message:      "Table created successfully",
			PlayingTable: table,
		})
	}
}

// func gameHandler(message backend.WebsocketMessage, client backend.WebsocketClientInterface[backend.WebsocketMessage], serverRegistry backend.ServerRegistryInterface) {
// 	switch message.Action {
// 	case "identify":
// 		username := message.Username
// 		if username == "" {
// 			client.SendJsonMessage(backend.WebsocketMessage{
// 				Action:  "error",
// 				Message: "Username cannot be empty.",
// 			})
// 			return
// 		}

// 		player := client.GetPlayer()
// 		player.Username = username
// 	}
// }

func LiveGameHandler(w http.ResponseWriter, r *http.Request, serverRegistry backend.ServerRegistryInterface) {
	conn, err := CustomRequestUpgrader.Upgrade(w, r, nil)

	if err != nil {
		http.Error(w, "Failed to upgrade to WebSocket", http.StatusInternalServerError)
		return
	}

	conn.SetReadLimit(1024)
	conn.SetReadDeadline(time.Now().Add(60 * time.Second))

	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	// Ping ticker: Sends a ping every 30s to keep the connection alive
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}()

	client := backend.NewWebsocketClient("temp_username", conn, true)
	err = serverRegistry.AddClient(client)

	defer func() {
		serverRegistry.RemoveClient(client)
		conn.Close()
	}()

	if err != nil {
		log.Println("❌ AddClient error:", err)
		return
	}

	client.SendJsonMessage(backend.WebsocketMessage{
		Action:     "must_identify",
		PlayerUuid: client.GetPlayer().Uuid,
		Message:    "Connection established successfully! Please identify yourself with your username.",
	})

	for {
		var message backend.WebsocketMessage
		err = client.ReadJsonMessage(&message)

		if err != nil {
			if IsWebsocketClose(err) {
				log.Println("🔌 Client disconnected:", client.GetPlayer().Uuid)
			} else {
				log.Println("❌ Read error:", err)
			}
			break
		}

		authenticationHandler(message, client, serverRegistry)
		tableHandler(message, client, serverRegistry)
	}
}
