package models

import (
	"context"

	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
)

const (
	PingRpc            string = "ping_rpc" // Tries to connect to an RPC server
	StartGame          string = "start_game"
	MustIdentify       string = "must_identify"
	ClientConnected    string = "client_connected"
	ClientDisconnected string = "client_disconnected"
	CreateTable        string = "create_table"
	JoinTable          string = "join_table"
	JoinWebsite        string = "join_website"
)

type AppInterface interface {
	Start(ctx context.Context) error
}

type WebsocketClientInterface interface {
	GetID() string
	IsActive() bool
	SetConn(conn *websocket.Conn)
	GetConn() *websocket.Conn
	WriteJSON(msg BaseWebsocketMessage) error
}

type ServerAppInterface interface {
	AppInterface
	GetRouter() *Router
	GetGrpcConnections() map[string]*grpc.ClientConn
	GetGrpcConnection(id string) (*grpc.ClientConn, error)
	UpdateClients(clients []WebsocketClientInterface)
	GetState() bool
	GetOptions() *ServerAppOptions
}

type BaseWebsocketMessage struct {
	Action      string         `json:"action"`
	Rpc         string         `json:"rpc"`
	RpcState    string         `json:"rpc_state"`
	ClientID    string         `json:"client_id"`
	StrMessage  string         `json:"message"`
	JsonMessage map[string]any `json:"json_message"`
	Routing     *Router        `json:"router"`
}
