package models

type BaseActionsInterface interface {
	AddPlayer(id string) error
	RemovePlayer(id string) error
}

type GameRegistryInterface interface {
	BaseActionsInterface
	AddTable(table TableInterface) error
	GetTable(id string) (TableInterface, error)
	RemoveTable(table TableInterface) error
}

type TableInterface interface {
	BaseActionsInterface
	GetID() string
	Execute(*Instruction) error
}
