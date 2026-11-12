package models

type CharacterSelection struct {
	Names  []string `json:"names"`
	Random bool     `json:"random"`
}

type WebsocketMessage struct {
	CharacterSelection `json:"character_selection"`
	Action             string         `json:"action"`
	Data               map[string]any `json:"data"`
	Message            string         `json:"message"`
}
