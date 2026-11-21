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

type BaseInstruction struct {
	Sender      string                   `json:"sender"` // The gorouting sending the instruction
	Action      string                   `json:"action"`
	StrMessage  string                   `json:"str_message"`
	JsonMessage map[string]any           `json:"json_message"`
	Table       TableInterface           `json:"table"` // The table to which the instruction is sent to
	TableID     string                   `json:"table_id"`
	Client      WebsocketClientInterface `json:"client"`
}

// Represents an instruction to send to the game
// server specifically. This is used internally to
// transform a websocket action to game logical instruction
type InternalInstruction struct {
	BaseInstruction
	Client       WebsocketClientInterface `json:"client"` // Client sending the instruction (if applicable)
	TableOptions *TableOptions            `json:"-"`      // Options to send to a table when creating or updating it
}

type InternalInstructionResponse struct {
	BaseInstruction
}

type InternalInstructionError struct {
	InternalInstruction
}

func (i *InternalInstruction) MatchesClient(client WebsocketClientInterface) bool {
	return i.Client == client
}
