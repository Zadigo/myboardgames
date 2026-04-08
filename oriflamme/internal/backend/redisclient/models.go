package redisclient

import (
	"context"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

type Message struct {
	Action    string `json:"action"`
	TableUuid string `json:"tableUuid,omitempty"`
	Payload   any    `json:"payload,omitempty"`
}

// Provides publish/subscribe functionality
// using Redis channels.
type RedisSubscription struct {
	client *redis.Client
	ctx    context.Context
}

// The table broadcaster manages subscriptions
// and fan-out to websocket clients for a single
// game table.
type Broadcaster struct {
	tableUuid    string
	subscription *RedisSubscription
	clients      map[*websocket.Conn]bool
	mu           sync.RWMutex
	cancel       context.CancelFunc
}

// The registry references all broadcasters, one per table,
// and is safe for concurrent use.
type BroadcastingRegistry struct {
	mu           sync.RWMutex
	broadcasters map[string]*Broadcaster
	Subscription *RedisSubscription
	cancel       context.CancelFunc
	ctx          context.Context
}

func NewBroadcastingRegistry(tableUuid string, subscription *RedisSubscription, parentCtx context.Context) *BroadcastingRegistry {
	return &BroadcastingRegistry{
		broadcasters: make(map[string]*Broadcaster),
		Subscription: subscription,
		ctx:          parentCtx,
	}
}

func NewBroadcaster(tableUuid string, subscription *RedisSubscription, parentCtx context.Context) *Broadcaster {
	_, cancel := context.WithCancel(parentCtx)

	b := &Broadcaster{
		tableUuid:    tableUuid,
		subscription: subscription,
		clients:      make(map[*websocket.Conn]bool),
		cancel:       cancel,
	}

	go func() {

	}()

	return b
}
