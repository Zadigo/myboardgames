package http

import (
	"errors"
	"log/slog"
	"myboardgames/minimalgo/internal/models"
	"net/http"
)

type BaseHandlers struct {
	app *HttpApp
}

func (h *BaseHandlers) SetApp(app models.AppInterface) {
	h.app = app.(*HttpApp)
}

func (b *BaseHandlers) JoinWebsite(w http.ResponseWriter, r *http.Request) {
	conn, err := CustomRequestUpgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, ConnectionUpgradeFailError.Error(), http.StatusInternalServerError)
		return
	}
	WsMiddleware(conn)
	client := NewWebsocketClient(conn)

	if b.app == nil {
		http.Error(w, "An internal server error occured", http.StatusInternalServerError)
		return
	}

	if err := b.app.AddClient(client); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer b.app.RemoveClient(client)

	// Return the routing table so that Nuxt knows what
	// rpc servers are actually registred. When client clicks
	// on a game, it pings the rpc server to check connectivity
	// and the different things related to gameplay etc.
	err = client.GetConn().WriteJSON(models.BaseWebsocketMessage{
		Action:   models.JoinWebsite,
		ClientID: client.GetID(),
		Routing:  b.app.server.GetRouter(),
	})
	if err != nil {
		slog.Error(err.Error())
		return
	}

	for {
		var message models.BaseWebsocketMessage
		err = conn.ReadJSON(&message)
		if err != nil {
			slog.Error(errors.Join(InvalidMessageError, err).Error())
			break
		}

		authenticationMessage(b.app, client, message)
		globalGameMessage(b.app, client, message)
		gameRelatedMessage(b.app, client, message)
	}
}

func (b *BaseHandlers) JoinTable(w http.ResponseWriter, r *http.Request) {}
