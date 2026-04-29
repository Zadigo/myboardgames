package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/Zadigo/splendor/internal/backend"
)

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

	func() {
		defer conn.Close()
	}()

	client := backend.NewWebsocketClient("temp_username", conn, true)
	err = serverRegistry.AddClient(client)
	
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
		err = conn.ReadJSON(&message)

		if err != nil {
			log.Println("❌ Read error:", err)
			break
		}

		// AuthHandler(message, client, serverRegistry)
		// GameHandler(message, client, serverRegistry)
	}
}
