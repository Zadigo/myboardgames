package logic

type CardQueue struct {
	// The card queue is a list of cards that are waiting
	// to be flipped by the player. The cards in the queue
	// are ordered by the time they were added to the queue,
	// with the most recently added card at the end of the list.
	Queue []*Card `json:"queue"`
	// Each card is resolved one by one in an ascending order
	// of their position in the queue.
	ResolutionIndex int `json:"resolutionIndex"`
}

func NewCardQueue() *CardQueue {
	return &CardQueue{
		Queue:           []*Card{},
		ResolutionIndex: -1,
	}
}

func (q *CardQueue) FlipCard(player *WebsocketClient) *Card {
	q.ResolutionIndex++

	card := q.Queue[q.ResolutionIndex]
	if card.IsSpecial {
		// Do something
	} else {
		player.SetFlipCard(q, card)
	}

	return card
}
