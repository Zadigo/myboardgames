package tests

import (
	"testing"

	"github.com/Zadigo/flipseven3/internal/backend/logic"
	"github.com/stretchr/testify/assert"
)

func TestCardQueue(t *testing.T) {
	conn, server := CreateConnection(t)
	defer server.Close()

	// julie := logic.CreatePlayer(conn, "testuser", "testuuid", false)
	// pauline := logic.CreatePlayer(conn, "testuser2", "testuuid2", false)

	julie := logic.NewWebsocketClient(conn)
	julie.Username = "julie"

	pauline := logic.NewWebsocketClient(conn)
	pauline.Username = "pauline"

	queue := CreateCardQueue(t)

	t.Run("Should flip card", func(t *testing.T) {
		queue.FlipCard(julie)

		assert.Equal(t, queue.Queue[0].Owner.Uuid, julie.Uuid)
		assert.Equal(t, queue.ResolutionIndex, 0)

		// Flip the next card
		queue.FlipCard(pauline)

		assert.Equal(t, queue.Queue[1].Owner.Uuid, pauline.Uuid)
		assert.Equal(t, queue.ResolutionIndex, 1)
	})
}
