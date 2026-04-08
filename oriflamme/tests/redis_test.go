package tests

import (
	"context"
	"testing"
	"time"

	"github.com/Zadigo/oriflamme/internal/backend"
	"github.com/Zadigo/oriflamme/internal/backend/redisclient"
	"github.com/stretchr/testify/assert"
)

func TestRedisBroadcasting(t *testing.T) {
	ctx := context.Background()
	redisClient := redisclient.NewRedisClient("redis://@localhost:6379/0")
	
	tableUuid := "test-table-id"
	subcription := redisclient.NewSubscription(redisClient, ctx)
	registry := redisclient.NewBroadcastingRegistry(tableUuid, subcription, ctx)

	t.Run("Register new table", func(t *testing.T) {
		broadcaster := registry.GetOrCreate(tableUuid)

		t.Run("Register new clients", func(t *testing.T) {
			conns, server := CreateMultipleGameConn(t, "pauline", "julie", "aurelie")
			defer server.Close()

			for name, conn := range conns {
				defer conn.Close()

				// Must identify response
				message := backend.WebsocketMessage{}
				conn.ReadJSON(&message)

				// Identify using player uuid
				conn.WriteJSON(backend.WebsocketMessage{
					Action:     "identify",
					PlayerUuid: message.PlayerUuid,
					Username:   name,
				})

				// Read identify response
				message = backend.WebsocketMessage{}
				err := conn.ReadJSON(&message)
				assert.NoError(t, err, err)

				broadcaster.AddClient(conn)

				time.Sleep(1 * time.Second)
			}
		})
	})
}
