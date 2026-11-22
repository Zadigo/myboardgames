package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path"
	"path/filepath"

	"github.com/redis/go-redis/v9"
	"go.yaml.in/yaml/v3"
)

type RulesBackend struct {
	redisDb *redis.Client
}

func (r *RulesBackend) SetRedis(client *redis.Client) {
	r.redisDb = client
}

func (r *RulesBackend) Load(ctx context.Context, rules *Rule) error {
	if r.redisDb == nil {
		return fmt.Errorf("Redis not set")
	}

	cmd := r.redisDb.HGet(ctx, CreateStorageKey("global"), "rules")
	if data, err := cmd.Bytes(); err != nil {
		return err
	} else {
		if err := json.Unmarshal(data, rules); err != nil {
			return err
		}
	}
	return cmd.Err()
}

func (r *RulesBackend) Save(ctx context.Context, rules *Rule) error {
	if r.redisDb == nil {
		return fmt.Errorf("Redis database not set")
	}

	data, _ := json.Marshal(rules)

	cmd := r.redisDb.HSet(ctx, CreateStorageKey("global"), "rules", data)
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}

type TimeLimitRule struct {
	State bool `yaml:"state"`
	Value int  `yaml:"value"`
}

type PlayerRule struct {
	Max int `yaml:"max"`
	Min int `yaml:"min"`
}

type GlobalRules struct {
	Players     PlayerRule    `yaml:"players"`
	StartsFirst string        `yaml:"starts_first"`
	TimeLimit   TimeLimitRule `yaml:"time_limit"`
	MaxPoints   int           `yaml:"max_points"`
}

type Rule struct {
	RulesBackend
	Rules GlobalRules `yaml:"rules"`
}

func LoadRules(filePath string, rule *Rule) error {
	filep, err := filepath.Abs(path.Join(filePath, "rules.yaml"))
	if err != nil {
		return err
	}

	slog.Info("Loaded rules for game", "rules", filep)

	if content, err := os.ReadFile(filep); err != nil {
		return err
	} else {
		if err := yaml.Unmarshal(content, rule); err != nil {
			return err
		}
		return nil
	}
}
