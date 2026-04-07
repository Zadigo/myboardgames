package tests

import (
	"testing"

	"github.com/Zadigo/flipseven2/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestCards(t *testing.T) {
	t.Run("Create number cards", func(t *testing.T) {
		cards := models.GetNumberCards()
		assert.Len(t, cards, 79)

		cardsWithValue0 := 0
		for _, card := range cards {
			if card.Value == 0 {
				cardsWithValue0++
			}
		}
		assert.Equal(t, 1, cardsWithValue0)

		cardsWithValue10 := 0
		for _, card := range cards {
			if card.Value == 10 {
				cardsWithValue10++
			}
		}
		assert.Equal(t, 10, cardsWithValue10)
	})

	t.Run("Create special cards", func(t *testing.T) {
		cards := models.GetSpecialCards()
		assert.Len(t, cards, 9)
	})
}
