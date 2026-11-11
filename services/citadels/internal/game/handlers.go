package game

import (
	"log"
	"net/http"

	"github.com/Zadigo/citadels/internal/models"
)

type GameHandler struct{}

func (h *GameHandler) Get(w http.ResponseWriter, r *http.Request) {
	conn, err := CustomRequestUpgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to upgrade", http.StatusInternalServerError)
		return
	}

	WsMiddleware(conn)

	message := models.WebsocketMessage{Action: "must_identify"}
	conn.WriteJSON(message)
	defer conn.Close()

	for {
		var message models.WebsocketMessage
		err = conn.ReadJSON(&message)

		if err != nil {
			log.Println("❌ Read error:", err)
			break
		}

		log.Println("Received:", message.Action)
	}
}
