package backend

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func CreateRedisConnection() (*redis.Client, error) {
	options, err := redis.ParseURL("redis://localhost:6379")
	if err != nil {
		return nil, fmt.Errorf("Failed to resolve redis url")
	}
	instance := redis.NewClient(options)
	return instance, nil
}
