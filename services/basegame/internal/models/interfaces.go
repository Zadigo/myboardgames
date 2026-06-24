package models

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type AppConfigInterface interface {
	LoadConfig(app AppInterface) error
}

type AppExtraInterface interface {
	GetRedisClient() *redis.Client
	GetBaseDir() string
}

type AppInterface interface {
	AppExtraInterface
	Start() error
	GetContext() context.Context
}

type ServerAppInterface any
