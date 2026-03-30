package tests

import (
	"testing"

	"github.com/Zadigo/flipseven/internal/logic"
)

func TestGetDeck(t *testing.T) {
	tableLayer := logic.CreateNewTableLayer()
	tableLayer.Layer.GetDeck()

	expected := []int{79, 6, 9}

	if len(tableLayer.Layer.GetTable().CurrentDeck) != logic.SumValues(expected...) {
		t.Errorf("Expected 79 cards, got %d", len(tableLayer.Layer.GetTable().CurrentDeck))
	}
}

func TestFlipCard(t *testing.T) {
	tableLayer := logic.CreateNewTableLayer()
	tableLayer.Layer.GetDeck()

	connection, _, _ := LiveGameWebsocketConnection(t)
	tableLayer.Layer.AddPlayer("Alice", connection)
	card := tableLayer.Layer.FlipCard("Alice", 1)

	if card == nil {
		t.Errorf("Expected a card to be flipped, got nil")
	}
}

func TestAddPlayer(t *testing.T) {
	tableLayer := logic.CreateNewTableLayer()

	connection, _, _ := LiveGameWebsocketConnection(t)
	tableLayer.Layer.AddPlayer("Alice", connection)

	if !tableLayer.Layer.HasPlayer("Alice") {
		t.Errorf("Expected player Alice to be added, but was not found")
	}

	if !(tableLayer.Layer.GetNumberOfPlayers() > 0) {
		t.Errorf("Expected number of players to be greater than 0, got %d", tableLayer.Layer.GetNumberOfPlayers())
	}
}

func TestCalculateAllPoints(t *testing.T) {
	tableLayer := logic.CreateNewTableLayer()

	connection, _, _ := LiveGameWebsocketConnection(t)
	tableLayer.Layer.AddPlayer("Alice", connection)

	player := tableLayer.Layer.GetPlayer("Alice")
	player.Details.Cards = []logic.Card{
		{Value: 2, IsNumber: true},
		{Value: 3, IsNumber: true},
		{Value: 2, IsMultiplier: true},
	}

	total := player.Details.TotalPoints()

	if total != 10 {
		t.Errorf("Expected total points to be 10, got %d", total)
	}

	allPoints := tableLayer.Layer.CalculateAllPoints()

	if allPoints != 10 {
		t.Errorf("Expected total points for all players to be 10, got %d", allPoints)
	}
}

func TestHasPlayer(t *testing.T) {
	tableLayer := logic.CreateNewTableLayer()

	connection, _, _ := LiveGameWebsocketConnection(t)
	tableLayer.Layer.AddPlayer("Alice", connection)

	if !tableLayer.Layer.HasPlayer("Alice") {
		t.Errorf("Expected player Alice to be found, but was not")
	}

	if tableLayer.Layer.HasPlayer("Bob") {
		t.Errorf("Expected player Bob to not be found, but was")
	}
}
