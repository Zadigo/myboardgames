package models

// Base websocket message structure for
// communication between client and server
type BaseWebsocketMessage struct {
	Action  string           `json:"action"`
	Message string           `json:"message"`
	Client  *WebsocketClient `json:"-"`
}

func (c *BaseWebsocketMessage) SendJsonMessage() error {
	c.Client.Mu.Lock()
	defer c.Client.Mu.Unlock()
	return c.Client.Conn.WriteJSON(c)
}

func (c *BaseWebsocketMessage) ReadJsonMessage() error {
	c.Client.Mu.RLock()
	defer c.Client.Mu.RUnlock()
	return c.Client.Conn.ReadJSON(c)
}

type AuthenticationMessage struct {
	BaseWebsocketMessage
	Username string `json:"username"`
}

type WebsocketMessageInterface interface {
	AuthenticationMessage
}

type WebsocketMessage[T WebsocketMessageInterface] = T
