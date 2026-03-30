package tests

import (
	"testing"

	"github.com/Zadigo/flipseven/internal/logic"
)

func TestDeck(t *testing.T) {
	cards := logic.GetDeck()

	totalCards := []int{79, 6, 9}

	if len(cards) != logic.SumValues(totalCards...) {
		t.Errorf("Expected 79 cards, got %d", len(cards))
	}
}
