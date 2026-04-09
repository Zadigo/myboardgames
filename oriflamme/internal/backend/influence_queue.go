package backend

import (
	"errors"
	"slices"
)

type InfluenceQueue struct {
	// The influence queue is a list of cards
	// that have been played but have not yet resolved.
	Queue []*Card
	// Each card is resolved one by one in an ascending order of their position in the queue.
	// This index tracks which card is currently being resolved. The index is initialized to
	// -1, meaning that no card has been resolved yet.
	ResolutionIndex int
}

func (queue *InfluenceQueue) updateCardIndexes() {
	for i := range queue.Queue {
		queue.Queue[i].PositionInQueue = i
	}
}

// Add a card to the left of the queue. This means that the card will be
// resolved before all other cards in the queue.
func (queue *InfluenceQueue) AddCardLeft(card *Card) {
	card.InQueue = true
	queue.Queue = slices.Insert(queue.Queue, 0, card)
	queue.updateCardIndexes()
}

// Add a card to the right of the queue. This means that the card will be
// resolved after all other cards in the queue.
func (queue *InfluenceQueue) AddCardRight(card *Card) {
	card.InQueue = true
	queue.Queue = append(queue.Queue, card)
	queue.updateCardIndexes()
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

// Resolve the next card in the queue. This function should be called to move to the next card
// being resolved. It returns an error if there are no cards in the queue or
// if the resolution index is out of bounds.
func (queue *InfluenceQueue) Resolve() error {
	if queue.NumberOfCards() == 0 {
		return errors.New("No cards in the queue")
	}

	if queue.ResolutionIndex >= queue.NumberOfCards() {
		return errors.New("Resolution index out of range")
	}

	queue.ResolutionIndex++
	return nil
}

// Get the current card being resolved. If the resolution index
// is out of bounds, return nil.
func (queue *InfluenceQueue) GetCurrentCard() (*Card, error) {
	if queue.NumberOfCards() == 0 {
		return nil, errors.New("No cards in the queue")
	}

	if queue.ResolutionIndex == -1 {
		return nil, errors.New("No card is being resolved yet")
	}

	if queue.ResolutionIndex < queue.NumberOfCards() {
		return queue.Queue[queue.ResolutionIndex], nil
	}

	return nil, errors.New("Resolution index out of range")
}

// Create a new influence queue with the given cards.
// The resolution index is initialized to -1.
func NewInfluenceQueue() *InfluenceQueue {
	return &InfluenceQueue{
		Queue:           []*Card{},
		ResolutionIndex: -1,
	}
}
