package models

const (
	MUST_IDENTIFY = "must_identify"
	START_GAME        = "start_game"
	END_GAME          = "end_game"
	START_TURN        = "start_turn"
	END_TURN          = "end_turn"
	JOIN              = "join"
	OBSERVE           = "observe"
	FLIP_CARD         = "flip_card"
	GIVE_CARD         = "give_card"
)

type WebsocketMessage struct {
	Action string `json:"action"`
}
