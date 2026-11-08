package models

type GameInterface interface {
	GetPlayers() []PlayerInterface
}

type BaseGame struct {
	Players []PlayerInterface
}

func (g *BaseGame) GetPlayers() []PlayerInterface {
	return g.Players
}
