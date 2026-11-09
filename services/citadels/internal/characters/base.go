package characters

import (
	"fmt"
	"uuid"

	"github.com/Zadigo/citadels/internal/models"
)

// const (
// 	Assassin = "assassin"
//  Witch = "witch"
//  Magistrate = "magistrate"
//  Thief = "thief"
//  Spy = "spy"
//  Magician = "magician"
//  Blackmailer = "blackmailer"
//  Wizard = "wizard"
//  Seer = "seer"
//  King = "king"
//  Patrician = "patrician"
// 	Emperor = "emperor"
//  Bishop = "bishop"
//  Abbot = "abbot"
// 	Alchemist = "alchemist"
//	Trader = "trader"
// 	Architect = "architect"
//  Cardinal = "cardinal"
//  Navigator = "navigator"
//  Merchant = "merchant"
//  Scholar = "scholar"
//  Warlord = "warlord"
//  Queen = "queen"
//  Diplomat = "diplomat"
//  Artist = "artist"
//  TaxCollector = "tax_collector"
//  Marshal = "marshal"
// )

func checkOptions(options models.CharacterOptions) error {
	if options.Target.GetPlayer() == nil {
		return fmt.Errorf("Target has no player")
	}

	if options.CurrentGame == nil {
		return fmt.Errorf("No game")
	}

	return nil
}

type Character struct {
	models.BaseCharacter
}

func (c *Character) Resolve(options models.CharacterOptions) error {
	return nil
}

func (c *Character) Kill() {
	c.Killed = true
}

type King struct {
	Character
}

type Bishop struct {
	Character
}

type Merchant struct {
	Character
}

type Architect struct {
	Character
}

type Warlord struct {
	Character
}

// NewCharacter creates a new Citadels character:
//
//	character := NewCharacter()
//	character.SetPlayer(player)
//
// A character can then be resolved by calling the Resolve method
//
//	err := character.Resolve(models.CharacterOptions{})
func NewCharacter(name string) (models.CharacterInterface, error) {
	characterId := uuid.New().String()
	baseCharacter := models.BaseCharacter{
		ID:        characterId,
		Level:     0,
		Name:      "",
		Image:     "",
		Killed:    false,
		Bewitched: nil,
		Player:    nil,
	}

	switch name {
	case "assassin":
		baseCharacter.Level = 1
		baseCharacter.Name = "assassin"
		return &Assassin{Character: Character{baseCharacter}}, nil
	case "thief":
		baseCharacter.Level = 2
		baseCharacter.Name = "thief"
		return &Thief{Character: Character{baseCharacter}}, nil
	case "magician":
		baseCharacter.Level = 3
		baseCharacter.Name = "magician"
		return &Magician{Character: Character{baseCharacter}}, nil
	case "king":
		baseCharacter.Level = 4
		baseCharacter.Name = "king"
		return &King{Character: Character{baseCharacter}}, nil
	case "bishop":
		baseCharacter.Level = 5
		baseCharacter.Name = "bishop"
		return &Bishop{Character: Character{baseCharacter}}, nil
	case "merchant":
		baseCharacter.Level = 6
		baseCharacter.Name = "merchant"
		return &Merchant{Character: Character{baseCharacter}}, nil
	case "warlord":
		baseCharacter.Level = 8
		baseCharacter.Name = "warlord"
		return &Warlord{Character: Character{baseCharacter}}, nil
	case "architect":
		baseCharacter.Level = 8
		baseCharacter.Name = "architect"
		return &Architect{Character: Character{baseCharacter}}, nil
	}
	return nil, fmt.Errorf("unknown character: %s", name)
}
