package game

type BaseGame struct {
	table   map[string]*Table
	players map[string]*Player
}

func (g *BaseGame) Start() error {
	return nil
}

func (g *BaseGame) CreateTable() error {
	return nil
}

func (g *BaseGame) CreateCards() error {
	return nil
}

func (g *BaseGame) Prepare() error {
	return nil
}

type StandardGame struct {
	BaseGame
}

func NewStandardGame() *StandardGame {
	return &StandardGame{}
}

type ExtensionGame struct {
	BaseGame
}

func NewExtensionGame() *ExtensionGame {
	return &ExtensionGame{}
}

type GameInterface interface {
	Start() error
	CreateTable() error
	CreateCards() error
	Prepare() error
}

func CreateGame(gameType string) GameInterface {
	switch gameType {
	case STANDARD:
		return NewStandardGame()
	case EXTENSION:
		return NewExtensionGame()
	default:
		return nil
	}
}
