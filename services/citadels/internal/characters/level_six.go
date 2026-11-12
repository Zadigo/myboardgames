package characters

import (
	"github.com/Zadigo/citadels/internal/districts"
	"github.com/Zadigo/citadels/internal/models"
)

type Merchant struct {
	Character
}

func (m *Merchant) Resolve(options models.CharacterOptions) error {
	points := 0

	player := m.GetPlayer()

	// You gain one gold for each trade district in your city.
	for _, d := range player.GetDistricts() {
		if d == nil {
			continue
		}

		if d.GetDistrictType() == districts.Trade {
			points++
		}
	}

	player.AddPoints(points)

	// You gain one extra gold. You can use this ability
	// regardless of which resource you gathered this turn.
	player.IncreaseGold(1)
	return nil
}

type Trader struct {
	Character
}

type Alchemist struct {
	Character
}
