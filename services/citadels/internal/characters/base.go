package characters

import (
	"fmt"
	"uuid"
)

type CharacterInterface interface {
	Resolve(options CharacterOptions) error
	GetId() string
	GetName() string
	IsKilled() bool
	GetLevel() int
	Kill()
	SetPlayer(id string)
	GetPlayer() string
}

type BaseCharacter struct {
	ID     string
	Level  int
	Name   string
	Image  string
	Killed bool
	Player string
}

func (b *BaseCharacter) Resolve(options CharacterOptions) error {
	for _, c := range options.CurrentGame.GetPlayers() {
		if c.SelectedCharacter == options.Target.GetId() {
			// Do something with the player
		}
	}
	return nil
}

func (b *BaseCharacter) GetId() string {
	return b.ID
}

func (b *BaseCharacter) GetName() string {
	return b.Name
}

func (b *BaseCharacter) IsKilled() bool {
	return b.Killed
}

func (b *BaseCharacter) GetLevel() int {
	return b.Level
}

func (b *BaseCharacter) Kill() {
	b.Killed = true
}

func (b *BaseCharacter) SetPlayer(id string) {
	b.Player = id
}

func (b *BaseCharacter) GetPlayer() string {
	return b.Player
}

type Assassin struct {
	BaseCharacter
}

func (a *Assassin) Resolve(options CharacterOptions) error {
	if options.Target == nil {
		return fmt.Errorf("Assassin needs target")
	} else {
		options.Target.Kill()
		return nil
	}
}

type Thief struct {
	BaseCharacter
}

func (t *Thief) Resolve(options CharacterOptions) error {
	return nil
}

type Magician struct {
	BaseCharacter
}

type King struct {
	BaseCharacter
}

type Bishop struct {
	BaseCharacter
}

type Merchant struct {
	BaseCharacter
}

type Architect struct {
	BaseCharacter
}

type Warlord struct {
	BaseCharacter
}

func NewCharacter(name string) (CharacterInterface, error) {
	characterId := uuid.New().String()

	baseCharacter := BaseCharacter{
		ID:    characterId,
		Level: 0,
		Name:  "",
		Image: "",
	}

	switch name {
	case "assassin":
		baseCharacter.Level = 1
		baseCharacter.Name = "assassin"
		return &Assassin{BaseCharacter: baseCharacter}, nil
	case "thief":
		baseCharacter.Level = 2
		baseCharacter.Name = "thief"
		return &Thief{BaseCharacter: baseCharacter}, nil
	case "magician":
		baseCharacter.Level = 3
		baseCharacter.Name = "magician"
		return &Magician{BaseCharacter: baseCharacter}, nil
	case "king":
		baseCharacter.Level = 4
		baseCharacter.Name = "king"
		return &King{BaseCharacter: baseCharacter}, nil
	case "bishop":
		baseCharacter.Level = 5
		baseCharacter.Name = "bishop"
		return &Bishop{BaseCharacter: baseCharacter}, nil
	case "merchant":
		baseCharacter.Level = 6
		baseCharacter.Name = "merchant"
		return &Merchant{BaseCharacter: baseCharacter}, nil
	case "warlord":
		baseCharacter.Level = 7
		baseCharacter.Name = "warlord"
		return &Warlord{BaseCharacter: baseCharacter}, nil
	case "architect":
		baseCharacter.Level = 8
		baseCharacter.Name = "architect"
		return &Architect{BaseCharacter: baseCharacter}, nil
	}
	return nil, fmt.Errorf("unknown character: %s", name)
}
