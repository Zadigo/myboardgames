package redis

import (
	"context"
	"encoding/json"

	"github.com/Zadigo/basegame/internal/models"
	"github.com/redis/go-redis/v9"
)

type GameRedis struct {
	BaseRedis
	storageKey string
}

// SaveCards saves the provided cards to Redis as a caching mechanism for the game.
func (g *GameRedis) SaveCards(cards []models.BaseCardInterface) error {
	byteCards, err := json.Marshal(cards)
	if err != nil {
		return err
	}

	err = g.redisClient.LPush(g.ctx, g.storageKey, byteCards).Err()
	if err != nil {
		return err
	}
	return nil
}

// GetCards retrieves the cards from Redis and returns them as a slice of BaseCardInterface.
func (g *GameRedis) GetCards() ([]models.BaseCardInterface, error) {
	cmd := g.redisClient.LRange(g.ctx, g.storageKey, 0, -1)
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}

	cachedCards := []models.BaseCardInterface{}
	for _, cmdVal := range cmd.Val() {
		var card models.BaseCardInterface

		err := json.Unmarshal([]byte(cmdVal), &card)
		if err != nil {
			return nil, err
		}

		cachedCards = append(cachedCards, card)
	}

	return cachedCards, nil
}

func NewGameRedis(ctx context.Context, redisClient *redis.Client) *GameRedis {
	return &GameRedis{
		storageKey: "game_cards", // Example storage key for cards in Redis
		BaseRedis: BaseRedis{
			ctx:         ctx,
			redisClient: redisClient,
		},
	}
}
