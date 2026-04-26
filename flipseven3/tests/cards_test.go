package tests

import (
	"testing"

	"github.com/Zadigo/flipseven3/internal/backend/logic"
	"github.com/stretchr/testify/assert"
)

func TestCards(t *testing.T) {
	t.Run("Should have correct number of number cards", func(t *testing.T) {
		cards := logic.GetNumberCards()
		assert.Equal(t, 79, len(cards))
	})

	t.Run("Should have correct number of special cards", func(t *testing.T) {
		cards := logic.GetSpecialCards()
		assert.Equal(t, 9, len(cards))
	})
}
