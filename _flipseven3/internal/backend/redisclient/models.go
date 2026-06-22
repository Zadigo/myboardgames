package redisclient

import (
	"context"

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
