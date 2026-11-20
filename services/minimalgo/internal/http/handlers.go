package http

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"myboardgames/minimalgo/internal/models"
	"net/http"
	"time"
)

type BaseHandlers struct {
	app *HttpApp
}

func (b *BaseHandlers) JoinWebsite(w http.ResponseWriter, r *http.Request) {
	conn, err := CustomRequestUpgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, ConnectionUpgradeFailError.Error(), http.StatusInternalServerError)
		return
	}
	WsMiddleware(conn)
	client := NewWebsocketClient(conn)

	b.app.AddClient(client)
	defer b.app.RemoveClient(client)

	send := func(value string) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		message := models.InternalMessage{Client: client, Sender: models.HttpAppSender, StrMessage: value}
		if err := b.app.server.WriteInternalMessage(ctx, message); err != nil {
			slog.Warn("internal message dropped", "err", err)
		}
	}

	send(models.ClientConnected)
	defer send(models.ClientDisconnected)

	for {
		var message models.BaseWebsocketMessage
		err = conn.ReadJSON(&message)
		if err != nil {
			slog.Error(errors.Join(InvalidMessageError, err).Error())
			break
		}

		fmt.Print(message)

		authenticationMessage(client, message)
		gameRelatedMessage(client, message)
	}
}

func (b *BaseHandlers) JoinTable(w http.ResponseWriter, r *http.Request) {}

func (b *BaseHandlers) ObserveTable(w http.ResponseWriter, r *http.Request) {}
