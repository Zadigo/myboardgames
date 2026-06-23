package models

import "github.com/redis/go-redis/v9"

type AppOptions struct {
	ServerApp   ServerAppInterface
	RedisClient *redis.Client
}

type GameAppOptions struct {
	AppOptions
	GameType string
}
