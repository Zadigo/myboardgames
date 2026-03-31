package handlers

import (
	"context"
	"net/http"

	"github.com/Zadigo/flipseven/internal/logic"
)

type WebsocketMessage struct {
	Action       string              `json:"action"`
	Message      string              `json:"message,omitempty"`
	Deck         []logic.Card        `json:"deck,omitempty"`
	TableId      string              `json:"tableId,omitempty"`
	PlayerId     string              `json:"playerId,omitempty"`
	Username     string              `json:"username,omitempty"`
	Players      []logic.PlayerLayer `json:"players,omitempty"`
	TableDetails *logic.PlayersTable `json:"tableDetails,omitempty"`
	CardDetails  *logic.Card         `json:"cardDetails,omitempty"`
}

type ContextHandlerFunc func(w http.ResponseWriter, r *http.Request, ctx context.Context)

type PostDataMessage struct {
	TableId string `json:"tableId"`
}
