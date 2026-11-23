package logic

import (
	"context"
	"encoding/json"
	"maps"
	"myboardgames/minimalgame/internal/models"
	"slices"
	"time"
	"uuid"
)

type Table struct {
	ID                 string                    `json:"id"`
	Direction          bool                      `json:"direction"`
	TimeLimit          int                       `json:"time_limit"`
	Duration           time.Duration             `json:"duration"`
	CurrentRound       int                       `json:"current_index"`
	CurrentPlayer      int                       `json:"current_player"`
	UpdatedAt          time.Time                 `json:"updated_at"`
	CreatedAt          time.Time                 `json:"created_at"`
	IsAcceptingPlayers bool                      `json:"is_accepting_players"`
	IsActive           bool                      `json:"is_active"`
	players            map[string]*models.Player `json:"-"`
	logic              *GameLogic                `json:"-"`
}

func (t *Table) GetPlayers() map[string]*models.Player {
	return t.players
}

func (t *Table) GetPlayer(id string) (*models.Player, error) {
	p, ok := t.players[id]
	if !ok {
		return nil, PlayerNotFoundError
	}
	return p, nil
}

func (t *Table) GetPlayerKeys() []string {
	keys := maps.Keys(t.players)
	return slices.Collect(keys)
}

func (t *Table) ChangeDirection() {
	if t.Direction {
		t.Direction = false
	} else {
		t.Direction = true
	}
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
	table := &Table{
		ID:                 uuid.New().String(),
		players:            make(map[string]*models.Player, 20),
		TimeLimit:          0,
		IsAcceptingPlayers: true,
		IsActive:           true,
		CurrentRound:       0,
		CurrentPlayer:      0,
		UpdatedAt:          time.Now().UTC(),
		CreatedAt:          time.Now().UTC(),
		logic:              &GameLogic{},
	}
	table.Duration = table.CreatedAt.Sub(table.UpdatedAt)
	return table
}
