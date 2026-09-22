package tests

import (
	"testing"
	"time"

	"github.com/Zadigo/oriflamme/internal/backend"
	"github.com/stretchr/testify/assert"
)

func TestRedisBroadcasting(t *testing.T) {
	t.Run("Register new table", func(t *testing.T) {
		conns, server := CreateMultipleGameConn(t, "pauline", "julie", "aurelie")
		defer server.Close()

		t.Run("Register new clients", func(t *testing.T) {

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

				time.Sleep(1 * time.Second)

				if name == "pauline" {
					// Read create game response
					conn.WriteJSON(backend.WebsocketMessage{
						Action: "create_game",
					})

					message = backend.WebsocketMessage{}
					err = conn.ReadJSON(&message)
					assert.NoError(t, err, err)
					assert.Equal(t, "create_game", message.Action)
				}
			}
		})
	})
}
