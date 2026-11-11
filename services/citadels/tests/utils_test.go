package tests

import (
	"testing"

	"github.com/Zadigo/citadels/internal/districts"
	"github.com/Zadigo/citadels/internal/models"
	"github.com/Zadigo/citadels/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestExchangeDistricts(t *testing.T) {
	players := helperCreatePlayers(2)

	p1 := players[0]
	p2 := players[1]

	t.Run("players have no districts", func(t *testing.T) {
		utils.ExchangeDistricts(p1, p2)

		assert.Len(t, p1.GetDistricts(), 0)
		assert.Len(t, p2.GetDistricts(), 0)
	})

	t.Run("players exhange districts", func(t *testing.T) {
		helperAddDistrictsToPlayers(p1, p2)

		assert.Len(t, p1.GetDistricts(), 1)
		assert.Len(t, p2.GetDistricts(), 3)

		utils.ExchangeDistricts(p1, p2)

		assert.Greater(t, len(p1.GetDistricts()), 0)
		assert.Greater(t, len(p2.GetDistricts()), 0)

		assert.Len(t, p1.GetDistricts(), 3)
		assert.Len(t, p2.GetDistricts(), 1)
	})
}

func TestCalculateScore(t *testing.T) {
	t.Run("should be 0", func(t *testing.T) {
		t.Parallel()

		player := models.NewPlayer(nil)

		player.AddDistrict(&districts.District{})
		player.AddDistrict(&districts.District{})

		assert.Equal(t, 0, utils.CalculateScore(player, utils.ScoringOptions{}))
	})

	t.Run("should be equal to 2", func(t *testing.T) {
		t.Parallel()

		player := models.NewPlayer(nil)

		player.AddDistrict(&districts.District{Cost: 1})
		player.AddDistrict(&districts.District{Cost: 1})

		assert.Equal(t, 2, utils.CalculateScore(player, utils.ScoringOptions{}))
	})

	t.Run("should be equal to 8 when each types", func(t *testing.T) {
		t.Parallel()

		player := models.NewPlayer(nil)

		player.AddDistrict(&districts.District{Cost: 1, DistrictType: districts.Religious})
		player.AddDistrict(&districts.District{Cost: 1, DistrictType: districts.Nobel})
		player.AddDistrict(&districts.District{Cost: 1, DistrictType: districts.Military})
		player.AddDistrict(&districts.District{Cost: 1, DistrictType: districts.Trade})
		player.AddDistrict(&districts.District{Cost: 1, DistrictType: districts.Unique})
		// 1 + 1 + 1 + 1 + 1 + 3
		assert.Equal(t, 8, utils.CalculateScore(player, utils.ScoringOptions{}))
	})

	t.Run("should be equal to 5 when finished first", func(t *testing.T) {
		t.Parallel()

		player := models.NewPlayer(nil)

		player.AddDistrict(&districts.District{Cost: 1})

		assert.Equal(t, 5, utils.CalculateScore(player, utils.ScoringOptions{FinishedFirst: true}))
	})
}
