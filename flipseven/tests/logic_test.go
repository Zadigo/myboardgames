package tests

import (
	"testing"

	"github.com/Zadigo/flipseven/internal"
	"github.com/Zadigo/flipseven/internal/logic"
)

func TestCalcualtePlayerScores(t *testing.T) {
	player := &internal.ConnectedPlayer{
		Details: internal.Player{
			Cards: []internal.Card{
				{Value: 1, IsNumber: true},
				{Value: 2, IsNumber: true},
				{Value: 3, IsNumber: true},
			},
		},
	}

	logic.CalcualtePlayerScores(player)

	if player.Details.Score != 6 {
		t.Errorf("Expected score to be 6, but got %d", player.Details.Score)
	}
}

func TestCalcualtePlayerScoresWithMultiplier(t *testing.T) {
	player := &internal.ConnectedPlayer{
		Details: internal.Player{
			Cards: []internal.Card{
				{Value: 2, IsNumber: true},
				{Value: 2, IsMultiplier: true},
			},
		},
	}

	logic.CalcualtePlayerScores(player)

	if player.Details.Score != 4 {
		t.Errorf("Expected score to be 4, but got %d", player.Details.Score)
	}
}

func TestCalcualtePlayerScoresWithBonus(t *testing.T) {
	player := &internal.ConnectedPlayer{
		Details: internal.Player{
			Cards: []internal.Card{
				{Value: 2, IsNumber: true},
				{Value: 2, IsBonus: true},
			},
		},
	}

	logic.CalcualtePlayerScores(player)

	if player.Details.Score != 4 {
		t.Errorf("Expected score to be 4, but got %d", player.Details.Score)
	}
}
