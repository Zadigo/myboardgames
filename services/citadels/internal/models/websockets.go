package models

type CharacterSelection struct {
	Names  []string `json:"names"`
	Random bool     `json:"random"`
}

type WebsocketMessage struct {
	CharacterSelection `json:"character_selection"`
	Action             string `json:"action"`
}
