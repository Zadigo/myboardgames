package game

import (
	"github.com/Zadigo/basegame/internal/game/cards"
	"github.com/gorilla/websocket"
)

type Player struct {
	Conn  *websocket.Conn
	Cards []cards.BaseCardInterface
}
