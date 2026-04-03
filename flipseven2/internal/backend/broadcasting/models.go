package broadcasting

import (
	"context"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

type RedisPubSub struct {
	client *redis.Client
	ctx    context.Context
}

// Broadcaster manages subscriptions and fan-out to websocket clients
// for a single game table.
type Broadcaster struct {
	tableId string
	pubsub  *RedisPubSub
	mu      sync.RWMutex
	clients map[*websocket.Conn]bool
	cancel  context.CancelFunc
}

// BroadcasterRegistry keeps one Broadcaster per table, safe for concurrent use.
type BroadcasterRegistry struct {
	mu           sync.RWMutex
	broadcasters map[string]*Broadcaster
	pubsub       *RedisPubSub
	ctx          context.Context
}

type Message struct {
	Action  string `json:"action"`
	TableId string `json:"tableId,omitempty"`
	Payload any    `json:"payload,omitempty"`
}
