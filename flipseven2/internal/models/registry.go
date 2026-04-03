package models

import (
	"sync"

	"github.com/gorilla/websocket"
)

type BaseRegistry struct {
	tableLayers map[string]*TableLayer
	Mu          sync.RWMutex
	Clients     map[*websocket.Conn][]bool
}

func (r *BaseRegistry) Get(tableId string) (*TableLayer, bool) {
	r.Mu.RLock()
	defer r.Mu.RUnlock()
	table, ok := r.tableLayers[tableId]
	return table, ok
}

func (r *BaseRegistry) Set(tableLayer *TableLayer) {
	r.Mu.Lock()
	defer r.Mu.Unlock()
	r.tableLayers[tableLayer.Uuid] = tableLayer
}

func (r *BaseRegistry) GetOrCreate(tableID string, factory func() *TableLayer) *TableLayer {
	r.Mu.Lock()
	defer r.Mu.Unlock()

	if table, ok := r.tableLayers[tableID]; ok {
		return table
	}

	table := factory()
	r.tableLayers[tableID] = table
	return table
}

func (r *BaseRegistry) Delete(tableID string) {
	r.Mu.Lock()
	defer r.Mu.Unlock()
	delete(r.tableLayers, tableID)
}

func NewBaseRegistry() *BaseRegistry {
	return &BaseRegistry{
		tableLayers: make(map[string]*TableLayer),
	}
}

// Create a new global registry
func CreateNewBaseRegistry() *BaseRegistry {
	return &BaseRegistry{
		tableLayers: make(map[string]*TableLayer),
	}
}
