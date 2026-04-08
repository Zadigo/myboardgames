package tests

import (
	"testing"

	"github.com/Zadigo/oriflamme/internal/backend"
	"github.com/stretchr/testify/assert"
)

func TestLiveGameHandler(t *testing.T) {
	conn, server := CreateLiveGameConn(t)

	defer conn.Close()
	defer server.Close()

	t.Run("Initial connection", func(t *testing.T) {
		message := backend.WebsocketMessage{}
		err := conn.ReadJSON(&message)
		assert.NoError(t, err, err)
		assert.Equal(t, "must_identify", message.Action)
		assert.NotEmpty(t, message.PlayerUuid)

		conn.WriteJSON(backend.WebsocketMessage{
			Action:     "identify",
			PlayerUuid: message.PlayerUuid,
			Username:   "pauline",
		})
	})

	t.Run("Create game", func(t *testing.T) {})

	t.Run("Another player joins the game", func(t *testing.T) {})

	t.Run("Start game", func(t *testing.T) {})

	t.Run("Players perform actions play card place card", func(t *testing.T) {})

	t.Run("Player A performs play card action reveal", func(t *testing.T) {})

	t.Run("Player B performs play card action place token", func(t *testing.T) {})

	t.Run("Resolve queue", func(t *testing.T) {})

	t.Run("Player A performs play card action stack card", func(t *testing.T) {})

	t.Run("Player B performs unrecognized card action", func(t *testing.T) {})

	t.Run("Player B performs unrecognized action", func(t *testing.T) {})

	t.Run("Player disconnects", func(t *testing.T) {})

	t.Run("Player reconnects within a given time limit", func(t *testing.T) {})

	t.Run("End game", func(t *testing.T) {})
}
