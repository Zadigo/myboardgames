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

type Character struct {
	models.BaseCharacter
}

func (c *Character) Resolve(options models.CharacterOptions) error {
	return nil
}

func (c *Character) GetId() string {
	return c.ID
}

func (c *Character) GetName() string {
	return c.Name
}

func (c *Character) IsKilled() bool {
	return c.Killed
}

func (c *Character) IsBewitched() bool {
	return c.Bewitched
}

func (c *Character) GetLevel() int {
	return c.Level
}

func (c *Character) Kill() {
	c.Killed = true
}

func (c *Character) SetPlayer(player models.PlayerInterface) {
	c.Player = player
}

func (c *Character) GetPlayer() models.PlayerInterface {
	return c.Player
}

type Assassin struct {
	Character
}

func (a *Assassin) Resolve(options models.CharacterOptions) error {
	if options.Target == nil {
		return fmt.Errorf("Assassin needs target")
	} else {
		options.Target.Kill()
		return nil
	}
}

// Call the name of another character whom you wish to rob.
// When a player reveals that character to take their turn, you immediately
// take all of their gold. You cannot rob the rank 1 character (Assassin, Witch, or Magistrate),
// the killed character, or the bewitched character.
type Thief struct {
	Character
}

func (t *Thief) Resolve(options models.CharacterOptions) error {
	if options.Target.GetPlayer() == nil {
		return fmt.Errorf("Target has no player")
	}

	if options.CurrentGame == nil {
		return fmt.Errorf("No game")
	}

	for _, p := range options.CurrentGame.GetPlayers() {
		// Skip the player that is attacking
		// which is "self"
		if p == t.GetPlayer() {
			return CannotAttackSelfError
		}

		selectedCharacter := p.GetSelectedCharacter()
		if selectedCharacter == nil {
			return TargetHasNoCharacterError
		}

		if p.GetSelectedCharacter().GetId() == options.Target.GetId() {
			if selectedCharacter.GetLevel() == 1 {
				return CannotStealLevelOneError
			}

			if selectedCharacter.IsKilled() {
				return CannotStealKilledError
			}

			if selectedCharacter.IsBewitched() {
				return CannotStealBewitchedError
			}

			goldCount := selectedCharacter.GetPlayer().GoldCount()
			selectedCharacter.GetPlayer().DecreaseGold(goldCount)
			t.GetPlayer().IncreaseGold(goldCount)
		}
	}
	return nil
}

type Magician struct {
	Character
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

func NewCharacter(name string) (models.CharacterInterface, error) {
	characterId := uuid.New().String()
	baseCharacter := models.BaseCharacter{
		ID:    characterId,
		Level: 0,
		Name:  "",
		Image: "",
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
