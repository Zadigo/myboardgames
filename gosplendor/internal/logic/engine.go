package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type GameEngineInterface interface {
	AddPlayingTable(table PlayingTableInterface) error
	GetPlayingTable(tableUuid string) (PlayingTableInterface, bool)
}

type GameEngine struct {
	ctx           context.Context
	redisClient   *redis.Client
	playingTables map[string]PlayingTableInterface
}

func (ge *GameEngine) AddPlayingTable(table PlayingTableInterface) error {
	ge.playingTables[table.GetUuid()] = table
	return nil
}

func (ge *GameEngine) GetPlayingTable(tableUuid string) (PlayingTableInterface, bool) {
	table, exists := ge.playingTables[tableUuid]
	return table, exists
}

func NewGameEngine(ctx context.Context, redisClient *redis.Client) GameEngineInterface {
	return &GameEngine{
		ctx:           ctx,
		redisClient:   redisClient,
		playingTables: make(map[string]PlayingTableInterface),
	}
}
