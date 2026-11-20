package tests

import (
	"context"
	"myboardgames/minimalgo/internal/games"
	"myboardgames/minimalgo/internal/logic"
	"myboardgames/minimalgo/internal/models"
	"myboardgames/minimalgo/tests/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGameServer(t *testing.T) {
	t.Run("should instantiate game server", func(t *testing.T) {
		tCtx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
		defer cancel()

		ms, ch := utils.HelperMainServer(t)
		s := games.NewGameServerApp(tCtx, ms)
		s.Start(tCtx)

		ch <- models.InternalMessage{Sender: "test_client"}
		msg, ok := <-ch
		if ok {
			assert.True(t, msg.Sender == "test_client")
		}
	})

	t.Run("should create table", func(t *testing.T) {
		t.Parallel()

		tCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		ms, _ := utils.HelperMainServer(t)
		s := games.NewGameServerApp(tCtx, ms)
		s.CreateTable(tCtx, nil)

		assert.Len(t, s.GetTables(), 1)
	})
}

func TestCreateTable(t *testing.T) {
	t.Run("should be able to run a table", func(t *testing.T) {
		tCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		ms, _ := utils.HelperMainServer(t)
		s := games.NewGameServerApp(tCtx, ms)
		table := s.CreateTable(tCtx, nil)

		assert.NotNil(t, table)
		assert.IsType(t, table, &logic.ExampleTable{})
		assert.Len(t, s.GetTables(), 1)
	})
}
