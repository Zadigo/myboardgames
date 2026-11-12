package tests

import (
	"testing"

	"github.com/Zadigo/citadels/internal/game"
	"github.com/stretchr/testify/assert"
)

func TestGames(t *testing.T) {
	t.Run("should be able to create games", func(t *testing.T) {
		t.Parallel()

		registry := game.NewGames()
		id := registry.Create()
		assert.NotNil(t, id)

		g, s := registry.Get(id)
		assert.NotNil(t, g)
		assert.True(t, s)
		assert.True(t, g.IsStarted())
	})

	t.Run("should remove game", func(t *testing.T) {

	})
}
