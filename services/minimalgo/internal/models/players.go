package models

import (
	"sync"
	"sync/atomic"

	"github.com/gorilla/websocket"
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
