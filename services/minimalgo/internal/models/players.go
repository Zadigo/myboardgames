package models

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
)

const (
	ReadTimeout    = 60 * time.Second
	PingPeriod     = 25 * time.Second // must be well below readTimeout
	WriteTimeout   = 10 * time.Second
	MaxMessageSize = 1024
)

type Player struct{}

type WebsocketClient struct {
	Player
	ID       string
	conn     *websocket.Conn
	isActive atomic.Bool
	mu       sync.Mutex
}

func (w *WebsocketClient) GetID() string {
	return w.ID
}

func (w *WebsocketClient) SetConn(conn *websocket.Conn) {
	w.mu.Lock()
	w.conn = conn
	w.mu.Unlock()
	w.isActive.Store(true)
}

func (w *WebsocketClient) GetConn() *websocket.Conn {
	return w.conn
}

func (w *WebsocketClient) IsActive() bool {
	return w.isActive.Load()
}

// WriteJSON is an enhanced version of conn.WriteJSON with timeout support
// that sets a write deadline before sending the message
func (w *WebsocketClient) WriteJSON(v BaseWebsocketMessage) error {
	w.mu.Lock()
	defer w.mu.Unlock()
	if err := w.conn.SetWriteDeadline(time.Now().Add(WriteTimeout)); err != nil {
		return err
	}
	return w.conn.WriteJSON(v)
}
