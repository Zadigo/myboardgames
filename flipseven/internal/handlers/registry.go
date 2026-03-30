package handlers

import (
	"sync"

	"github.com/Zadigo/flipseven/internal/logic"
)

// BaseRegistry is a thread-safe in-memory registry for managing game tables.
// It provides methods to get, set, and delete tables, as well as a factory method
// to create tables if they do not exist.
type BaseRegistry struct {
	mu     sync.RWMutex
	tables map[string]*logic.TableLayer
}

// Get retrieves a table by its ID. It returns the table and a boolean indicating
// whether the table exists in the registry.
func (r *BaseRegistry) Get(tableId string) (*logic.TableLayer, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	table, exists := r.tables[tableId]
	return table, exists
}

// Set adds or updates a table in the registry with the given ID.
func (r *BaseRegistry) Set(tableID string, table *logic.TableLayer) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.tables[tableID] = table
}

// GetOrCreate retrieves a table by its ID. If the table does not exist, it uses the provided
// factory function to create a new table, adds it to the registry, and returns it.
func (r *BaseRegistry) GetOrCreate(tableID string, factory func() *logic.TableLayer) *logic.TableLayer {
	r.mu.Lock()
	defer r.mu.Unlock()

	tableLayer, exists := r.tables[tableID]
	if !exists {
		tableLayer = factory()
		r.tables[tableLayer.Layer.GetTable().TableId] = tableLayer
	}

	return tableLayer
}

// Delete removes a table from the registry by its ID.
func (r *BaseRegistry) Delete(tableID string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.tables, tableID)
}

// NewBaseRegistry creates and returns a new instance of BaseRegistry with
// an initialized tables map.
func NewBaseRegistry() *BaseRegistry {
	return &BaseRegistry{
		tables: make(map[string]*logic.TableLayer),
	}
}
