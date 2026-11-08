package models

type CharacterOptions struct {
	Target      CharacterInterface
	CurrentGame GameInterface
}

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
