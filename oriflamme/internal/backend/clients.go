package backend

import (
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type WebsocketMessage struct {
	Action   string `json:"action"`
	Username string `json:"username,omitempty"`
	Message  string `json:"message,omitempty"`
	// The unique identifier for the player/client sending
	// the message. This can be used to associate messages with
	// specific players and manage game state accordingly.
	PlayerUuid string `json:"player_uuid,omitempty"`
	CardAction string `json:"card_action,omitempty"`
	// The unique identifier for the game table that the player is trying to join or create.
	// This can be used to associate players with specific games and manage game state accordingly.
	TableUuid string `json:"table_uuid,omitempty"`
	// Indicates the player/client is the one that initiated the new game
	Initiator bool `json:"initiator,omitempty"`
}

// The WebsocketClient struct represents a client that is connected to
// the server via a websocket connection.
type WebsocketClient struct {
	Uuid      string          `json:"uuid"`
	Username  string          `json:"username"`
	Initiator bool            `json:"initiator"`
	conn      *websocket.Conn `json:"-"`

	// The send channel is used to queue messages 
	// that need to be sent to the client.
	send chan WebsocketMessage
	mu   sync.Mutex
}

func (client *WebsocketClient) SendJsonMessage(message WebsocketMessage) error {
	return client.conn.WriteJSON(message)
}

func (client *WebsocketClient) ReceiveJsonMessage() (WebsocketMessage, error) {
	var message WebsocketMessage
	err := client.conn.ReadJSON(&message)
	return message, err
}

// Send messages from the send channel 
// to the websocket connection.
// func (c *WebsocketClient) Broadcast() {
// 	for msg := range c.send {
// 		c.mu.Lock()
// 		err := c.conn.WriteJSON(msg)
// 		c.mu.Unlock()

// 		if err != nil {
// 			return
// 		}
// 	}
// }

func NewWebsocketClient(conn *websocket.Conn) *WebsocketClient {
	return &WebsocketClient{
		Uuid:     uuid.NewString(),
		Username: "",
		conn:     conn,
		// TEST: for broadcasting
		send: make(chan WebsocketMessage),
	}
}
