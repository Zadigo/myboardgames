package tests

import (
	"testing"

	"github.com/Zadigo/oriflamme/internal/backend"
)

func constructQueue(t *testing.T) (*backend.InfluenceQueue, *backend.WebsocketClient, *backend.WebsocketClient) {
	t.Helper()

	ownerOne := &backend.WebsocketClient{Username: "Player1"}
	ownerTwo := &backend.WebsocketClient{Username: "Player2"}

	simpleQueue := backend.InfluenceQueue{
		ResolutionIndex: -1,
		Queue: []backend.CardInterface{
			&backend.ArcherCard{BaseCard: backend.NewBaseCard("Archer", "Character", ownerOne, "Red")},
			&backend.AssassinationCard{BaseCard: backend.NewBaseCard("Assassination", "Character", ownerTwo, "Blue")},
			&backend.SpyCard{BaseCard: backend.NewBaseCard("Spy", "Character", ownerTwo, "Red")},
			&backend.SoldierCard{BaseCard: backend.NewBaseCard("Soldier", "Character", ownerTwo, "Blue")},
		},
	}
	return &simpleQueue, ownerOne, ownerTwo
}

// func TestInfluenceQueue(t *testing.T) {
// 	simpleQueue, ownerOne, ownerTwo := constructQueue()

// 	otherCard := &backend.ArcherCard{BaseCard: &backend.BaseCard{Uuid: uuid.NewString(), Name: "Archer", Color: "Red", PositionInQueue: 0, Owner: ownerOne}}
// 	simpleQueue.AddCardLeft(otherCard)

// 	if simpleQueue.NumberOfCards() != 5 {
// 		t.Errorf("Expected 5 cards in the queue, got %d", simpleQueue.NumberOfCards())
// 	}

// 	otherCard = &backend.Card{Uuid: uuid.NewString(), Name: "Spy", Color: "Blue", PositionInQueue: 4, Owner: ownerTwo}
// 	simpleQueue.AddCardRight(otherCard)

// 	if simpleQueue.NumberOfCards() != 6 {
// 		t.Errorf("Expected 6 cards in the queue, got %d", simpleQueue.NumberOfCards())
// 	}

// 	simpleQueue.RemoveCardAtPosition(1)
// 	if !simpleQueue.Queue[1].IsRemoved {
// 		t.Error("Expected card at position 1 to be marked as removed")
// 	}
// }

// func TestInfluenceQueueResolution(t *testing.T) {
// 	simpleQueue, _, _ := constructQueue()
// 	simpleQueue.Resolve()

// 	if simpleQueue.ResolutionIndex != 0 {
// 		t.Errorf("Expected resolution index to be 0, got %d", simpleQueue.ResolutionIndex)
// 	}

// 	card, _ := simpleQueue.GetCurrentCard()
// 	if card == nil {
// 		t.Error("Expected current card to be non-nil")
// 	} else if card.Name != "Archer" {
// 		t.Errorf("Expected current card to be Archer, got %s", card.Name)
// 	}

// 	result := simpleQueue.ApplyEffect(card)
// 	if result {
// 		t.Error("Expected effect to be applied successfully")
// 	}

// 	card.Reveal(simpleQueue)
// 	result = simpleQueue.ApplyEffect(card)
// 	if !result {
// 		t.Error("Expected effect to be applied successfully")
// 	}
// }
