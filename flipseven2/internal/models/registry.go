package models

import (
	"sync"

	"github.com/gorilla/websocket"
)

type BaseRegistry struct {
	tableLayers map[string]*PlayersTable
	Mu          sync.RWMutex
	Clients     map[*websocket.Conn][]bool
}

// Get retrieves a PlayersTable from the registry based on the
// provided tableId. It returns the PlayersTable and a boolean indicating
// whether the table was found in the registry.
func (r *BaseRegistry) Get(tableId string) (*PlayersTable, bool) {
	r.Mu.RLock()
	defer r.Mu.RUnlock()
	table, ok := r.tableLayers[tableId]
	return table, ok
}

// Set adds or updates a PlayersTable in the registry. It takes
// a pointer to a PlayersTable as an argument and stores it in
// the registry's tableLayers map, using the table's UUID as the key.
// The method locks the registry for writing to ensure thread
// safety while modifying the map.
func (r *BaseRegistry) Set(tableLayer *PlayersTable) {
	r.Mu.Lock()
	defer r.Mu.Unlock()
	r.tableLayers[tableLayer.Layer.GetUuid()] = tableLayer
}

// GetOrCreate retrieves a PlayersTable from the registry based on the
// provided tableID. If the PlayersTable does not exist, it creates a new
// one using the provided factory function, adds it to the registry, and
// returns it. The method locks the registry for writing to ensure thread
// safety while modifying the map.
func (r *BaseRegistry) GetOrCreate(tableID string, factory func() *PlayersTable) *PlayersTable {
	r.Mu.Lock()
	defer r.Mu.Unlock()

	if table, ok := r.tableLayers[tableID]; ok {
		return table
	}

	table := factory()
	r.tableLayers[table.Layer.GetUuid()] = table
	return table
}

// Delete removes a PlayersTable from the registry based on the provided tableID.
// It locks the registry for writing to ensure thread safety while
// modifying the map, and then deletes the entry corresponding to the
// provided tableID from the tableLayers map.
func (r *BaseRegistry) Delete(tableID string) {
	r.Mu.Lock()
	defer r.Mu.Unlock()
	delete(r.tableLayers, tableID)
}

// Create a new global registry
func CreateNewBaseRegistry() *BaseRegistry {
	return &BaseRegistry{
		tableLayers: make(map[string]*PlayersTable),
	}
}
