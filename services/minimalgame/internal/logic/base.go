package logic

import (
	"context"
	"myboardgames/minimalgame/internal/models"
)

// GameLogic holds the core game rules. This is
// the section you should modify for game logic
type GameLogic struct{}

// RunMove executes a move in the game
func (g *GameLogic) RunMove(ctx context.Context, table *Table, options *models.Instruction) error {
	// Implement move logic here
	return nil
}
