package logic

import (
	"github.com/Zadigo/flipseven/cards"
	"github.com/gorilla/websocket"
)

type Player struct {
	// Player's username
	username string
	// Unique identifier for the player
	uuid string
	// The table on which the player is currently playing
	tableUuid string
	// The current number of cards the player has flipped
	numberOfCards int
	// Whether the player was stopped by the freeze card
	isFreezed bool
	// The number of rounds played by the player
	numberOfRounds int
	// The cards that the player has hands
	cards []cards.Card
}

type ConnectedPlayer struct {
	conn    websocket.Conn
	details Player
}

type PlayersTable struct {
	clients []ConnectedPlayer
}
