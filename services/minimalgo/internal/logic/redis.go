package logic

import "github.com/redis/go-redis/v9"

type TableRedis struct {
	client *redis.Client
}
