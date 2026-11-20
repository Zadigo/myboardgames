package games

import (
	"context"
	"fmt"
	"myboardgames/minimalgo/internal/models"
	"sync"
)

type GameServerApp struct {
	tables  []models.TableInterface
	clients []models.WebsocketClientInterface
	mu      sync.Mutex
}

func (g *GameServerApp) Start(ctx context.Context) {
	fmt.Println("Initializing game server")
}

func NewGameServerApp(ctx context.Context) models.AppInterface {
	return &GameServerApp{}
}
