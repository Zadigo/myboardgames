package models

type BaseActionsInterface interface {
	RemovePlayer(id string) error
}

type GameRegistryInterface interface {
	BaseActionsInterface
	AddPlayer(id string) *Player
	AddTable(table TableInterface) error
	GetTable(id string) (TableInterface, error)
	RemoveTable(table TableInterface) error
}

type TableInterface interface {
	BaseActionsInterface
	AddPlayer(*Player) error
	GetID() string
	Execute(*Instruction) error
}
