package logic

import (
	"context"
	"myboardgames/minimalgame/internal/models"
)

// GameLogic holds the core game rules. This is
// the section you should modify for game logic
type GameLogic struct{}

func (g *GameLogic) preMove(table *Table, instruction *models.Instruction) {
	// Pre-move logic here
}

func (g *GameLogic) postMove(table *Table, instruction *models.Instruction) {
	// Post-move logic here
	// table.PublishSubscriptionMessage(instruction)
}

func (g *GameLogic) RunMove(ctx context.Context, table *Table, instruction *models.Instruction) error {
	// Code to execute after the move is done
	defer g.postMove(table, instruction)

	// Code that execute before the move is done
	go g.preMove(table, instruction)

	// Implement move logic here

	return nil
}
