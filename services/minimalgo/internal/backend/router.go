package backend

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
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

func (r *RouterBackend) SetRedis(client *redis.Client) {
	r.redisDb = client
}

func (r *RouterBackend) LoadFromCache(ctx context.Context, router *Router) error {
	if r.redisDb == nil {
		return fmt.Errorf("Redis not set")
	}

	cmd1 := r.redisDb.HExists(ctx, utils.CreateStorageKey("global"), "router")
	if !cmd1.Val() {
		return errors.New("Routes do not exist on the cache")
	}

	cmd2 := r.redisDb.HGet(ctx, utils.CreateStorageKey("global"), "router")
	if data, err := cmd2.Bytes(); err != nil {
		return errors.Join(errors.New("Error trying to load routes from cache"), err)
	} else {
		if err := json.Unmarshal(data, router); err != nil {
			return err
		}
	}
	return cmd2.Err()
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
	Hash        string `json:"hash" yaml:"-"` // Hash of the name of the route used for quick lookup
	NameFr      string `json:"name_fr" yaml:"name_fr"`
	NameEn      string `json:"name_en" yaml:"name_en"`
	Domain      string `json:"domain" yaml:"domain"`
	Port        int    `json:"port" yaml:"port"`
	Active      bool   `json:"active" yaml:"active"`
	Environment string `json:"environment" yaml:"environment"`
	Version     int    `json:"version" yaml:"version"`
}

func (r *Route) Address() string {
	if r.Domain == "" {
		return ""
	} else {
		if r.Environment == "test" || r.Environment == "dev" {
			return fmt.Sprintf("dns:///localhost:%d", r.Port)
		} else {
			return fmt.Sprintf("dns:///%s:%d", r.Domain, r.Port)
		}
	}
}

type Router struct {
	RouterBackend
	Routes []*Route `json:"routes" yaml:"routes"`
}

func (r *Router) Find(name string) *Route {
	hashed := sha256.Sum256([]byte(name))
	for _, route := range r.Routes {
		if route.Hash == fmt.Sprintf("%x", hashed) {
			return route
		}
	}
	return nil
}

func (r *Router) LoadFile(ctx context.Context, filePath string) error {
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

	for _, route := range r.Routes {
		route.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(route.NameFr)))
	}

	// Save in redis directly
	r.Save(ctx, r)
	return nil
}

func NewRouter(ctx context.Context, filePath string) (*Router, error) {
	router := &Router{Routes: []*Route{}}
	if err := router.LoadFile(ctx, filePath); err != nil {
		return nil, err
	} else {
		return router, nil
	}
}
