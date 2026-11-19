package http

import (
	"myboardgames/minimalgo/internal/models"

	"github.com/gorilla/websocket"
)

func NewWebsocketClient(conn *websocket.Conn) models.WebsocketClientInterface {
	client := &models.WebsocketClient{}
	client.SetConn(conn)
	return client
}
