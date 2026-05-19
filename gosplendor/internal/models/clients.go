package models

import (
	"sync"

	"github.com/gorilla/websocket"
)

// Base websocket message structure for
// communication between client and server
type WebsocketMessage struct {
	Action  string           `json:"action"`
	Message string           `json:"message"`
	client  *WebsocketClient `json:"-"`
}

func (c *WebsocketMessage) SendJsonMessage() error {
	c.client.Mu.Lock()
	defer c.client.Mu.Unlock()
	return c.client.Conn.WriteJSON(c)
}

func (c *WebsocketMessage) ReadJsonMessage() error {
	c.client.Mu.RLock()
	defer c.client.Mu.RUnlock()
	return c.client.Conn.ReadJSON(c)
}

type WebsocketClient struct {
	Conn *websocket.Conn
	Mu   *sync.RWMutex
}

type WebsocketClientInterface interface {
	SendJsonMessage(message WebsocketMessage) error
	ReadJsonMessage(message *WebsocketMessage) error
}
