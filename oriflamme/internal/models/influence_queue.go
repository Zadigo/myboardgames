package models

type InfluenceQueue struct {
	// The influence queue is a list of cards that have been played but have not yet resolved.
	Queue []Card
	// Each card is resolved one by one in an ascending order of their position in the queue.
	// This index tracks which card is currently being resolved.
	ResolutionIndex int
}

func (queue *InfluenceQueue) AddCardLeft(card Card) {

}

func (queue *InfluenceQueue) AddCardRight(card Card) {

}

// Remove a card at a specific position in the queue. The card is simply
// marked as removed and will be skipped during resolution.
func (queue *InfluenceQueue) RemoveCardAtPosition(index int) *Card {
	card := queue.Queue[index]
	card.IsRemoved = true
	return &card
}

func (queue *InfluenceQueue) StackCard(index int) *Card {
	return nil
}

func (queue *InfluenceQueue) NumberOfCards() int {
	return len(queue.Queue)
}

func (queue *InfluenceQueue) Resolve(card *Card) {
	if queue.NumberOfCards() == 0 {
		return
	}

	if card == nil {
		return
	}

	index := IndexOfCard(queue.Queue, card.Uuid)

	if index == -1 {
		return
	}

	if index < queue.NumberOfCards() {
		queue.ResolutionIndex = index
	}
}

func (queue *InfluenceQueue) GetCurrentCard() *Card {
	if queue.NumberOfCards() == 0 {
		return nil
	}

	if queue.ResolutionIndex < queue.NumberOfCards() {
		return &queue.Queue[queue.ResolutionIndex]
	}

	return nil
}

func IndexOfCard(slice []Card, value string) int {
	for i, v := range slice {
		if v.Uuid == value {
			return i
		}
	}
	return -1
}
