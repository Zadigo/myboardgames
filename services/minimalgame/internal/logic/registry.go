package logic

import (
	"myboardgames/minimalgame/internal/models"
	"sync"
)

type Player struct {
	ID string
}

type GameRegistry struct {
	tables  sync.Map
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

func (g *GameRegistry) AddTable(table models.TableInterface) error {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.tables.Store(table.(interface{ GetID() string }).GetID(), table)
	return nil
}

func (g *GameRegistry) RemoveTable(table models.TableInterface) error {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.tables.Delete(table.(interface{ GetID() string }).GetID())
	return nil
}

func (g *GameRegistry) GetTable(id string) (models.TableInterface, error) {
	t, ok := g.tables.Load(id)
	if !ok {
		return nil, nil
	}
	return t.(models.TableInterface), nil
}

func NewGameRegistry() models.GameRegistryInterface {
	return &GameRegistry{
		tables:  sync.Map{},
		players: make(map[string]*Player),
	}
}
