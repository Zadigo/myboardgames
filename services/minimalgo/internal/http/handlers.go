package http

import (
	"context"
	"errors"
	"log/slog"
	"myboardgames/minimalgo/internal/models"
	"net/http"
	"time"
)

type BaseHandlers struct {
	app *HttpApp
}

// func (b *BaseHandlers) getClientMessages(client models.WebsocketClientInterface) {
// 	incoming := b.app.server.GetInternalChannel()
// 	for msg := range incoming {
// 		if msg.Client == client {

// 		}
// 	}
// }

func (h *BaseHandlers) SetApp(app models.AppInterface) {
	h.app = app.(*HttpApp)
}

func (b *BaseHandlers) send(client models.WebsocketClientInterface, value string) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	message := models.InternalMessage{Client: client, Sender: models.HttpAppSender, StrMessage: value}
	if err := b.app.server.WriteInternalMessage(ctx, message); err != nil {
		slog.Warn("internal message dropped", "err", err)
	}
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

	b.send(client, models.ClientConnected)
	defer b.send(client, models.ClientDisconnected)

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
