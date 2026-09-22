package games

import (
	"github.com/Zadigo/basegame/internal/utils"
	"github.com/gorilla/websocket"
)

// WebsocketClient represents a client connected to
// the server via a websocket connection. It adds a layer of
// abstraction over the raw websocket connection, allowing for
// easier message handling and communication with the client.
type WebsocketClient struct {
	utils.BaseWebsocketClient
	Player *Player         `json:"player"`
}

func NewWebsocketClient(conn *websocket.Conn) *WebsocketClient {
	client := &WebsocketClient{}
	client.Setup(conn)
	return client
}
