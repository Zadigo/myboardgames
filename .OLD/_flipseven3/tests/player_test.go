package tests

import (
	"testing"

	"github.com/Zadigo/flipseven3/internal/backend/logic"
	"github.com/stretchr/testify/assert"
)

func TestPlayer(t *testing.T) {
	conn, server := CreateConnection(t)
	defer server.Close()

	// player := logic.CreatePlayer(conn, "testuser", uuid.NewString(), false)
	client := logic.NewWebsocketClient(conn)
	client.Username = "testuser"

	cards := CreateTestCards(t)
	queue := logic.NewCardQueue()
	queue.Queue = cards

	for i := range 5 {
		queue.Queue[i].Owner = client
	}

	t.Run("Should return the player's cards", func(t *testing.T) {
		playerCards := client.GetCards(queue)

		for _, card := range playerCards {
			assert.Equal(t, client.Uuid, card.Owner.Uuid)
		}

		assert.Equal(t, len(playerCards), 5)
	})

	t.Run("Score should have correct value", func(t *testing.T) {
		cumulatedScore := 0
		scores := []int{0, 1, 2, 2, 3}

		for _, score := range scores {
			cumulatedScore += score
		}

		assert.Equal(t, cumulatedScore, client.GetCurrentScore(queue))
	})
}
