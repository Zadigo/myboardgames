package models

const (
	StartGame    string = "start_game"
	MustIdentify string = "must_identify"
)

type BaseWebsocketMessage struct {
	Action string
}
