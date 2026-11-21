package logic

import (
	"context"
	"log/slog"
	"myboardgames/minimalgo/internal/models"
	"time"
	"uuid"
)

type ExampleTable struct {
	models.Table
}

func (e *ExampleTable) Run(ctx context.Context, options *models.TableOptions) {
	slog.Info("Running table", "tableId", e.ID)

	// db := &TableRedis{client: e.}

	for {
		select {
		case event, ok := <-e.GetInbox():
			if !ok {
				return
			}

			switch event.Details.EventName {
			case models.TableEventStealCard:
				if event.Client != nil {
					event.Client.GetConn().WriteMessage(1, []byte("Stole card"))
				}

			}
		case <-ctx.Done():
			timeoutCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
			defer cancel()

			slog.Error(timeoutCtx.Err().Error())
		}
	}
}

func NewExampleTable() *ExampleTable {
	table := &ExampleTable{
		ID:        uuid.New().String(),
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	table.Setup(nil)
	return table
}
