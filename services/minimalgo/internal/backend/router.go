package backend

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log/slog"
	"myboardgames/minimalgo/internal/utils"
	"net/url"
	"os"
	"path/filepath"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert/yaml"
)

type RouterBackend struct {
	redisDb *redis.Client
}

func (r *RouterBackend) LoadFromCache(ctx context.Context, router *Router) error {
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
	Hash        string `json:"hash" yaml:"-"` // Hash of the name of the route used for quick lookup
	NameFr      string `json:"name_fr" yaml:"name_fr"`
	NameEn      string `json:"name_en" yaml:"name_en"`
	Domain      string `json:"domain" yaml:"domain"`
	Port        int    `json:"port" yaml:"port"`
	Active      bool   `json:"active" yaml:"active"`
	Environment string `json:"environment" yaml:"environment"`
	Version     int    `json:"version" yaml:"version"`
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

func (r *Router) FindDomain(name string) (string, error) {
	if route := r.Find(name); route == nil {
		return "", fmt.Errorf("route not found")
	} else {
		if _domain, err := url.Parse(route.Domain); err != nil {
			return "", err
		} else {
			if route.Environment == "dev" || route.Environment == "test" {
				return fmt.Sprintf("%s:%d", "localhost", route.Port), nil
			} else {
				return _domain.String(), nil
			}
		}
	}
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
		for _, route := range router.Routes {
			route.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(route.NameFr)))
		}
		return router, nil
	}
}
