package redis

import (
	"context"

	"github.com/Zadigo/basegame/internal/game/cards"
	"github.com/redis/go-redis/v9"
)

type GameRedis struct {
	BaseRedis
}

// SaveCards saves the provided cards to Redis as a caching mechanism for the game.
func (g *GameRedis) SaveCards(cards []cards.BaseCardInterface) error {
	return nil
}

// GetCards retrieves the cards from Redis and returns them as a slice of BaseCardInterface.
func (g *GameRedis) GetCards() ([]cards.BaseCardInterface, error) {
	return nil, nil
}

// AttributeCard associates a card with a player in Redis, allowing for tracking of which player has which card.
func (g *GameRedis) AttributeCard(card cards.BaseCardInterface, player any) error {
	return nil
}

func NewGameRedis(ctx context.Context, redisClient *redis.Client) *GameRedis {
	return &GameRedis{
		BaseRedis: BaseRedis{
			ctx:         ctx,
			redisClient: redisClient,
		},
	}
}
