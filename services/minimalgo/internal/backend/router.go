package backend

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"myboardgames/minimalgo/internal/utils"
	"os"
	"path/filepath"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert/yaml"
)

type RouterBackend struct {
	redisDb *redis.Client
}

func (r *RouterBackend) Load(ctx context.Context, router *Router) error {
	if r.redisDb == nil {
		return fmt.Errorf("Redis not set")
	}

	cmd := r.redisDb.HGet(ctx, utils.CreateStorageKey("global"), "router")
	if data, err := cmd.Bytes(); err != nil {
		return err
	} else {
		if err := json.Unmarshal(data, router); err != nil {
			return err
		}
	}
	return cmd.Err()
}

func (r *RouterBackend) Save(ctx context.Context, router *Router) error {
	if r.redisDb == nil {
		return fmt.Errorf("Redis database not set")
	}

	data, _ := json.Marshal(router)

	cmd := r.redisDb.HSet(ctx, utils.CreateStorageKey("global"), "router", data)
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}

type Route struct {
	NameFr      string `yaml:"name_fr"`
	NameEn      string `yaml:"name_en"`
	Domain      string `yaml:"domain"`
	Port        int    `yaml:"port"`
	Active      bool   `yaml:"active"`
	Environment string `yaml:"environment"`
	Version     int    `yaml:"version"`
}

type Router struct {
	RouterBackend
	Routes []*Route `yaml:"routes"`
}

func (r *Router) Load(ctx context.Context, filePath string) error {
	filep, err := filepath.Abs(filepath.Join(filePath, "routes.yaml"))
	if err != nil {
		return err
	}

	content, err := os.ReadFile(filep)
	if err != nil {
		return err
	}

	// Unmarshal into r (which is already a *Router pointer)
	if err := yaml.Unmarshal(content, r); err != nil {
		return err
	}

	slog.Info("Loaded routes", "file", filep)
	return nil
}

func NewRouter(ctx context.Context, filePath string) (*Router, error) {
	router := &Router{Routes: []*Route{}}
	if err := router.Load(ctx, filePath); err != nil {
		return nil, err
	} else {
		return router, nil
	}
}
