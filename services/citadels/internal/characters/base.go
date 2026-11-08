package characters

import (
	"fmt"
	"uuid"

	"github.com/Zadigo/citadels/internal/models"
)

type Character struct {
	models.BaseCharacter
}

func (c *Character) Resolve(options models.CharacterOptions) error {
	for _, c := range options.CurrentGame.GetPlayers() {
		if c.SelectedCharacter == options.Target.GetId() {
			// Do something with the player
		}
	}
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

func (c *Character) GetLevel() int {
	return c.Level
}

func (c *Character) Kill() {
	c.Killed = true
}

func (c *Character) SetPlayer(id string) {
	c.Player = id
}

func (c *Character) GetPlayer() string {
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

type Thief struct {
	Character
}

func (t *Thief) Resolve(options models.CharacterOptions) error {
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
		baseCharacter.Level = 7
		baseCharacter.Name = "warlord"
		return &Warlord{Character: Character{baseCharacter}}, nil
	case "architect":
		baseCharacter.Level = 8
		baseCharacter.Name = "architect"
		return &Architect{Character: Character{baseCharacter}}, nil
	}
	return nil, fmt.Errorf("unknown character: %s", name)
}
