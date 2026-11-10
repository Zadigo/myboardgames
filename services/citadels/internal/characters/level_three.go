package characters

import (
	"github.com/Zadigo/citadels/internal/models"
	"github.com/Zadigo/citadels/internal/utils"
)

type Magician struct {
	Character
}

func (m *Magician) Resolve(options models.CharacterOptions) error {
	if err := checkOptions(options); err == nil {
		switch options.Magician.Choice {
		case "exchange":
			utils.ExchangeDistricts(m.GetPlayer(), options.Target.GetPlayer())
		case "discard":
		default:
			return InvalidMagicianChoiceError
		}
		return nil
	} else {
		return err
	}
}

type Seer struct {
	Character
}

type Wizard struct {
	Character
}
