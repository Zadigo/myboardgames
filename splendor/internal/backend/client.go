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
	return nil
}

func (c *WebsocketClient) ReadJsonMessage() WebsocketMessage {
	return WebsocketMessage{}
}

func (c *WebsocketClient) GetPlayer() *Player {
	return c.Player.(*Player)
}

type WebsocketClientInterface[T any] interface {
	GetPlayer() *Player
	SendJsonMessage(message T) error
	ReadJsonMessage() T
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
}

type GameRegistry struct {
}
