package tests

import (
	"testing"
	"time"

	"github.com/Zadigo/oriflamme/internal/backend"
	"github.com/stretchr/testify/assert"
)

func TestRegistry(t *testing.T) {
	registry := backend.NewGameRegistry()

	t.Run("Has Uuid", func(t *testing.T) {
		assert.NotEqual(t, "", registry.Uuid)
	})

	t.Run("Is running", func(t *testing.T) {
		assert.True(t, registry.IsRunning)
	})

	t.Run("Is started", func(t *testing.T) {
		assert.False(t, registry.IsStarted)
	})

	t.Run("Has start date", func(t *testing.T) {
		assert.NotNil(t, registry.StartedAt)
	})

	conns, server := CreateMultipleGameConn(t, "pauline", "marie")
	defer server.Close()

	clients := ConvertToWebsocketClient(t, conns)

	for _, client := range clients {
		registry.Clients[client.Uuid] = client
	}

	t.Run("Start game", func(t *testing.T) {
		d, ok := t.Deadline()
		d.Add(time.Duration(5 * time.Second))

		registry.StartGame("test-table-id")
		assert.False(t, ok)
	})
}
