package models

const (
	StartGame          string = "start_game"
	MustIdentify       string = "must_identify"
	ClientConnected    string = "client_connected"
	ClientDisconnected string = "client_disconnected"
)

type BaseWebsocketMessage struct {
	Action string
}

const (
	GameAppSender string = "GameAppSender"
	HttpAppSender string = "HttpAppSender"
)

type InternalMessage struct {
	Client      WebsocketClientInterface
	Sender      string
	StrMessage  string
	JsonMessage map[string]any
}

type InternalErrorMessage struct {
	InternalMessage
}
