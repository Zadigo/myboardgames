package internal

import "github.com/Zadigo/flipseven/internal/cards"

type WebsocketMessage struct {
	Action   string       `json:"action"`
	Message  string       `json:"message"`
	Deck     []cards.Card `json:"deck"`
	TableId  string       `json:"table_id"`
	PlayerId string       `json:"player_id"`
}
