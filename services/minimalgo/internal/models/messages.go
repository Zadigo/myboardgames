package models

const (
	HelloPing          string = "hello_ping"
	StartGame          string = "start_game"
	MustIdentify       string = "must_identify"
	ClientConnected    string = "client_connected"
	ClientDisconnected string = "client_disconnected"
	CreateTable        string = "create_table"
	JoinTable          string = "join_table"
)

type BaseWebsocketMessage struct {
	Action string `json:"action"`
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
	Table       TableInterface
}

type InternalErrorMessage struct {
	InternalMessage
}

func (i *InternalMessage) MatchesClient(client WebsocketClientInterface) bool {
	return i.Client == client
}
