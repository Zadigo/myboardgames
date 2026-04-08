package redisclient

import "github.com/gorilla/websocket"

func (b *Broadcaster) AddClient(conn *websocket.Conn) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.clients[conn] = true
}

func (br *BroadcastingRegistry) GetOrCreate(tableUuid string) *Broadcaster {
	br.mu.Lock()
	defer br.mu.Unlock()

	// Check if a broadcaster already exists for this table
	b, exists := br.broadcasters[tableUuid]
	if !exists {
		// If not, create a new one
		b = NewBroadcaster(tableUuid, br.Subscription, br.ctx)
		br.broadcasters[tableUuid] = b
	}

	return b
}
