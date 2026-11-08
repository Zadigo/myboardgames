package tests

import (
	"testing"

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
		helperAddDistrictsToPlayers(players)
		utils.ExchangeDistricts(p1, p2)

		assert.Greater(t, len(p1.GetDistricts()), 0)
		assert.Greater(t, len(p2.GetDistricts()), 0)

		for _, d := range p1.GetDistricts() {
			assert.Equal(t, d.GetPlayer(), p1.GetId())
		}
	})
}
