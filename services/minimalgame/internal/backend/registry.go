package backend

import (
	"myboardgames/minimalgame/internal/models"
	"sync"
)

type Table struct {
	ID      string
	players map[string]*Player
}

func (t *Table) AddPlayer(id string) error {
	t.players[id] = &Player{ID: id}
	return nil
}

func (t *Table) RemovePlayer(id string) error {
	delete(t.players, id)
	return nil
}

type Player struct {
	ID string
}

type GameRegistry struct {
	tables  map[string]models.TableInterface
	players map[string]*Player
	mu      sync.Mutex
}

func (g *GameRegistry) AddPlayer(id string) error {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.players[id] = &Player{ID: id}
	return nil
}

func (g *GameRegistry) RemovePlayer(id string) error {
	g.mu.Lock()
	defer g.mu.Unlock()
	delete(g.players, id)
	return nil
}

func NewGameRegistry() models.GameRegistryInterface {
	return &GameRegistry{
		tables:  make(map[string]models.TableInterface),
		players: make(map[string]*Player),
	}
}
