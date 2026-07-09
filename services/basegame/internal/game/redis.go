package game

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type GameRedis struct {
	ctx         context.Context
	redisClient *redis.Client
}

// Save persists the current state of the game application to Redis.
// It serializes the GameApp instance into JSON and stores it in a Redis
// hash with the game ID as the key.
func (gr *GameRedis) Save(app *GamesMiniServer) error {
	b, err := json.Marshal(app)
	if err != nil {
		return err
	}

	return gr.redisClient.SetNX(gr.ctx, "gameServer", b, 5*time.Minute).Err()
}

func NewGameRedis(ctx context.Context, redisClient *redis.Client) *GameRedis {
	return &GameRedis{
		ctx:         ctx,
		redisClient: redisClient,
	}
}
