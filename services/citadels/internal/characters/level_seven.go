package characters

import (
	"fmt"

	"github.com/Zadigo/citadels/internal/models"
)

type Architect struct {
	Character
}

type Navigator struct {
	Character
}

func (n *Navigator) Resolve(options models.CharacterOptions) error {
	n.CanBuildDistricts = false
	if options.Navigator.GainFourGolds && options.Navigator.GainFourCards {
		return fmt.Errorf("Navigator cannot gain four golds and four cards at the same time. Choose between the two options")
	}

	if !options.Navigator.GainFourCards && !options.Navigator.GainFourGolds {
		return fmt.Errorf("Navigator must choose an option")
	}

	if options.Navigator.GainFourGolds {
		n.GetPlayer().IncreaseGold(4)
	} else {
		// Gain four cards from the deck
	}

	return nil
}

type Scholar struct {
	Character
}
