package models

import "slices"

type InfluenceQueue struct {
	// The influence queue is a list of cards that have been played but have not yet resolved.
	Queue []*Card
	// Each card is resolved one by one in an ascending order of their position in the queue.
	// This index tracks which card is currently being resolved. The inde is initialized to
	// -1, meaning that no card has been resolved yet.
	ResolutionIndex int
}

func (queue *InfluenceQueue) AddCardLeft(card *Card) {
	queue.Queue = slices.Insert(queue.Queue, 0, card)

	for i := range queue.Queue {
		queue.Queue[i].Position = i
	}
}

func (queue *InfluenceQueue) AddCardRight(card *Card) {
	queue.Queue = append(queue.Queue, card)

	for i := range queue.Queue {
		queue.Queue[i].Position = i
	}
}

// Remove a card at a specific position in the queue. The card is simply
// marked as removed and will be skipped during resolution.
func (queue *InfluenceQueue) RemoveCardAtPosition(index int) *Card {
	card := queue.Queue[index]
	card.IsRemoved = true
	return card
}

func (queue *InfluenceQueue) StackCard(index int) *Card {
	return nil
}

func (queue *InfluenceQueue) NumberOfCards() int {
	return len(queue.Queue)
}

func (queue *InfluenceQueue) Resolve() {
	if queue.NumberOfCards() == 0 {
		return
	}

	if queue.ResolutionIndex >= queue.NumberOfCards() {
		return
	}

	queue.ResolutionIndex++
}

// Get the current card being resolved. If the resolution index
// is out of bounds, return nil.
func (queue *InfluenceQueue) GetCurrentCard() *Card {
	if queue.NumberOfCards() == 0 {
		return nil
	}

	if queue.ResolutionIndex == -1 {
		return nil
	}

	if queue.ResolutionIndex < queue.NumberOfCards() {
		return queue.Queue[queue.ResolutionIndex]
	}

	return nil
}

// Apply the effect of the current card being resolved. This function should be called
// after calling Resolve() to apply the effect of the current card. It returns true if 
// the effect was successfully applied, and false otherwise.
// Deprecated: use the card effect directly
func (queue *InfluenceQueue) ApplyEffect(card *Card) bool {
	if card.IsRevealed {
		if card.Name == "Archer" {
			return card.ApplyArcher(queue, false)
		}
	}
	return false
}

// Helper function to find the index of a card in
// the queue by its UUID.
func IndexOfCard(slice []*Card, value string) int {
	for i, v := range slice {
		if v.Uuid == value {
			return i
		}
	}
	return -1
}
