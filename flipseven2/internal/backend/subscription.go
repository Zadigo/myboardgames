package backend

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

const (
	TableChannelPrefix = "table:"
	GlobalChannel      = "global"
)

// Create a new Redis subscription manager with the given Redis client and context
func CreateRedisSubcription(client *redis.Client, ctx context.Context) *RedisPubSub {
	return &RedisPubSub{client: client, ctx: ctx}
}

// TableChannel returns the Redis channel name for a specific table
func TableChannel(tableId string) string {
	return fmt.Sprintf("%s%s", TableChannelPrefix, tableId)
}

// Publish sends a message to a table-specific Redis channel
func (r *RedisPubSub) Publish(tableId string, msg Message) error {
	channel := TableChannel(tableId)

	data, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	if err := r.client.Publish(r.ctx, channel, data).Err(); err != nil {
		return fmt.Errorf("failed to publish to channel %s: %w", channel, err)
	}

	log.Printf("📢 Published to channel %s: %s", channel, msg.Action)
	return nil
}

// PublishGlobal sends a message to the global channel
func (redisPubSub *RedisPubSub) PublishGlobal(msg Message) error {
	data, err := json.Marshal(msg)

	if err != nil {
		return fmt.Errorf("Failed to marshal message: %w", err)
	}

	return redisPubSub.client.Publish(redisPubSub.ctx, GlobalChannel, data).Err()
}

// Subscribe listens on a table channel and calls handler for each message.
// Blocks until ctx is cancelled — run in a goroutine.
func (redisPubSub *RedisPubSub) Subscribe(tableId string, handler func(Message)) error {
	channel := TableChannel(tableId)
	sub := redisPubSub.client.Subscribe(redisPubSub.ctx, channel)

	defer func() {
		if err := sub.Close(); err != nil {
			log.Printf("⚠️ Error closing subscription on %s: %v", channel, err)
		}
	}()

	log.Printf("👂 Subscribed to channel: %s", channel)

	ch := sub.Channel()

	for {
		select {
		case <-redisPubSub.ctx.Done():
			log.Printf("🛑 Subscription cancelled for channel: %s", channel)
			return nil

		case redisMsg, ok := <-ch:
			if !ok {
				return fmt.Errorf("channel %s closed unexpectedly", channel)
			}

			var msg Message
			if err := json.Unmarshal([]byte(redisMsg.Payload), &msg); err != nil {
				log.Printf("⚠️ Failed to unmarshal message on %s: %v", channel, err)
				continue
			}

			handler(msg)
		}
	}
}
