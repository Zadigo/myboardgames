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

func (c *WebsocketClient) SendJsonMessage(message WebsocketMessage) error {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	return c.Conn.WriteJSON(message)
}

func (c *WebsocketClient) ReadJsonMessage(message *WebsocketMessage) error {
	c.Mu.RLock()
	defer c.Mu.RUnlock()
	return c.Conn.ReadJSON(message)
}

func (c *WebsocketClient) GetPlayer() *Player {
	return c.Player.(*Player)
}

type WebsocketClientInterface[T any] interface {
	GetPlayer() *Player
	SendJsonMessage(message T) error
	ReadJsonMessage(message *T) error
}

func NewWebsocketClient(username string, conn *websocket.Conn, isNormalGame bool) WebsocketClientInterface[WebsocketMessage] {
	player := NewPlayer(username, isNormalGame)
	return &WebsocketClient{
		Player: player,
		Conn:   conn,
		Mu:     &sync.RWMutex{},
	}
}

type WebsocketMessage struct {
	Action       string        `json:"action"`
	PlayerUuid   string        `json:"playerUuid"`
	Message      string        `json:"message"`
	PlayingTable *PlayingTable `json:"playingTable,omitempty"`
	Username     string        `json:"username,omitempty"`
	IsNormalGame bool          `json:"isNormalGame,omitempty"`
}
