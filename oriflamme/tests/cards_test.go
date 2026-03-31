package tests

import (
	"testing"

	"github.com/Zadigo/oriflamme/internal/models"
)

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

	result = card.ApplySoldier(&models.InfluenceQueue{})
	if result {
		t.Error("Expected ApplySoldier to return false")
	}

	result = card.ApplySpy(&models.InfluenceQueue{}, true)
	if result {
		t.Error("Expected ApplySpy to return false")
	}
}
