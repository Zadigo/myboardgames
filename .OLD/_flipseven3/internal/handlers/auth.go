package handlers

import (
	"log"

	"github.com/Zadigo/flipseven3/internal/backend/logic"
)

func AuthHandler(message logic.WebsocketMessage, client *logic.WebsocketClient, serverRegistry *logic.ServerRegistry) {
	switch message.Action {
	case "identify":
		// Update the client's username in the server registry
		client, exists := serverRegistry.GetClient(message.PlayerUuid)

		if exists {
			client.Username = message.Username
			client.SendJsonMessage(logic.WebsocketMessage{
				Action:  "identify",
				Message: "Identification successful",
			})
		} else {
			log.Printf("❌ Client with UUID %s not found for identification", message.PlayerUuid)
			client.SendJsonMessage(logic.WebsocketMessage{Action: "error", Message: "Client was not found"})
		}
	default:
		// Handle unrecognized actions
	}
}
