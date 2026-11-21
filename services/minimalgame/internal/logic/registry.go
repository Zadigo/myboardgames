package logic

import (
	"myboardgames/minimalgame/internal/models"
	"sync"
)

type GameRegistry struct {
	tables sync.Map
	// Tracks all active players on this current game
	// regardless of whether they are on a table or not
	players map[string]*models.Player
	mu      sync.Mutex
}

func (g *GameRegistry) AddPlayer(id string) *models.Player {
	g.mu.Lock()
	defer g.mu.Unlock()

	existingPlayer, ok := g.players[id]
	if ok {
		return existingPlayer
	} else {
		player := &models.Player{ID: id}
		g.players[id] = player
		return player
	}
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
		players: make(map[string]*models.Player),
	}
}
