package tests

import (
	"testing"

	"github.com/Zadigo/flipseven/backend"
)


func TestRedisBackend(t *testing.T) {
	serverConfig := backend.ServerConfig{
		Config: backend.ServerBaseConfig{
			Backends: &backend.ServerBackendsConfig{
				Redis: &backend.ServerBackendConfig{
					Url: "redis://:@localhost:6379/0",
				},
			},
		},
	}

	client, err := backend.CreateRedisClient(serverConfig.Config.Backends.Redis)
	if err != nil {
		t.Fatalf("Failed to create Redis client: %v", err)
	}

	if client == nil {
		t.Fatal("Redis client is nil")
	}
}
