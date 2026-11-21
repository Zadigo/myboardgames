package models

type BaseActionsInterface interface {
	AddPlayer(id string) error
	RemovePlayer(id string) error
}

type GameRegistryInterface interface {
	BaseActionsInterface
}

type TableInterface interface {
	BaseActionsInterface
}
