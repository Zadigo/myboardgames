package backend

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/redis/go-redis/v9"
)

type GameHistoryEntry struct {
	Action     string `json:"action"`
	PlayerUuid string `json:"player_uuid"`
}

type GameHistory struct {
	mu sync.Mutex
}

func (h *GameHistory) GetStreamName(gameUuid string) string {
	return fmt.Sprintf("oriflamme:history:%s", gameUuid)
}

func (h *GameHistory) PublishToStream(redisClient *redis.Client, gameUuid string, entry GameHistoryEntry) (string, error) {
	payload, err := json.Marshal(entry)
	if err != nil {
		return "", fmt.Errorf("Failed to marshal entry: %w", err)
	}

	cmd := redisClient.XAdd(context.Background(), &redis.XAddArgs{
		Stream: h.GetStreamName(gameUuid),
		Values: map[string]any{
			"payload": string(payload),
		},
	})

	if cmd.Err() != nil {
		return "", fmt.Errorf("Error publishing to stream: %w", cmd.Err())
	}

	return cmd.Val(), nil
}

func (h *GameHistory) GetStreamEntries(redisClient *redis.Client, gameUuid string) ([]GameHistoryEntry, error) {
	cmd := redisClient.XRange(context.Background(), h.GetStreamName(gameUuid), "-", "+")
	if cmd.Err() != nil {
		return nil, fmt.Errorf("Error fetching stream entries: %w", cmd.Err())
	}

	var entries []GameHistoryEntry
	for _, msg := range cmd.Val() {
		payload, ok := msg.Values["payload"].(string)
		if !ok {
			return nil, fmt.Errorf("Missing or invalid 'payload' field in stream entry %s", msg.ID)
		}

		var entry GameHistoryEntry
		if err := json.Unmarshal([]byte(payload), &entry); err != nil {
			return nil, fmt.Errorf("Failed to unmarshal entry %s: %w", msg.ID, err)
		}

		entries = append(entries, entry)
	}

	return entries, nil
}
