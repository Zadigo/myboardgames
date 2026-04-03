package models

import "sync"

type BaseRegistry struct {
	tableLayers map[string]*TableLayer
	mu          sync.RWMutex
}

func (r *BaseRegistry) Get(tableId string) (*TableLayer, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	table, ok := r.tableLayers[tableId]
	return table, ok
}

func (r *BaseRegistry) Set(tableLayer *TableLayer) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.tableLayers[tableLayer.Uuid] = tableLayer
}

func (r *BaseRegistry) GetOrCreate(tableID string, factory func() *TableLayer) *TableLayer {
	r.mu.Lock()
	
	defer r.mu.Unlock()

	if table, ok := r.tableLayers[tableID]; ok {
		return table
	}
	
	table := factory()
	r.tableLayers[tableID] = table
	return table
}

func (r *BaseRegistry) Delete(tableID string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.tableLayers, tableID)
}

func NewBaseRegistry() *BaseRegistry {
	return &BaseRegistry{
		tableLayers: make(map[string]*TableLayer),
	}
}
