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

func (e *ExampleTable) Run(ctx context.Context) {
	// db := &TableRedis{client: e.}
	for {
		select {
		case instruction, ok := <-e.GetInbox():
			if !ok {
				return
			}

			switch instruction.Details.EventName {
			case models.TableEventIdle:
				// Do something here

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
