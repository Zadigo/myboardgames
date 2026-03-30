package handlers

import (
	"context"
	"net/http"

	"github.com/Zadigo/flipseven/internal/cards"
	"github.com/Zadigo/flipseven/internal/logic"
)

type WebsocketMessage struct {
	Action       string                  `json:"action"`
	Message      string                  `json:"message,omitempty"`
	Deck         []cards.Card            `json:"deck,omitempty"`
	TableId      string                  `json:"tableId,omitempty"`
	PlayerId     string                  `json:"playerId,omitempty"`
	Username     string                  `json:"username,omitempty"`
	Players      []logic.ConnectedPlayer `json:"players,omitempty"`
	TableDetails logic.PlayersTable      `json:"tableDetails,omitempty"`
}

type ContextHandlerFunc func(w http.ResponseWriter, r *http.Request, ctx context.Context)

type PostDataMessage struct {
	TableId string `json:"tableId"`
}
