package wsclient

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"myboardgames/minimalgo/internal/models"
	"net"
	"time"

	"github.com/gorilla/websocket"
)

// closeWith sends a close message to the websocket connection with code and reason
// since is hijacked, can't use http.Error, use closeWith instead to close the ws connection
func CloseWith(conn *websocket.Conn, code int, reason string) {
	_ = conn.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(code, reason),
		time.Now().Add(time.Second),
	)
}

// A decode error leaves the connection intact. Everything else
// (network error, close frame, timeout, oversized message) is fatal.
func IsFatalReadError(err error) bool {
	var syntaxErr *json.SyntaxError
	var typeErr *json.UnmarshalTypeError

	switch {
	case errors.As(err, &syntaxErr),
		errors.As(err, &typeErr),
		errors.Is(err, io.ErrUnexpectedEOF): // empty or truncated JSON
		return false
	}
	return true
}

func LogReadError(client models.WebsocketClientInterface, err error) {
	var netErr net.Error
	switch {
	case websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway):
		slog.Debug("client disconnected", "id", client.GetID())
	case errors.As(err, &netErr) && netErr.Timeout():
		slog.Info("client timed out", "id", client.GetID())
	case errors.Is(err, websocket.ErrReadLimit):
		slog.Warn("message too large", "id", client.GetID())
	default:
		slog.Warn("websocket read failed", "id", client.GetID(), "err", err)
	}
}

// Logs an error if it exists. If err is nil, does nothing
func tryError(client models.WebsocketClientInterface, err error) bool {
	if err != nil {
		slog.Warn("", "id", client.GetID(), "err", err)
		return true
	}
	return false
}

// Sends real ping control frames. Browsers answer with a pong automatically,
// and the pong handler then extends the read deadline.
func KeepAlive(ctx context.Context, conn *websocket.Conn) {
	t := time.NewTicker(models.PingPeriod)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			// WriteControl is safe to call concurrently with other write methods
			err := conn.WriteControl(websocket.PingMessage, nil, time.Now().Add(models.WriteTimeout))
			if err != nil {
				_ = conn.Close() // makes the blocked ReadJSON return, so the handler exits
				return
			}
		}
	}
}
