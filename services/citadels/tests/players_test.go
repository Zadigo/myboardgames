package tests

import (
	"testing"

	"github.com/Zadigo/citadels/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestPlayers(t *testing.T) {
	t.Run("should instiantiate player", func(t *testing.T) {
		player := models.NewPlayer()

		assert.IsType(t, player, &models.Player{})
		assert.Len(t, player.GetDistricts(), 0)
		assert.Equal(t, player.GoldCount(), 2)
		assert.Nil(t, player.GetSelectedCharacter())
	})

	t.Run("should get ID", func(t *testing.T) {
		t.Parallel()

		player := models.NewPlayer()

		result := player.GetId()
		assert.NotEmpty(t, result)
	})

	t.Run("should apply choice gather resources", func(t *testing.T) {
		t.Parallel()

		t.Run("take two golds", func(t *testing.T) {
			player := models.NewPlayer()
			choices := models.PlayerChoices{GatherResources: true, TakeTwoGolds: true}
			
			player.ApplyChoice(choices)
			
			assert.Equal(t, 4, player.GoldCount())
		})

		t.Run("draw two district cards", func(t *testing.T) {

		})

		t.Run("should fail with GatherResouces false and BuildDistrict false", func(t *testing.T) {

		})
	})

	t.Run("should increase gold", func(t *testing.T) {
		t.Parallel()

		player := models.NewPlayer()

		player.IncreaseGold(2)
		assert.Equal(t, player.GoldCount(), 4)
	})

	t.Run("should descrease gold by margin", func(t *testing.T) {
		t.Parallel()

		player := models.NewPlayer()

		player.DecreaseGold(1)
		assert.Equal(t, 1, player.GoldCount())
	})

	t.Run("should descrease equals to zero", func(t *testing.T) {
		t.Parallel()

		player := models.NewPlayer()

		player.DecreaseGold(15)
		assert.Equal(t, 0, player.GoldCount())
	})
}
