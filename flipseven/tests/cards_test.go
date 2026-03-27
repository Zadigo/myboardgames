package tests

import (
	"testing"

	"github.com/Zadigo/flipseven/cards"
	"github.com/Zadigo/flipseven/logic"
)

func TestBaseCardGenerator(t *testing.T) {
	cards := cards.GetDeck()

	totalCards := []int{79, 6, 9}

	if len(cards) != logic.SumValues(totalCards) {
		t.Errorf("Expected 79 cards, got %d", len(cards))
	}
}
