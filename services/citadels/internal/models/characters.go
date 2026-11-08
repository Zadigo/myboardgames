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
	IsBewitched() bool
	GetLevel() int
	Kill()
	SetPlayer(player PlayerInterface)
	GetPlayer() PlayerInterface
}

type BaseCharacter struct {
	ID        string
	Level     int
	Name      string
	Image     string
	Killed    bool
	Bewitched bool
	Player    PlayerInterface
}
