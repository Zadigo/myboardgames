package http

import (
	"myboardgames/minimalgo/internal/models"
	"uuid"

	"github.com/gorilla/websocket"
)

func NewWebsocketClient(conn *websocket.Conn) models.WebsocketClientInterface {
	client := &models.WebsocketClient{ID: uuid.New().String()}
	client.SetConn(conn)
	return client
}
