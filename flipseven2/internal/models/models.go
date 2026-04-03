package models

type WebsocketMessage struct {
	Action       string        `json:"action"`
	Message      string        `json:"message,omitempty"`
	Deck         []Card        `json:"deck,omitempty"`
	TableId      string        `json:"tableId,omitempty"`
	PlayerId     string        `json:"playerId,omitempty"`
	Username     string        `json:"username,omitempty"`
	Players      []*Player     `json:"players,omitempty"`
	TableDetails *PlayersTable `json:"tableDetails,omitempty"`
	CardDetails  *Card         `json:"cardDetails,omitempty"`
}
