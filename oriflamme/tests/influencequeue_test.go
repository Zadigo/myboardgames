package tests

import (
	"testing"

	"github.com/Zadigo/oriflamme/internal/models"
	"github.com/google/uuid"
)

func constructQueuee() (models.InfluenceQueue, *models.Player, *models.Player) {
	ownerOne := &models.Player{Username: "Player1"}
	ownerTwo := &models.Player{Username: "Player2"}

	simpleQueue := models.InfluenceQueue{
		ResolutionIndex: -1,
		Queue: []*models.Card{
			{Uuid: uuid.NewString(), Name: "Archer", Color: "Red", Position: 0, Owner: ownerOne},
			{Uuid: uuid.NewString(), Name: "Assassination", Color: "Blue", Position: 1, Owner: ownerTwo},
			{Uuid: uuid.NewString(), Name: "Spy", Color: "Red", Position: 2, Owner: ownerTwo},
			{Uuid: uuid.NewString(), Name: "Soldier", Color: "Blue", Position: 3, Owner: ownerTwo},
		},
	}
	return simpleQueue, ownerOne, ownerTwo
}

func TestInfluenceQueue(t *testing.T) {
	simpleQueue, ownerOne, ownerTwo := constructQueuee()

	otherCard := &models.Card{Uuid: uuid.NewString(), Name: "Archer", Color: "Red", Position: 0, Owner: ownerOne}
	simpleQueue.AddCardLeft(otherCard)

	if simpleQueue.NumberOfCards() != 5 {
		t.Errorf("Expected 5 cards in the queue, got %d", simpleQueue.NumberOfCards())
	}

	otherCard = &models.Card{Uuid: uuid.NewString(), Name: "Spy", Color: "Blue", Position: 4, Owner: ownerTwo}
	simpleQueue.AddCardRight(otherCard)

	if simpleQueue.NumberOfCards() != 6 {
		t.Errorf("Expected 6 cards in the queue, got %d", simpleQueue.NumberOfCards())
	}

	simpleQueue.RemoveCardAtPosition(1)
	if !simpleQueue.Queue[1].IsRemoved {
		t.Error("Expected card at position 1 to be marked as removed")
	}
}

func TestInfluenceQueueResolution(t *testing.T) {
	simpleQueue, _, _ := constructQueuee()
	simpleQueue.Resolve()

	if simpleQueue.ResolutionIndex != 0 {
		t.Errorf("Expected resolution index to be 0, got %d", simpleQueue.ResolutionIndex)
	}

	card := simpleQueue.GetCurrentCard()
	if card == nil {
		t.Error("Expected current card to be non-nil")
	} else if card.Name != "Archer" {
		t.Errorf("Expected current card to be Archer, got %s", card.Name)
	}

	result := simpleQueue.ApplyEffect(card)
	if result {
		t.Error("Expected effect to be applied successfully")
	}

	card.Reveal()
	result = simpleQueue.ApplyEffect(card)
	if !result {
		t.Error("Expected effect to be applied successfully")
	}
}
