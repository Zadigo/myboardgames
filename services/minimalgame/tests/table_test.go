package tests

import (
	"myboardgames/minimalgame/internal/logic"
	"myboardgames/minimalgame/internal/models"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTable(t *testing.T) {
	t.Run("should create new table", func(t *testing.T) {
		t.Parallel()

		table := logic.NewTable(t.Context())
		assert.NotNil(t, table)
		assert.IsType(t, &logic.Table{}, table)
	})

	// t.Run("should add players", func(t *testing.T) {
	// 	t.Parallel()

	// 	table := logic.NewTable(t.Context())
	// 	err := table.AddPlayer(&models.Player{ID: "1"})

	// 	assert.NoError(t, err)
	// 	assert.Equal(t, 1, table.PlayerCount())

	// 	t.Run("should remove existing player", func(t *testing.T) {
	// 		err := table.RemovePlayer("1")
	// 		assert.NoError(t, err)
	// 		assert.Equal(t, 0, table.PlayerCount())
	// 	})
	// })

	// t.Run("should error if player exists", func(t *testing.T) {
	// 	t.Parallel()

	// 	table := logic.NewTable(t.Context())

	// 	err := table.AddPlayer(&models.Player{ID: "1"})
	// 	err = table.AddPlayer(&models.Player{ID: "1"})

	// 	assert.Error(t, err)
	// 	assert.ErrorIs(t, err, logic.PlayerAlreadyExistsError)
	// 	assert.Equal(t, 1, table.PlayerCount())
	// })

	// t.Run("should run logic on table", func(t *testing.T) {
	// 	t.Parallel()

	// 	table := logic.NewTable(t.Context())

	// 	inst := &models.Instruction{}

	// 	err := table.Execute(t.Context(), inst)
	// 	assert.NoError(t, err)
	// 	assert.NotNil(t, table)
	// 	assert.NotEmpty(t, inst.TableID)
	// })

	// t.Run("should return snapshot", func(t *testing.T) {
	// 	t.Parallel()

	// 	table := logic.NewTable(t.Context())
	// 	snap := table.Snapshot()
	// 	assert.NotNil(t, snap)
	// 	assert.Contains(t, snap, "snapshot")
	// 	t.Log(snap)
	// })
}

func BenchmarkTable(b *testing.B) {
	b.ResetTimer()
	b.SetParallelism(30)

	table := logic.NewTable(b.Context())

	b.RunParallel(func(pb *testing.PB) {
		for b.Loop() {
			err := table.AddPlayer(&models.Player{ID: strconv.Itoa(b.N)})
			if err != nil {
				b.Fatal(err)
			}
		}
	})

}
