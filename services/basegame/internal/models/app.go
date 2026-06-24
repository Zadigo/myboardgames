package models

import "github.com/redis/go-redis/v9"

type AppOptions struct {
	ServerApp ServerAppInterface
	// Deprecated: Get the Redis client from the AppInterface
	// using GetRedisClient instead of passing it here.
	RedisClient *redis.Client
}

type GameAppOptions struct {
	AppOptions
	GameType string
	Debug    bool
}
