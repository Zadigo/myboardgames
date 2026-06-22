package models

import (
	"github.com/gorilla/websocket"
)

const (
	// A player requests to identify themselves
	MUST_IDENTIFY = "must_identify"
	// A game was started
	START_GAME = "start_game"
	// A player ends the game
	END_GAME = "end_game"
	// A player starts their turn
	START_TURN = "start_turn"
	// A player ends their turn
	END_TURN = "end_turn"
	// A player wants to join a game
	JOIN = "join"
	// A player wants to observe a game
	OBSERVE = "observe"
	// A player flips a card
	FLIP_CARD = "flip_card"
	// The value of the card is resolved
	RESOLVE_CARD = "resolve_card"
	// A player gives a card to another player
	GIVE_CARD = "give_card"
)

// WebsocketMessage represents a message sent
// over the websocket connection.
type BaseWebsocketMessage struct {
	Action string `json:"action"`
}

type WebsocketMessage struct {
	BaseWebsocketMessage
}

type WebscoketClientInterface interface {
	GetUuid() string
	SetConn(conn *websocket.Conn)
	SendJsonMessage(message WebsocketMessage) error
	ReceiveJsonMessage() (WebsocketMessage, error)
}
