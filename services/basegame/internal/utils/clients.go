package utils

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
type BaseWebsocketClient struct {
	Uuid   string          `json:"uuid"`
	conn   *websocket.Conn `json:"-"`
	mu     sync.Mutex      `json:"-"`
	// The send channel is used to queue messages 
	// that need to be sent to the client.
	messagesChannel chan models.WebsocketMessage `json:"-"`
}

// Setup initializes the websocket client with the given connection.
// It sets up the connection, initializes the messages channel,
// and generates a UUID if it hasn't been set already.
func (c *BaseWebsocketClient) Setup(conn *websocket.Conn) {
	c.SetConn(conn)
	c.messagesChannel = make(chan models.WebsocketMessage, 256)
	if c.Uuid == "" {
		c.Uuid = uuid.NewString()
	}
}

func (c *BaseWebsocketClient) GetUuid() string {
	return c.Uuid
}

func (c *BaseWebsocketClient) SetConn(conn *websocket.Conn) {
	c.conn = conn
}

func (c *BaseWebsocketClient) GetConn() *websocket.Conn {
	return c.conn
}

func (c *BaseWebsocketClient) SendJsonMessage(message models.WebsocketMessage) error {
	return c.conn.WriteJSON(message)
}

func (c *BaseWebsocketClient) ReceiveJsonMessage() (models.WebsocketMessage, error) {
	var message models.WebsocketMessage
	err := c.conn.ReadJSON(&message)
	return message, err
}

func (c *BaseWebsocketClient) GetMessagesChannel() chan models.WebsocketMessage {
	return c.messagesChannel
}

func (c *BaseWebsocketClient) CloseMessagesChannel() {
	close(c.messagesChannel)
}

func (c *BaseWebsocketClient) SendMessageToChannel(message models.WebsocketMessage) {
	c.messagesChannel <- message
}
