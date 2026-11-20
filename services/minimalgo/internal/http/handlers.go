package http

import (
	"fmt"
	"myboardgames/minimalgo/internal/models"
	"net/http"
)

type BaseHandlers struct {
	app *HttpApp
}

func (b *BaseHandlers) JoinWebsite(w http.ResponseWriter, r *http.Request) {
	conn, err := CustomRequestUpgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "upgrade failed", http.StatusInternalServerError)
		return
	}
	WsMiddleware(conn)
	client := NewWebsocketClient(conn)

	defer func() {
		b.app.RemoveClient(client)
		conn.Close()
	}()

	b.app.AddClient(client)

	for {
		var message models.BaseWebsocketMessage
		err = conn.ReadJSON(&message)
		if err != nil {
			break
		}
		fmt.Print(message)

		authenticationMessage(client, message)
		gameRelatedMessage(client, message)
	}
}

func (b *BaseHandlers) JoinTable(w http.ResponseWriter, r *http.Request) {}

func (b *BaseHandlers) ObserveTable(w http.ResponseWriter, r *http.Request) {}
