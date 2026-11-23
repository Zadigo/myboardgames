package http

import (
	"context"
	"errors"
	"log/slog"
	"myboardgames/minimalgo/internal/models"
	"myboardgames/minimalgo/internal/wsclient"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type BaseHandlers struct {
	app *HttpApp
}

func (h *BaseHandlers) SetApp(app models.AppInterface) {
	h.app = app.(*HttpApp)
}

func (b *BaseHandlers) JoinWebsite(w http.ResponseWriter, r *http.Request) {
	// Check before upgrading: after a successful Upgrade the HTTP connection is
	// hijacked and http.Error can no longer be used.
	if b.app == nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	conn, err := wsclient.CustomRequestUpgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Warn("websocket upgrade failed", "err", err)
		return
	}

	defer conn.Close()

	wsclient.WsMiddleware(conn)
	client := wsclient.NewWebsocketClient(conn)

	if err := b.app.AddClient(client); err != nil {
		wsclient.CloseWith(conn, 1011, "internal error")
		return
	}

	defer b.app.RemoveClient(client)

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()
	go wsclient.KeepAlive(ctx, conn)

	// Return the routing table so that Nuxt knows what
	// rpc servers are actually registred. When client clicks
	// on a game, it pings the rpc server to check connectivity
	// and the different things related to gameplay etc.
	if err := client.WriteJSON(models.BaseWebsocketMessage{
		Action:   models.JoinWebsite,
		ClientID: client.GetID(),
		Routing:  b.app.server.GetRouter(),
	}); err != nil {
		slog.Warn("initial write failed", "err", err)
		return
	}

	const maxConsecutiveInvalid = 5
	invalid := 0

	for {
		var message models.BaseWebsocketMessage
		if err = conn.ReadJSON(&message); err != nil {
			if wsclient.IsFatalReadError(err) {
				wsclient.LogReadError(client, err)
				return // never read again from a failed connection
			}

			// Only the payload was bad, the connection is still usable
			invalid++
			if invalid >= maxConsecutiveInvalid {
				wsclient.CloseWith(conn, websocket.ClosePolicyViolation, "too many invalid messages")
				return
			}

			slog.Error(errors.Join(InvalidMessageError, err).Error())
			continue
		}

		invalid = 0

		// Any valid traffic also proves the client is alive
		_ = conn.SetReadDeadline(time.Now().Add(models.ReadTimeout))

		NewMessageBroker(b.app).execute(ctx, client, &message)
	}
}

func (b *BaseHandlers) JoinTable(w http.ResponseWriter, r *http.Request) {}
