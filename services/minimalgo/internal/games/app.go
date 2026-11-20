package games

import (
	"context"
	"fmt"
	"myboardgames/minimalgo/internal/models"
	"sync"
	"uuid"
)

type GameServerApp struct {
	tables map[string]models.TableInterface
	server models.ServerAppInterface
	mu     sync.Mutex
}

func (g *GameServerApp) Start(ctx context.Context) {
	fmt.Println("Initializing game server")
}

func (g *GameServerApp) CreateTable() {
	g.mu.Lock()
	defer g.mu.Unlock()
	table := &models.Table{
		ID: uuid.New().String(),
	}
	g.tables[table.ID] = table
}

func NewGameServerApp(ctx context.Context, server models.ServerAppInterface) models.AppInterface {
	return &GameServerApp{
		tables: make(map[string]models.TableInterface),
		server: server,
	}
}
