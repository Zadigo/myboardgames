package handlers

import (
	"log"
	"net/http"

	"github.com/Zadigo/basegame/internal/models"
)

func (g *GenericHandler) JoinGameHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("JoinGameHandler called")

	conn, err := CustomRequestUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("❌ Failed to upgrade connection: %v", err)
		http.Error(w, "Failed to upgrade connection", http.StatusInternalServerError)
		return
	}

	WsMiddleware(conn)

	defer func() {
		conn.Close()
	}()

	for {
		var message models.WebsocketMessage
		err = conn.ReadJSON(&message)

		if err != nil {
			log.Println("❌ Read error:", err)
			break
		}

		AuthMessageProcessor(conn, message)
		GameMessageProcessor(conn, message)
	}
}

func (g *GenericHandler) ObserveGameHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ObserveGameHandler called")

	conn, err := CustomRequestUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("❌ Failed to upgrade connection: %v", err)
		http.Error(w, "Failed to upgrade connection", http.StatusInternalServerError)
		return
	}

	WsMiddleware(conn)

	defer func() {
		conn.Close()
	}()

	for {
		var message models.WebsocketMessage
		err = conn.ReadJSON(&message)

		if err != nil {
			log.Println("❌ Read error:", err)
			break
		}

		g.app.GetGameApp()

		AuthMessageProcessor(conn, message)
		ObserveMessageProcessor(conn, message)
	}
}
