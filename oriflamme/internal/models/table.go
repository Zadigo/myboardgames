package models

type PlayingTable struct {
	Players []Player
}

type LayerInterface interface {
	AddCardLeft(card *Card)
	AddCardRight(card *Card)
	RemoveCardAtPosition(index int) *Card
	StackCard(index int) *Card
}

type TableLayer struct {
	Layer               LayerInterface
	CurrentRound        int
	CurrentPlayer       *Player
	MaxPlayers          int
	MaxRounds           int
	TotalCardsPerPlayer int
	MaxCardsPerPlayer   int
}

func (table *PlayingTable) AddPlayer(player *Player) {}

func (table *PlayingTable) StartGame() {}

func (table *PlayingTable) EndGame() {}

// Create the initial influnce queue with 5 cards of
// each color for each player
func CreateTableLayer() *TableLayer {
	return &TableLayer{
		CurrentRound:        0,
		CurrentPlayer:       nil,
		MaxPlayers:          5,
		MaxRounds:           7,
		TotalCardsPerPlayer: 10,
		MaxCardsPerPlayer:   7,
		Layer: &InfluenceQueue{
			Queue: []*Card{},
		},
	}
}
