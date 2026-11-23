package models

import "context"

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
	GetPlayers() map[string]*Player
	GetPlayer(id string) (*Player, error)
	GetPlayerKeys() []string
	AddPlayer(*Player) error
	GetID() string
	Snapshot() map[string]any
	PlayerCount() int
	ChangeDirection()
	// Execute is the main entrypoint used by
	// by a gRPC call in order to execute a
	// move in the game
	Execute(ctx context.Context, instruction *Instruction) error
}
