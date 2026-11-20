package models

import (
	"context"

	"github.com/gorilla/websocket"
)

type AppInterface interface {
	Start(ctx context.Context) error
}

type GameServerInterface interface {
	AppInterface
	CreateTable(game GameInterface) (string, TableInterface)
	GetTables() []TableInterface
}

type WebsocketClientInterface interface {
	IsActive() bool
	SetConn(conn *websocket.Conn)
	GetConn() *websocket.Conn
}

type ServerAppInterface interface {
	AppInterface
	UpdateClients(clients []WebsocketClientInterface)
	WriteInternalMessage(ctx context.Context, msg InternalMessage) error
	GetInternalChannel() <-chan InternalMessage
	GetState() bool
}
