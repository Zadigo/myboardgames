package backend

import (
	"sync"

	"github.com/gorilla/websocket"
)

// WebsocketClient represents a client connected
// to the game via websocket. It wraps the player's
// information and the websocket connection.
type WebsocketClient struct {
	Player PlayerInterface
	Conn   *websocket.Conn
	Mu     *sync.RWMutex
}

func NewWebsocketClient(username string, conn *websocket.Conn, isNormalGame bool) *WebsocketClient {
	player := NewPlayer(username, isNormalGame)
	return &WebsocketClient{
		Player: player,
		Conn:   conn,
		Mu:     &sync.RWMutex{},
	}
}

type WebsocketMessage struct{}

type GameRegistry struct {
	
}
