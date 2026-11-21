package games

import (
	"context"
	"fmt"
	"log/slog"
	"myboardgames/minimalgo/internal/logic"
	"myboardgames/minimalgo/internal/models"
	"sync"
)

type GameServerApp struct {
	tables map[string]models.TableInterface
	server models.ServerAppInterface
	mu     sync.RWMutex
}

func (g *GameServerApp) getTable(id string) (models.TableInterface, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	table, ok := g.tables[id]
	if !ok {
		return nil, fmt.Errorf("Table not found")
	}
	return table, nil
}

// Start listens for messages coming from the Chi HTTP
// handler and executes actions based on messages
func (g *GameServerApp) Start(ctx context.Context) error {
	slog.Info("Initializing game server")

	in := g.server.GetInstructions()
	for {
		select {
		case msg, ok := <-in:
			if !ok {
				return nil
			}

			if msg.Sender != models.HttpAppSender {
				continue
			}

			switch msg.Action {
			case models.CreateTable:
				table := g.CreateTable(ctx, msg.TableOptions)

				g.server.SendInstructionResponse(ctx, models.InternalInstructionResponse{
					Table:      table,
					StrMessage: "Table created successfully",
				})
			case models.TableEventStealCard:
				if table, err := g.getTable(msg.TableID); err == nil {
					// handle event
				} else {
					fmt.Printf("Got table %s", table.GetID())
				}
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// CreateTable creates a goroutine for the specific table and returns it to the caller
func (g *GameServerApp) CreateTable(ctx context.Context, options *models.TableOptions) models.TableInterface {
	table := logic.NewExampleTable()
	table.Setup(options.Game)

	g.mu.Lock()
	g.tables[table.ID] = table
	g.mu.Unlock()

	go func() {
		table.Run(ctx, options)

		g.mu.Lock()
		delete(g.tables, table.ID) // cleanup when the table ends
		g.mu.Unlock()
	}()

	return table
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
