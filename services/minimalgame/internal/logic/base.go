package logic

import (
	"context"
	"myboardgames/minimalgame/internal/models"
)

// GameLogic holds the core game rules. This is
// the section you should modify for game logic
type GameLogic struct{}

func (g *GameLogic) preMove(table *Table, options *models.Instruction) {
	// Pre-move logic here
}

func (g *GameLogic) postMove(table *Table, options *models.Instruction) {
	// Post-move logic here
}

func (g *GameLogic) RunMove(ctx context.Context, table *Table, options *models.Instruction) error {
	// Code to execute after the move is done
	defer g.postMove(table, options)

	// Code that execute before the move is done
	go g.preMove(table, options)

	// Implement move logic here

	return nil
}
