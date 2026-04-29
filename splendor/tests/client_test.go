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

	t.Run("Should not be nil", func(t *testing.T) {
		assert.NotNil(t, client)

		t.Run("Should be able to send json message", func(t *testing.T) {
			err := client.SendJsonMessage(backend.WebsocketMessage{})
			assert.NoError(t, err)
		})

		t.Run("Should be able to receive json message", func(t *testing.T) {
			message := client.ReadJsonMessage()
			assert.IsType(t, backend.WebsocketMessage{}, message)
			assert.Equal(t, backend.WebsocketMessage{}, message)
			// assert.NotNil(t, message)
			// assert.Empty(t, message)
			// assert.EqualValues(t, backend.WebsocketMessage{}, message)
			// assert.NotEqual(t, backend.WebsocketMessage{SomeField: "value"}, message)
			// assert.NotEqualValues(t, backend.WebsocketMessage{SomeField: "value"}, message)
			// assert.False(t, message.SomeField != "")
			// assert.True(t, message.SomeField == "")
			// assert.Zero(t, message.SomeField)
			// assert.NotZero(t, "non-empty")
			// assert.Contains(t, "hello world", "world")
			// assert.NotContains(t, "hello world", "universe")
			// assert.Len(t, message.SomeField, 0)
			// assert.NotLen(t, "non-empty", 0)
		})
	})
}
