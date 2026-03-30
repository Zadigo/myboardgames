package logic

import "github.com/gorilla/websocket"

type Card struct {
	Value int
	Owner string
}

type Player struct {
	Username   string
	Uuid       string
	Connection *websocket.Conn
}

func (player *Player) New(connection *websocket.Conn, username string) *Player {
	return &Player{
		Username:   username,
		Connection: connection,
	}
}

func (player *Player) AttributeCard(card *Card) {
	card.Owner = player.Uuid
}

type Table struct {
	TableId string
	Clients map[string]*Player
}

func (table Table) StartGame() {
	// Do something
}

func (table Table) EndGame() {
	// Do something
}

func (table Table) FlipCard() {
	// Do something
}

func (table Table) FreezePlayer(player string) {
	// Do something
}

func (table Table) NextRound() {
	// Do something
}

func (table Table) GetTableId() string {
	// Do something
	return table.TableId
}

func (table Table) GetPlayer(playerId string) *Player {
	return table.Clients[playerId]
}

func (table Table) HasPlayer(playerId string) bool {
	_, ok := table.Clients[playerId]
	return ok
}

type TableInterface interface {
	StartGame()
	EndGame()
	FlipCard()
	FreezePlayer(player string)
	NextRound()
	GetTableId() string
}

type TableLayer struct {
	Layer TableInterface
}

func NewTableLayer(tableId string) *TableLayer {
	return &TableLayer{
		Layer: Table{
			TableId: tableId,
		},
	}
}
