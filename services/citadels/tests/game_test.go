package tests

import (
	"testing"

	"github.com/Zadigo/citadels/internal/game"
	"github.com/Zadigo/citadels/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestGame(t *testing.T) {
	t.Run("should instantiate game", func(t *testing.T) {
		instance := game.NewGame()

		assert.IsType(t, instance, &game.Game{})
		assert.NotEmpty(t, instance.GetID())
		assert.Len(t, instance.GetPlayers(), 8)
		assert.Len(t, instance.GetCharacters(), 8)
		assert.Nil(t, instance.GetCrownOwner())
		assert.Equal(t, instance.GetStage(), models.WaitingForPlayers)
	})

	t.Run("should add player when below 8", func(t *testing.T) {
		t.Parallel()

		players := helperCreatePlayers(8)
		instance := game.NewGame()

		remainders := make([]int, 0)
		for _, player := range players {
			rem := instance.GetTable().AddPlayer(player)
			remainders = append(remainders, rem)
			assert.IsDecreasing(t, remainders, rem)
		}
	})
}
