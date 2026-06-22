package tests

import (
	"testing"

	"github.com/Zadigo/splendor/internal/backend"
)

func TestReadJson(t *testing.T) {
	data := backend.ReadJsonFile("/Volumes/Coding/Projects/Websites/myboardgames/splendor/internal/data/marvel_cards.json")

	t.Run("Should be able to read file", func(t *testing.T) {
		t.Log(data)
	})
}
