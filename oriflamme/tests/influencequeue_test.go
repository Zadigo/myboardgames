package tests

import (
	"testing"

	"github.com/Zadigo/oriflamme/internal/models"
)

func TestInfluenceQueue(t *testing.T) {
	simpleQueue := models.InfluenceQueue{
		Queue: []*models.Card{
			{Name: "Archer", Color: "Red", Position: 0},
			{Name: "Assassination", Color: "Blue", Position: 1},
			{Name: "Spy", Color: "Blue", Position: 2},
			{Name: "Soldier", Color: "Blue", Position: 3},
		},
	}

	otherCard := &models.Card{Name: "Archer", Color: "Black", Position: 0}
	simpleQueue.AddCardLeft(otherCard)

	if simpleQueue.NumberOfCards() != 5 {
		t.Errorf("Expected 5 cards in the queue, got %d", simpleQueue.NumberOfCards())
	}

	otherCard = &models.Card{Name: "Spy", Color: "Black", Position: 4}
	simpleQueue.AddCardRight(otherCard)

	if simpleQueue.NumberOfCards() != 6 {
		t.Errorf("Expected 6 cards in the queue, got %d", simpleQueue.NumberOfCards())
	}

	simpleQueue.RemoveCardAtPosition(1)
	if !simpleQueue.Queue[1].IsRemoved {
		t.Error("Expected card at position 1 to be marked as removed")
	}
}
