package tests

import (
	"testing"

	"github.com/Zadigo/flipseven2/internal/models"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestPlayers(t *testing.T) {
	type testCase struct {
		usemultiplier bool
		isInitiator   bool
		expectedScore int
		error         error
	}

	testCases := []testCase{
		{
			usemultiplier: false,
			isInitiator:   false,
			expectedScore: 0,
			error:         nil,
		},
	}

	for _, tc := range testCases {
		player := models.CreatePlayer(&websocket.Conn{}, "pauline", "test-table-uuid", tc.isInitiator)
		assert.NotNil(t, player)

		t.Run("Freshly created player", func(t *testing.T) {
			assert.Equal(t, tc.isInitiator, player.IsInitiator)
			if tc.expectedScore == 0 {
				assert.Equal(t, tc.expectedScore, player.Score)
			}
		})
	}

	testCase2 := []testCase{
		{
			usemultiplier: false,
			isInitiator:   false,
			expectedScore: 10,
			error:         nil,
		},
		{
			usemultiplier: true,
			isInitiator:   false,
			expectedScore: 20,
			error:         nil,
		},
	}

	for _, tc := range testCase2 {
		t.Run("Player receives cards", func(t *testing.T) {
			cards := []models.Card{
				{
					Value:    10,
					Category: "number",
					IsNumber: true,
				},
			}

			player := models.CreatePlayer(&websocket.Conn{}, "pauline", "test-table-uuid", tc.isInitiator)
			player.Cards = append(player.Cards, &cards[0])

			// Multiplier cards should multiply the score
			if tc.usemultiplier {
				multiplierCard := models.Card{
					Value:        2,
					IsMultiplier: true,
				}

				player.CalculatePoints(&multiplierCard)
				assert.Equal(t, tc.expectedScore, player.Score)
			}
		})
	}
}
