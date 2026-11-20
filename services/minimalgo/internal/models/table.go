package models

type TableInterface interface {
}

type Table struct {
	ID       string
	_players []WebsocketClientInterface
}
