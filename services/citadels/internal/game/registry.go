package game

import (
	"sync"

	"github.com/Zadigo/citadels/internal/models"
)

type Games struct {
	games sync.Map
	mu    sync.Mutex
}

func (g *Games) Create() string {
	game := NewGame()
	g.Add(game)
	return game.GetID()
}

func (g *Games) Add(game models.GameInterface) {
	g.games.Store(game.GetID(), game)
}

func (g *Games) Get(id string) (models.GameInterface, bool) {
	v, ok := g.games.Load(id)
	if !ok {
		return nil, false
	}
	return v.(models.GameInterface), true
}

func (g *Games) Remove(id string) {
	g.games.Delete(id)
}

func NewGames() *Games {
	return &Games{
		games: sync.Map{},
	}
}
