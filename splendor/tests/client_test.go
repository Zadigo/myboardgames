package tests

import (
	"testing"

	"github.com/Zadigo/splendor/internal/backend"
	"github.com/stretchr/testify/assert"
)

func TestWebsocketClient(t *testing.T) {
	conn, server, _ := CreateWebsocketConnection(t)

	defer server.Close()
	defer conn.Close()

	client := backend.NewWebsocketClient("alice", conn, true)
	message := &backend.WebsocketMessage{}

	t.Run("Should not be nil", func(t *testing.T) {
		assert.NotNil(t, client)

		t.Run("Should be able to identify", func(t *testing.T) {
			err := conn.ReadJSON(&message)
			assert.NoError(t, err)
			assert.Contains(t, message.Action, "must_identify")
			assert.NotEmpty(t, message.PlayerUuid)

			playerUuid := message.PlayerUuid

			t.Run("Should be able to send identification", func(t *testing.T) {
				err := client.SendJsonMessage(backend.WebsocketMessage{
					Action:     "identify",
					PlayerUuid: playerUuid,
					Username:   "alice",
				})
				assert.NoError(t, err)
			})

			t.Run("Should be able to create table", func(t *testing.T) {
				err := client.SendJsonMessage(backend.WebsocketMessage{
					Action:       "create_table",
					PlayerUuid:   playerUuid,
					IsNormalGame: true,
				})
				assert.NoError(t, err)

				err = conn.ReadJSON(&message)
				assert.NoError(t, err)
				assert.Contains(t, message.Action, "table_created")

				t.Log(message)
			})
		})
	})
}
