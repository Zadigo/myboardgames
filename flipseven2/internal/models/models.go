package models

type WebsocketMessage struct {
	Action       string     `json:"action"`
	Message      string     `json:"message,omitempty"`
	Deck         []Card     `json:"deck,omitempty"`
	TableId      string     `json:"tableId,omitempty"`
	PlayerId     string     `json:"playerId,omitempty"`
	Username     string     `json:"username,omitempty"`
	Players      []*Player  `json:"players,omitempty"`
	TableDetails TableLayer `json:"tableDetails"`
	CardDetails  *Card      `json:"cardDetails,omitempty"`
}
