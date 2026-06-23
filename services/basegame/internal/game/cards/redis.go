package cards

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

// CardsRedis is a struct that provides methods to interact
// with Redis for storing and retrieving game cards.
type CardsRedis struct {
	ctx         context.Context
	redisClient *redis.Client
	storageKey  string
}

// SaveCards saves the provided cards to Redis as a caching mechanism for the game.
func (g *CardsRedis) Save(cards []CardInterface) error {
	// Since the game provides a standard set of cards, we can clear the
	// existing cards in Redis before saving the new ones.
	if err := g.redisClient.Del(g.ctx, g.storageKey).Err(); err != nil {
		return err
	}

	values := make([]any, len(cards))

	for i, card := range cards {
		b, err := json.Marshal(card)
		if err != nil {
			return err
		}
		values[i] = b
	}

	return g.redisClient.RPush(g.ctx, g.storageKey, values...).Err()
}

// GetCards retrieves the cards from Redis and returns them as a slice of BaseCardInterface.
func (g *CardsRedis) Get() ([]CardInterface, error) {
	raws, err := g.redisClient.LRange(g.ctx, g.storageKey, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	cards := make([]CardInterface, 0, len(raws))
	for _, raw := range raws {
		var card CardInterface
		if err := json.Unmarshal([]byte(raw), &card); err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}

	return cards, nil
}

func NewCardsRedis(ctx context.Context, redisClient *redis.Client) *CardsRedis {
	return &CardsRedis{
		ctx:         ctx,
		storageKey:  "game_cards",
		redisClient: redisClient,
	}
}
