package backend

import (
	"context"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

// ServerBackendConfig represents the configuration for a specific backend in the YAML file
type ServerBackendConfig struct {
	Url string `json:"url" yaml:"url"`
}

type DatabaeClientConfig struct {
	Client string `json:"client" yaml:"client"`
}

// ServerBackendsConfig represents the backend configuration in the YAML file
type ServerBackendsConfig struct {
	Database *DatabaeClientConfig `json:"database" yaml:"database"`
	Postgres *ServerBackendConfig `json:"postgres" yaml:"postgres"`
	Redis    *ServerBackendConfig `json:"redis" yaml:"redis"`
	RabbitMQ *ServerBackendConfig `json:"rabbitmq" yaml:"rabbitmq"`
}

type ServerBaseConfig struct {
	Backends *ServerBackendsConfig `json:"backends" yaml:"backends"`
}

// YamlConfig represents the structure of the configuration YAML file
type ServerConfig struct {
	Config ServerBaseConfig `json:"config" yaml:"config"`
}

type Message struct {
	Action  string `json:"action"`
	TableId string `json:"tableId,omitempty"`
	Payload any    `json:"payload,omitempty"`
}

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
