package backend

import (
	"context"
	"log"

	"github.com/Zadigo/flipseven/internal/logic"
	"github.com/gorilla/websocket"
)

func NewBroadcaster(tableId string, ps *RedisPubSub, parentCtx context.Context) *Broadcaster {
	ctx, cancel := context.WithCancel(parentCtx)

	b := &Broadcaster{
		tableId: tableId,
		pubsub:  ps,
		clients: make(map[*websocket.Conn]bool),
		cancel:  cancel,
	}

	// Start listening on the Redis channel for this table
	go func() {
		if err := ps.Subscribe(tableId, b.fanOut); err != nil {
			log.Printf("⚠️ Broadcaster for table %s exited: %v", tableId, err)
		}
	}()

	_ = ctx
	return b
}

// AddClient registers a websocket connection for this table
func (b *Broadcaster) AddClient(conn *websocket.Conn) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.clients[conn] = true
	log.Printf("➕ Client added to broadcaster for table %s (total: %d)", b.tableId, len(b.clients))
}

// RemoveClient unregisters a websocket connection
func (b *Broadcaster) RemoveClient(conn *websocket.Conn) {
	b.mu.Lock()
	defer b.mu.Unlock()

	delete(b.clients, conn)

	log.Printf("➖ Client removed from broadcaster for table %s (total: %d)", b.tableId, len(b.clients))

	if len(b.clients) == 0 {
		log.Printf("🗑️ No clients left for table %s, stopping broadcaster", b.tableId)
		b.cancel()
	}
}

// Publish sends a message via Redis so ALL server instances receive it
func (b *Broadcaster) Publish(action string, payload any) error {
	return b.pubsub.Publish(b.tableId, Message{
		Action:  action,
		TableId: b.tableId,
		Payload: payload,
	})
}

// fanOut is called by the Redis subscriber and writes to all local websocket clients
func (b *Broadcaster) fanOut(msg Message) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for conn := range b.clients {
		if err := conn.WriteJSON(toWsMessage(msg)); err != nil {
			log.Printf("⚠️ Failed to write to client: %v — removing", err)
			// Schedule removal outside the read lock to avoid deadlock
			go b.RemoveClient(conn)
		}
	}
}

// toWsMessage converts a pub/sub Message to the format expected by the frontend.
// Adjust the target type to match handlers.WebsocketMessage in your project.
func toWsMessage(msg Message) map[string]any {
	out := map[string]any{
		"action":  msg.Action,
		"tableId": msg.TableId,
	}

	if msg.Payload != nil {
		out["payload"] = msg.Payload
	}

	return out
}

func NewRegistry(ps *RedisPubSub, ctx context.Context) *BroadcasterRegistry {
	return &BroadcasterRegistry{
		broadcasters: make(map[string]*Broadcaster),
		pubsub:       ps,
		ctx:          ctx,
	}
}

// GetOrCreate returns an existing broadcaster or creates one for the given tableId
func (r *BroadcasterRegistry) GetOrCreate(tableId string) *Broadcaster {
	r.mu.Lock()
	defer r.mu.Unlock()

	if b, ok := r.broadcasters[tableId]; ok {
		return b
	}

	b := NewBroadcaster(tableId, r.pubsub, r.ctx)
	r.broadcasters[tableId] = b
	log.Printf("🆕 Created broadcaster for table %s", tableId)
	return b
}

// Get returns an existing broadcaster or nil
func (r *BroadcasterRegistry) Get(tableId string) (*Broadcaster, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	b, ok := r.broadcasters[tableId]
	return b, ok
}

// Remove deletes a broadcaster from the registry
func (r *BroadcasterRegistry) Remove(tableId string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.broadcasters, tableId)
}

// Ensure internal.ConnectedPlayer is importable — this avoids a circular import
// by keeping pubsub free of handler-layer types.
var _ = logic.ConnectedPlayer{}
