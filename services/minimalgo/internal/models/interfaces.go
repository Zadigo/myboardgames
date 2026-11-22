package models

import (
	"context"

	"github.com/gorilla/websocket"
)

const (
	HelloPing          string = "hello_ping"
	StartGame          string = "start_game"
	MustIdentify       string = "must_identify"
	ClientConnected    string = "client_connected"
	ClientDisconnected string = "client_disconnected"
	CreateTable        string = "create_table"
	JoinTable          string = "join_table"
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
	GetState() bool
	GetOptions() *ServerAppOptions
}

type BaseWebsocketMessage struct {
	Action string `json:"action"`
}
