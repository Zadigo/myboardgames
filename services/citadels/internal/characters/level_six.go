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

	// You gain one gold for each trade district in your city.
	for _, d := range m.GetPlayer().GetDistricts() {
		if d.GetDistrictType() == districts.Trade {
			points++
		}
	}

	// You gain one extra gold. You can use this ability
	// regardless of which resource you gathered this turn.
	points++

	m.GetPlayer().AddPoints(points)
	return nil
}

type Trader struct {
	Character
}

type Alchemist struct {
	Character
}
