package tests

import (
	"testing"

	"github.com/Zadigo/oriflamme/internal/models"
)

func cardConstructor() (*models.Player, []models.Card) {
	alice := &models.Player{Username: "Alice", Tokens: 1}
	blueCards := models.CreateBlueCards(alice)

	return alice, blueCards
}

func TestCard(t *testing.T) {
	card := models.Card{
		IsRemoved:   false,
		IsSelected:  true,
		IsDiscarded: false,
	}

	card.Reveal()
	if !card.IsRevealed {
		t.Error("Card is not revealed")
	}

	card.Discard()
	if !card.IsDiscarded {
		t.Error("Card is not discarded")
	}

	// No tokens
	card.AddToken(3)
	if card.Tokens != 3 {
		t.Errorf("Expected 3 tokens, got %d", card.Tokens)
	}

	card.RemoveToken(2)
	if card.Tokens != 1 {
		t.Errorf("Expected 1 token, got %d", card.Tokens)
	}

	result := card.ApplyAssassination(&models.InfluenceQueue{}, 1)
	if result {
		t.Error("Expected ApplyAssassination to return false")
	}

	result = card.ApplyArcher(&models.InfluenceQueue{}, true)
	if result {
		t.Error("Expected ApplyArcher to return false")
	}

	result = card.ApplySoldier(&models.InfluenceQueue{}, "after")
	if result {
		t.Error("Expected ApplySoldier to return false")
	}

	result = card.ApplySpy(&models.InfluenceQueue{}, true)
	if result {
		t.Error("Expected ApplySpy to return false")
	}
}

func TestArcherCard(t *testing.T) {
	alice := &models.Player{Username: "Alice"}
	blueCards := models.CreateBlueCards(alice)

	simpleQueue := models.InfluenceQueue{
		ResolutionIndex: -1,
		Queue: []*models.Card{
			{Uuid: "1", Name: "Test Card"},
			&blueCards[0],
			{Uuid: "1", Name: "Test Card"},
		},
	}

	simpleQueue.ResolutionIndex = 1
	currentCard := simpleQueue.GetCurrentCard()

	currentCard.Reveal()
	// Eliminate the first card in the queue
	result := currentCard.ApplyArcher(&simpleQueue, true)

	if !result {
		t.Error("Expected effect to be applied successfully")
	}

	if !simpleQueue.Queue[0].IsRemoved {
		t.Error("Expected first card to be removed")
	}

	// Eliminate the last card in the queue
	result = currentCard.ApplyArcher(&simpleQueue, false)

	if !result {
		t.Error("Expected effect to be applied successfully")
	}

	if !simpleQueue.Queue[len(simpleQueue.Queue)-1].IsRemoved {
		t.Error("Expected last card to be removed")
	}

	if currentCard.Owner.Tokens != 2 {
		t.Errorf("Expected player to gain 1 token per elimination, got %d", currentCard.Owner.Tokens)
	}
}

func TestSoldierCard(t *testing.T) {
	alice := &models.Player{Username: "Alice"}
	blueCards := models.CreateBlueCards(alice)

	simpleQueue := models.InfluenceQueue{
		ResolutionIndex: -1,
		Queue: []*models.Card{
			{Uuid: "1", Name: "Test Card"},
			&blueCards[1],
			{Uuid: "1", Name: "Test Card"},
		},
	}

	simpleQueue.ResolutionIndex = 1
	currentCard := simpleQueue.GetCurrentCard()

	currentCard.Reveal()

	// Eliminate the card immediately before the soldier
	result := currentCard.ApplySoldier(&simpleQueue, "before")

	if !result {
		t.Error("Expected effect to be applied successfully")
	}

	if !simpleQueue.Queue[0].IsRemoved {
		t.Error("Expected first card to be removed")
	}

	// Eliminate the card immediately after the soldier
	result = currentCard.ApplySoldier(&simpleQueue, "after")

	if !result {
		t.Error("Expected effect to be applied successfully")
	}

	if !simpleQueue.Queue[len(simpleQueue.Queue)-1].IsRemoved {
		t.Error("Expected last card to be removed")
	}

	if currentCard.Owner.Tokens != 2 {
		t.Errorf("Expected player to gain 1 token per elimination, got %d", currentCard.Owner.Tokens)
	}
}

func TestSoldierPositionFirst(t *testing.T) {
	_, blueCards := cardConstructor()

	simpleQueue := models.InfluenceQueue{
		ResolutionIndex: -1,
		Queue: []*models.Card{
			&blueCards[1],
			{Uuid: "1", Name: "Test Card"},
			{Uuid: "1", Name: "Test Card"},
		},
	}

	simpleQueue.ResolutionIndex = 0
	currentCard := simpleQueue.GetCurrentCard()

	currentCard.Reveal()

	// Attempt to eliminate the card before the soldier, which should fail since the soldier is at the front of the queue
	result := currentCard.ApplySoldier(&simpleQueue, "none")

	if !result {
		t.Error("Expected effect to be applied successfully")
	}
}

func TestSoldierPositionLast(t *testing.T) {
	_, blueCards := cardConstructor()

	simpleQueue := models.InfluenceQueue{
		ResolutionIndex: -1,
		Queue: []*models.Card{
			{Uuid: "1", Name: "Test Card"},
			{Uuid: "1", Name: "Test Card"},
			&blueCards[1],
		},
	}

	simpleQueue.ResolutionIndex = 2
	currentCard := simpleQueue.GetCurrentCard()

	currentCard.Reveal()

	// Attempt to eliminate the card before the soldier, which should fail since the soldier is at the front of the queue
	result := currentCard.ApplySoldier(&simpleQueue, "none")

	if !result {
		t.Error("Expected effect to be applied successfully")
	}
}

func TestSpyCard(t *testing.T) {
	_, blueCards := cardConstructor()

	simpleQueue := models.InfluenceQueue{
		ResolutionIndex: -1,
		Queue: []*models.Card{
			{Uuid: "1", Name: "Test Card", Tokens: 1, Owner: &models.Player{Username: "Pierre"}},
			&blueCards[2], // Spy
			{Uuid: "1", Name: "Test Card", Tokens: 1, Owner: &models.Player{Username: "Pauline"}},
		},
	}

	simpleQueue.ResolutionIndex = 1
	currentCard := simpleQueue.GetCurrentCard()

	currentCard.Reveal()

	result := currentCard.ApplySpy(&simpleQueue, true)

	if !result {
		t.Error("Expected effect to be applied successfully")
	}

	prevCard := simpleQueue.Queue[0]

	if prevCard.Owner.Tokens != 0 {
		t.Errorf("Expected previous card to lose 1 token, got %d", prevCard.Tokens)
	}
}

func TestConspiracyCard(t *testing.T) {
	_, blueCards := cardConstructor()

	simpleQueue := models.InfluenceQueue{
		ResolutionIndex: -1,
		Queue: []*models.Card{
			&blueCards[8],
		},
	}

	simpleQueue.ResolutionIndex = 0
	currentCard := simpleQueue.GetCurrentCard()
	currentCard.Tokens = 2

	currentCard.Reveal()
	result := currentCard.ApplyConspiracy(&simpleQueue)

	if !result {
		t.Error("Expected effect to be applied successfully")
	}

	if currentCard.Owner.Tokens != 5 {
		t.Errorf("Expected player to gain 4 tokens for a total of 5, got %d", currentCard.Owner.Tokens)
	}
}
