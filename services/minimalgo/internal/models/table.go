package models

// GameInterface represents a game logic that a given
// table should host. The interface implements basic
// game logic that most games need.
type GameInterface interface {
	StartGame() error
	EndGame() error
	ResetGame() error
}

type TableInterface interface {
}

// Table represents a game table with players.
type Table struct {
	ID       string
	game     GameInterface
	_players []WebsocketClientInterface
}

func (t *Table) SetGame(g GameInterface) {
	t.game = g
}
