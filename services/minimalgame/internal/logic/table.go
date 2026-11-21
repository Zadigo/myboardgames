package logic

import (
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

func (t *Table) Execute(instruction *models.Instruction) error {
	// Implement custom logic here for the current game
	return nil
}

func NewTable() models.TableInterface {
	return &Table{
		ID:      uuid.New().String(),
		players: make(map[string]*models.Player, 20),
	}
}
