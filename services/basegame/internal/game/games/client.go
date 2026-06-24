package games

import (
	"sync"

	"github.com/Zadigo/basegame/internal/models"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// WebsocketClient represents a client connected to
// the server via a websocket connection. It adds a layer of
// abstraction over the raw websocket connection, allowing for
// easier message handling and communication with the client.
type WebsocketClient struct {
	Uuid   string          `json:"uuid"`
	Player *Player         `json:"player"`
	conn   *websocket.Conn `json:"-"`
	mu     sync.Mutex      `json:"-"`
}

func (c *WebsocketClient) GetUuid() string {
	return c.Uuid
}

func (c *WebsocketClient) SetConn(conn *websocket.Conn) {
	c.conn = conn
}

func (c *WebsocketClient) SendJsonMessage(message models.WebsocketMessage) error {
	return c.conn.WriteJSON(message)
}

func (c *WebsocketClient) ReceiveJsonMessage() (models.WebsocketMessage, error) {
	var message models.WebsocketMessage
	err := c.conn.ReadJSON(&message)
	return message, err
}

func NewWebsocketClient(conn *websocket.Conn) *WebsocketClient {
	return &WebsocketClient{
		Uuid: uuid.NewString(),
		conn: conn,
		// Player: &game.Player{},
	}
}
