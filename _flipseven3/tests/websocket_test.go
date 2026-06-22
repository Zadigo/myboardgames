package tests

import (
	"testing"

	"github.com/Zadigo/flipseven3/internal/backend/logic"
	"github.com/stretchr/testify/assert"
)

func TestWebsocketClient(t *testing.T) {
	conn, server := CreateConnection(t)
	defer server.Close()

	var tableUuid string

	message := &logic.WebsocketMessage{}

	t.Run("Should receive player uuid", func(t *testing.T) {
		err := conn.ReadJSON(message)
		assert.NoError(t, err)
		assert.NotNil(t, message.PlayerUuid)
	})

	t.Run("Should identify user", func(t *testing.T) {
		conn.WriteJSON(logic.WebsocketMessage{
			Action:     "identify",
			Username:   "julie",
			PlayerUuid: message.PlayerUuid,
		})

		err := conn.ReadJSON(message)
		assert.NoError(t, err)
		assert.Equal(t, message.Message, "Identification successful")
	})

	t.Run("Should be able to create a game", func(t *testing.T) {
		conn.WriteJSON(logic.WebsocketMessage{
			Action: "create_game",
		})

		err := conn.ReadJSON(message)
		assert.NoError(t, err)
		assert.NotNil(t, message.GameRegistry)

		tableUuid = message.GameRegistry.Uuid
	})

	t.Run("Should be able to start game", func(t *testing.T) {
		conn.WriteJSON(logic.WebsocketMessage{
			Action:    "start_game",
			TableUuid: tableUuid,
		})

		err := conn.ReadJSON(message)
		assert.NoError(t, err)
		assert.Equal(t, message.Action, "start_game")

		t.Logf("message: %v", message)
	})

	t.Run("Should be able to flip card", func(t *testing.T) {
		
	})
}
