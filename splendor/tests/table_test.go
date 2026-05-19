package tests

import (
	"testing"

	"github.com/Zadigo/splendor/internal/backend"
	"github.com/stretchr/testify/assert"
)

func TestPlayingTable(t *testing.T) {
	table := backend.NewPlayingTable(true)

	t.Run("Should have UUID and date set", func(t *testing.T) {
		assert.NotEmpty(t, table.Uuid)
		assert.False(t, table.StartedAt.IsZero())
	})

	t.Run("Should add a player without error", func(t *testing.T) {
		client := &backend.WebsocketClient{
			Player: &backend.Player{
				Uuid: "test-player-uuid",
			},
		}
		err := table.AddPlayer(client)
		assert.NoError(t, err)

		t.Run("Should return the correct number of players", func(t *testing.T) {
			numPlayers := table.NumberOfPlayers()
			assert.Equal(t, 1, numPlayers)
		})

		t.Cleanup(func() {
			table.CardsLevelOne = []backend.CardInterface{}
			table.CardsLevelTwo = []backend.CardInterface{}
			table.CardsLevelThree = []backend.CardInterface{}

			table.DeckLevelOne = []backend.CardInterface{}
			table.DeckLevelTwo = []backend.CardInterface{}
			table.DeckLevelThree = []backend.CardInterface{}

			table.CurrentPlayer = nil
			table.CurrentRound = 0
			table.IsStarted = false
		})
	})

	t.Run("Should create a new deck without error", func(t *testing.T) {
		err := table.StartGame()
		assert.NoError(t, err)

		assert.Len(t, table.DeckLevelOne, 40)
		assert.Len(t, table.DeckLevelTwo, 30)
		assert.Len(t, table.DeckLevelThree, 20)

		assert.Len(t, table.CardsLevelOne, 4)
		assert.Len(t, table.CardsLevelTwo, 4)
		assert.Len(t, table.CardsLevelThree, 4)

		t.Run("Should end the game without error", func(t *testing.T) {
			err := table.EndGame()
			assert.NoError(t, err)
		})
	})
}
