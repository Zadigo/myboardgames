package tests

import (
	"testing"

	"github.com/Zadigo/splendor/internal/backend"
)

func TestPlayingTable(t *testing.T) {
	table := backend.NewPlayingTable(true)

	t.Run("Should", func(t *testing.T) {
		table.StartGame()
	})

	t.Run("Should", func(t *testing.T) {})
	
	t.Run("Should", func(t *testing.T) {})
	
	t.Run("Should", func(t *testing.T) {})
	
	t.Run("Should", func(t *testing.T) {})
}
