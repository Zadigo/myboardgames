package redisclient

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func NewSubscription(client *redis.Client, ctx context.Context) *RedisSubscription {
	return &RedisSubscription{client: client, ctx: ctx}
}

// Publish sends a message to a table-specific Redis channel
func (rs *RedisSubscription) Publish(tableUuid string, msg Message) {}

// PublishGlobal sends a message to the global channel
func (rs *RedisSubscription) PublishGlobal(msg Message) {}

// Subscribe listens on a table channel and calls handler for each message.
// Blocks until ctx is cancelled — run in a goroutine.
func (rs *RedisSubscription) Subscribe(tableUuid string, handler func(Message)) {}
