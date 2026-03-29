package internal

import (
	"context"
	"net/http"

	"github.com/Zadigo/flipseven/internal/cards"
)

type WebsocketMessage struct {
	Action   string       `json:"action"`
	Message  string       `json:"message"`
	Deck     []cards.Card `json:"deck"`
	TableId  string       `json:"table_id"`
	PlayerId string       `json:"player_id"`
}

type ContextHandlerFunc func(w http.ResponseWriter, r *http.Request, ctx context.Context)

type PostDataMessage struct {
	TableId string `json:"tableId"`
}
