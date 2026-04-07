package tests

import (
	"testing"

	"github.com/Zadigo/flipseven2/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestBaseRegistry(t *testing.T) {
	type testCase struct {
		expected               any
		expectedNumberOfTables int
		error                  error
	}

	testCases := []testCase{
		{
			expected:               true,
			expectedNumberOfTables: 1,
			error:                  nil,
		},
	}

	baseRegistry := models.CreateBaseRegistry()
	assert.NotNil(t, baseRegistry)

	for _, tc := range testCases {
		t.Run("Read, write and delete of player's table", func(t *testing.T) {
			playersTable := models.PlayersTable{
				Layer: &models.TableLayer{
					IsStarted: false,
				},
			}

			baseRegistry.Set(&playersTable)

			tableId := playersTable.Layer.GetUuid()
			table, state := baseRegistry.Get(tableId)

			assert.True(t, state)
			assert.Equal(t, tc.expected, state)
			assert.Equal(t, &playersTable, table)

			// With the same tableId, GetOrCreate should return the same table
			// and not create a new one
			table2 := baseRegistry.GetOrCreate(tableId, models.CreatePlayersTable)
			assert.Equal(t, &playersTable, table2)
			assert.Equal(t, baseRegistry.NumberOfTables(), tc.expectedNumberOfTables)

			// Delete the table and check that it is removed from the registry
			baseRegistry.Delete(tableId)
			table, state = baseRegistry.Get(tableId)
			assert.False(t, state)
		})
	}
}
