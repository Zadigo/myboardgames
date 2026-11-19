package models

import (
	"context"

	"github.com/gorilla/websocket"
)

type AppInterface interface {
	Start(ctx context.Context)
}

type WebsocketClientInterface interface {
	IsActive() bool
	SetConn(conn *websocket.Conn)
}

type ServerAppInterface interface {
	AppInterface
	UpdateClients(clients []WebsocketClientInterface)
}
