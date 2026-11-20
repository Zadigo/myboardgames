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

type InternalInstruction struct {
	Client      WebsocketClientInterface `json:"-"`
	Sender      string                   `json:"sender"`
	Action      string                   `json:"action"`
	StrMessage  string                   `json:"str_message"`
	JsonMessage map[string]any           `json:"json_message"`
	Table       TableInterface           `json:"-"`
}

type InternalErrorMessage struct {
	InternalInstruction
}

func (i *InternalInstruction) MatchesClient(client WebsocketClientInterface) bool {
	return i.Client == client
}
