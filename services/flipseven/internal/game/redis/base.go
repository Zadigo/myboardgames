package redis

import "github.com/redis/go-redis/v9"

type GameRedis struct {
	redisClient *redis.Client
}

func NewGameRedis(redisClient *redis.Client) *GameRedis {
	return &GameRedis{
		redisClient: redisClient,
	}
}
