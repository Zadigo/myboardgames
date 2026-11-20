package games

import (
	"context"
	"fmt"
	"log/slog"
	"myboardgames/minimalgo/internal/models"
	"sync"
	"uuid"
)

type GameServerApp struct {
	tables map[string]models.TableInterface
	server models.ServerAppInterface
	mu     sync.Mutex
}

// Start listens for messages coming from the Chi HTTP
// handler and executes actions based on messages
func (g *GameServerApp) Start(ctx context.Context) {
	slog.Info("Initializing game server")

	in := g.server.GetInternalChannel()
	for {
		select {
		case msg, ok := <-in:
			if !ok {
				return
			}

			if msg.Sender != models.HttpAppSender {
				continue
			}
			fmt.Println(msg)
			// switch msg.StrMessage {
			// case models.CreateTable:
			// 	_, table := g.CreateTable(nil)
			// 	g.server.WriteInternalMessage(ctx, models.InternalMessage{
			// 		Table: table,
			// 	})
			// }
		case <-ctx.Done():
			return
		}
	}
}

func (g *GameServerApp) CreateTable(game models.GameInterface) (string, models.TableInterface) {
	g.mu.Lock()
	defer g.mu.Unlock()

	table := &models.Table{ID: uuid.New().String()}
	table.SetGame(game)

	g.tables[table.ID] = table
	return table.ID, table
}

func (g *GameServerApp) GetTables() []models.TableInterface {
	values := []models.TableInterface{}
	for _, v := range g.tables {
		values = append(values, v)
	}
	return values
}

func NewGameServerApp(ctx context.Context, server models.ServerAppInterface) models.GameServerInterface {
	return &GameServerApp{
		tables: make(map[string]models.TableInterface),
		server: server,
	}
}
