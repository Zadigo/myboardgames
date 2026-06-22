package logic

import "context"

type GameEngine struct {
	ctx context.Context
}

func NewGameEngine(ctx context.Context) *GameEngine {
	return &GameEngine{ctx: ctx}
}
