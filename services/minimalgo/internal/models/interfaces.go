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
	CreateTable(ctx context.Context, options *TableOptions) TableInterface
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
	SendInstruction(ctx context.Context, msg InternalInstruction) error
	SendInstructionResponse(ctx context.Context, msg InternalInstructionResponse) error
	GetInstructions() <-chan InternalInstruction
	GetState() bool
	GetOptions() *ServerAppOptions
}
