package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type BaseRedis struct {
	ctx         context.Context
	redisClient *redis.Client
}

func (r *BaseRedis) GetClient() *redis.Client {
	return r.redisClient
}

func (r *BaseRedis) SetClient(redisClient *redis.Client) {
	r.redisClient = redisClient
}

func (r *BaseRedis) Close() error {
	return r.redisClient.Close()
}

func (r *BaseRedis) Ping() error {
	return r.redisClient.Ping(r.ctx).Err()
}
