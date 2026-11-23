package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"maps"
	"myboardgames/minimalgame/internal/backend"
	"myboardgames/minimalgame/internal/models"
	"slices"
	"time"
	"uuid"

	"github.com/redis/go-redis/v9"
)

type Table struct {
	ID                 string                    `json:"id"`
	Creator            string                    `json:"creator"` // The player that created the table
	Direction          bool                      `json:"direction"`
	TimeLimit          int                       `json:"time_limit"`
	Duration           time.Duration             `json:"duration"`
	CurrentRound       int                       `json:"current_index"`
	CurrentPlayer      int                       `json:"current_player"`
	UpdatedAt          time.Time                 `json:"updated_at"`
	CreatedAt          time.Time                 `json:"created_at"`
	IsAcceptingPlayers bool                      `json:"is_accepting_players"`
	IsActive           bool                      `json:"is_active"`
	Sub                string                    `json:"sub"` // The Redis subscription channel used to send messages to all the players
	redisDb            *redis.Client             `json:"-"`
	players            map[string]*models.Player `json:"-"`
	logic              *GameLogic                `json:"-"`
	diffusion          chan *models.Instruction  `json:"-"`
	cancel             context.CancelFunc        `json:"-"`
}

func (t *Table) publishLoop(ctx context.Context) {
	defer t.cancel()
	defer func() { // a panic here would kill the whole process
		if r := recover(); r != nil {
			slog.Error("table publisher panic", "table", t.ID, "panic", r)
		}
	}()

	redisDb, err := backend.CreateRedisConnection(ctx)
	if err != nil {
		return
	}

	defer func() {
		err := redisDb.Close()
		if err != nil {
			slog.Error("Failed to close redis connection")
		}
	}()

	slog.Info("Started subscription channel", "id", t.Sub)

	for {
		select {
		case msg := <-t.diffusion:
			payload, err := json.Marshal(msg)
			if err != nil {
				slog.Error("marshal event", "table", t.ID, "err", err)
				continue
			}

			pubCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
			err = t.redisDb.Publish(pubCtx, t.Sub, payload).Err()
			cancel()

			if err != nil {
				slog.Error("publish event", "table", t.ID, "err", err)
			}
		case <-ctx.Done():
			timeoutCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			_ = t.redisDb.Publish(timeoutCtx, t.Sub, `{"action":"table_closed"}`).Err()
			cancel()
			return
		}
	}
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

func (t *Table) PublishSubscriptionMessage(instruction *models.Instruction) {
	t.diffusion <- instruction
}

func (t *Table) GetSubscriptionKey() string {
	return t.Sub
}

func (t *Table) Execute(ctx context.Context, instruction *models.Instruction) error {
	instruction.TableID = t.ID
	return t.logic.RunMove(ctx, t, instruction)
}

func NewTable(redisDb *redis.Client) models.TableInterface {
	ctx, cancel := context.WithCancel(context.Background())

	table := &Table{
		ID:                 uuid.New().String(),
		Creator:            "",
		TimeLimit:          0,
		IsAcceptingPlayers: true,
		IsActive:           true,
		CurrentRound:       0,
		CurrentPlayer:      0,
		Sub:                "",
		UpdatedAt:          time.Now().UTC(),
		CreatedAt:          time.Now().UTC(),
		cancel:             cancel,
		players:            make(map[string]*models.Player, 20),
		logic:              &GameLogic{},
		diffusion:          make(chan *models.Instruction, 100),
	}

	table.Sub = "table:" + table.ID
	table.Duration = table.CreatedAt.Sub(table.UpdatedAt)
	table.redisDb = redisDb

	// Start a Redis pub-sub that will update the table for
	// every player that will be listening for changes on
	// the table
	go table.publishLoop(ctx)

	table.PublishSubscriptionMessage(&models.Instruction{
		Action: models.TableEventConfigured,
		MetaData: models.InstructionMetaData{
			Description: fmt.Sprintf("Table with ID %s was created", table.ID),
		},
	})

	return table
}
