package models

import (
	"sync"

	"github.com/gorilla/websocket"
)

type WebsocketClient struct {
	Conn *websocket.Conn `json:"-"`
	Mu   *sync.RWMutex   `json:"-"`
}
