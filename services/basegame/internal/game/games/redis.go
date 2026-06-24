package games

import "github.com/redis/go-redis/v9"

type GameRedis struct {
	redisClient *redis.Client
}

