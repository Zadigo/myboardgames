package models

import (
	"sync/atomic"

	"github.com/gorilla/websocket"
)

type Player struct{}

type WebsocketClient struct {
	Player
	conn     *websocket.Conn
	isActive atomic.Bool
}
