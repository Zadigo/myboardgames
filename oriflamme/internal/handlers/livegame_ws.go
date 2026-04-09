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

	// go client.Broadcast()

	for {
		var message backend.WebsocketMessage
		err = conn.ReadJSON(&message)

		if err != nil {
			log.Println("❌ Read error:", err)
			break
		}

		backend.AuthenticationLogic(message, client, serverRegistry)
		backend.GameLogic(message, client, serverRegistry)
	}
}
