package logic

import (
	"context"
	"encoding/json"
	"myboardgames/minimalgame/internal/models"
	"time"
	"uuid"
)

type Table struct {
	ID            string                    `json:"id"`
	TimeLimit     int                       `json:"time_limit"`
	Duration      time.Duration             `json:"duration"`
	CurrentRound  int                       `json:"current_index"`
	CurrentPlayer int                       `json:"current_player"`
	UpdatedAt     time.Time                 `json:"updated_at"`
	CreatedAt     time.Time                 `json:"created_at"`
	players       map[string]*models.Player `json:"-"`
	logic         *GameLogic
}

func (t *Table) GetID() string {
	return t.ID
}

func (t *Table) AddPlayer(player *models.Player) error {
	if _, ok := t.players[player.ID]; ok {
		return PlayerAlreadyExistsError
	}
	t.players[player.ID] = &models.Player{ID: player.ID}
	return nil
}

func (t *Table) RemovePlayer(id string) error {
	delete(t.players, id)
	return nil
}

func (t *Table) PlayerCount() int {
	return len(t.players)
}

func (t *Table) Snapshot() map[string]any {
	if b, err := json.Marshal(t); err == nil {
		return map[string]any{"snapshot": string(b)}
	}
	return nil
}

func (t *Table) Execute(ctx context.Context, instruction *models.Instruction) error {
	instruction.TableID = t.ID
	return t.logic.RunMove(ctx, t, instruction)
}

func NewTable() models.TableInterface {
	return &Table{
		ID:      uuid.New().String(),
		players: make(map[string]*models.Player, 20),
		logic:   &GameLogic{},
	}
}
