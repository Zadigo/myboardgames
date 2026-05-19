package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Zadigo/gosplendor/internal/models"
	"github.com/gorilla/websocket"
)

type LiveGame struct {
	ctx context.Context
}

func (h *LiveGame) Connect(w http.ResponseWriter, r *http.Request) {
	conn, err := CustomRequestUpgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to upgrade connection", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	middleware := &WebsocketMiddleware{}
	middleware.Handle(conn)

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	message := models.WebsocketMessage{
		Action:  "must_identify",
		Message: "Connection established successfully! Please identify yourself with your username.",
	}

	message.SendJsonMessage()

	go func() {
		for range ticker.C {
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}()

	for {
		var message models.WebsocketMessage
		err := message.ReadJsonMessage()

		if err != nil {
			if IsWebsocketClose(err) {
				log.Println("🔌 Client disconnected:", "playerUuid")
			} else {
				log.Println("❌ Read error:", err)
			}
			break
		}
	}
}
