package redisclient

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(url string) *redis.Client {
	options, err := redis.ParseURL(url)

	if err != nil {
		panic(err)
	}

	client := redis.NewClient(options)

	// Test the connection
	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	return client
}
