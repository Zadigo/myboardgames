package models

import (
	"github.com/google/uuid"
)

type Card struct {
	// Uuid is a unique identifier for the card,
	// used to track it in the influence queue
	// and other game mechanics.
	Uuid string
	// Indicates the position of the card in the influence queue.
	Position int
	// Name is the name of the card
	Name string
	// Description is a brief text describing the card's effect or role in the game.
	Description string
	// Type indicates whether the card is a "Character" or an "Intrigue".
	Type string
	// Stack represents the stack of cards that have been played on top of this card.
	Stack []Card
	// Color represents the color of the card, which can be "Red", "Blue", "Green", "Yellow", or "Purple".
	Color string
	// Owner is a reference to the player who owns this card.
	Owner *Player
	// IsSelected indicates whether the card was selected by the player.
	IsSelected bool
	// IsRemoved indicates whether the card has been removed from the game.
	IsRemoved bool
	// IsDiscarded indicates whether the card has been discarded from the queue.
	IsDiscarded bool
	// A player can choose to stack tokens on his cards which
	// he can win when the card is revealed during the resolution phase.
	Tokens int
	// Indicates whether the card has been revealed during the resolution phase.
	IsRevealed bool
}

func createBaseCards() []Card {
	return []Card{
		{Uuid: uuid.NewString(), Position: 0, Name: "Archer", Description: "", Type: "Character", IsSelected: false, Tokens: 0, IsRevealed: false},
		{Uuid: uuid.NewString(), Position: 0, Name: "Soldier", Description: "", Type: "Character", IsSelected: false, Tokens: 0, IsRevealed: false},
		{Uuid: uuid.NewString(), Position: 0, Name: "Spy", Description: "", Type: "Character", IsSelected: false, Tokens: 0, IsRevealed: false},
		{Uuid: uuid.NewString(), Position: 0, Name: "Heir", Description: "", Type: "Character", IsSelected: false, Tokens: 0, IsRevealed: false},
		{Uuid: uuid.NewString(), Position: 0, Name: "Shapeshifter", Description: "", Type: "Character", IsSelected: false, Tokens: 0, IsRevealed: false},
		{Uuid: uuid.NewString(), Position: 0, Name: "Lord", Description: "", Type: "Character", IsSelected: false, Tokens: 0, IsRevealed: false},

		{Uuid: uuid.NewString(), Position: 0, Name: "Assassination", Description: "", Type: "Intrigue", IsSelected: false, Tokens: 0, IsRevealed: false},
		{Uuid: uuid.NewString(), Position: 0, Name: "Royal Decree", Description: "", Type: "Intrigue", IsSelected: false, Tokens: 0, IsRevealed: false},
		{Uuid: uuid.NewString(), Position: 0, Name: "Conspiracy", Description: "", Type: "Intrigue", IsSelected: false, Tokens: 0, IsRevealed: false},
		{Uuid: uuid.NewString(), Position: 0, Name: "Ambush", Description: "", Type: "Intrigue", IsSelected: false, Tokens: 0, IsRevealed: false},
	}
}

func createColorCards(owner *Player, color string) []Card {
	cards := createBaseCards()
	for i := range cards {
		cards[i].Color = color
		cards[i].Owner = owner
	}
	return cards
}

func CreateRedCards(owner *Player) []Card {
	return createColorCards(owner, "Red")
}

func CreateBlueCards(owner *Player) []Card {
	return createColorCards(owner, "Blue")
}

func CreateGreenCards(owner *Player) []Card {
	return createColorCards(owner, "Green")
}

func CreateYellowCards(owner *Player) []Card {
	return createColorCards(owner, "Yellow")
}

func CreatePurpleCards(owner *Player) []Card {
	return createColorCards(owner, "Purple")
}

func (card *Card) IsCharacter() bool {
	return card.Type == "Character"
}

func (card *Card) IsIntrigue() bool {
	return card.Type == "Intrigue"
}

// Reveal a card during the resolution phase.
// The card's effect will be applied if it is not
// removed or discarded.
func (card *Card) Reveal() {
	if card.IsRemoved || card.IsDiscarded {
		return
	} else {
		card.IsRevealed = true
	}
}

// Discard a card from the queue. The card is simply marked
// as discarded and will be skipped during resolution.
func (card *Card) Discard() {
	card.IsDiscarded = true
}

// Add tokens to a card. The number of tokens on a card can be
// increased by the player, and these tokens can be won when
// the card is revealed during the resolution phase.
func (card *Card) AddToken(k int) {
	card.Tokens += k
}

// Remove tokens from a card. If the number of tokens to remove exceeds
// the current number of tokens on the card, the card's
// tokens will be set to zero.
func (card *Card) RemoveToken(k int) {
	if card.Tokens > 0 {
		card.Tokens -= k

		if card.Tokens < 0 {
			card.Tokens = 0
		}
	}
}

// Check if two cards are owned by the same player. This can be used to determine
// if a card's special ability can be applied to another card in the queue.
func (card *Card) IsSameOwnerAs(b *Card) bool {
	return card.Owner.Username == b.Owner.Username
}

// Check if the card's name matches the specified name and if it is revealed and
// also if the queue has cards to resolve. This is a common check for the card's special abilities.
func (card *Card) preActionCheck(queue *InfluenceQueue, checkName string) bool {
	if card.Name != checkName || !card.IsRevealed {
		return false
	}

	if queue.NumberOfCards() == 0 {
		return false
	}

	return true
}

// Eliminate a card anywhere in the queue
func (card *Card) ApplyAssassination(queue *InfluenceQueue, index int) bool {
	if !card.preActionCheck(queue, "Assassination") {
		return false
	}

	if index < 0 || index >= queue.NumberOfCards() {
		return false
	}

	// Eliminate the card at the specified index
	wasCard := queue.RemoveCardAtPosition(index)

	if wasCard.Name == "Ambush" {
		card.ApplyAmbushAttackedl(queue, card)
		return true
	}

	card.Owner.IncreaseTokens(1)
	return true
}

// Elimate the first or last card in the queue
func (card *Card) ApplyArcher(queue *InfluenceQueue, firstCard bool) bool {
	if !card.preActionCheck(queue, "Archer") {
		return false
	}

	wasCard := &Card{}

	if firstCard {
		wasCard = queue.RemoveCardAtPosition(0)
	} else {
		wasCard = queue.RemoveCardAtPosition(queue.NumberOfCards() - 1)
	}

	card.Owner.IncreaseTokens(1)

	if wasCard.Name == "Ambush" {
		card.ApplyAmbushAttackedl(queue, card)
	}

	return true
}

// Eliminates a card adjacent to the soldier in the queue.
// The player can choose to eliminate either the card immediately before
// or immediately after the soldier.
func (card *Card) ApplySoldier(queue *InfluenceQueue) bool {
	if !card.preActionCheck(queue, "Soldier") {
		return false
	}

	// The soldier is the only card in the queue,
	// so there are no adjacent cards to eliminate.
	if queue.NumberOfCards() == 1 {
		return false
	}

	var wasCard *Card

	finalize := func() {
		if wasCard.Name == "Ambush" {
			card.ApplyAmbushAttackedl(queue, card)
		}
	}

	// The soldier is at the front of the queue,
	// so only the card immediately after it
	// can be eliminated.
	if queue.ResolutionIndex == 0 {
		wasCard = queue.RemoveCardAtPosition(1)
		finalize()
		return true
	}

	// The soldier is at the end of the queue,
	// so only the card immediately before it
	// can be eliminated.
	if queue.ResolutionIndex == queue.NumberOfCards()-1 {
		wasCard = queue.RemoveCardAtPosition(queue.ResolutionIndex - 1)
		finalize()
		return true
	}

	// The soldier is in the middle of the queue,
	// so the player can choose to eliminate either
	// the card immediately before or immediately after it.
	wasCard = queue.RemoveCardAtPosition(queue.ResolutionIndex - 1)
	card.Owner.IncreaseTokens(1)
	finalize()
	return true
}

// Steal a token from the player's card immediately before or
// after the spy in the queue and add it to the spy.
func (card *Card) ApplySpy(queue *InfluenceQueue, cardBefore bool) bool {
	if !card.preActionCheck(queue, "Spy") {
		return false
	}

	// The spy is the only card in the queue,
	// so there are no adjacent cards to steal tokens from.
	if queue.NumberOfCards() == 1 {
		return false
	}

	// The spy is at the front of the queue,
	// so only the card immediately after it
	// can be stolen from.
	if queue.ResolutionIndex == 0 {
		card.Owner.IncreaseTokens(1)

		nextCard := queue.Queue[1]
		// If the next card is owned by the same player as the spy,
		// the spy cannot steal from his own self.
		if card.IsSameOwnerAs(nextCard) {
			return false
		}
		nextCard.Owner.DecreaseTokens(1)
		return true
	}

	// The spy is at the end of the queue,
	// so only the card immediately before it
	// can be stolen from.
	if queue.ResolutionIndex == queue.NumberOfCards()-1 {
		card.Owner.IncreaseTokens(1)
		prevCard := queue.Queue[queue.ResolutionIndex-1]

		if card.IsSameOwnerAs(prevCard) {
			return false
		}
		prevCard.Owner.DecreaseTokens(1)
		return true
	}

	// The spy is in the middle of the queue,
	// so the player can choose to steal from either
	// the card immediately before or immediately after it.
	if cardBefore {
		card.Owner.IncreaseTokens(1)
		prevCard := queue.Queue[queue.ResolutionIndex-1]

		if card.IsSameOwnerAs(prevCard) {
			return false
		}

		prevCard.Owner.DecreaseTokens(1)
	} else {
		card.Owner.IncreaseTokens(1)
		nextCard := queue.Queue[queue.ResolutionIndex+1]

		if card.IsSameOwnerAs(nextCard) {
			return false
		}

		nextCard.Owner.DecreaseTokens(1)
	}
	return true
}

// Move a card from its current position in the queue to another position.
func (card *Card) ApplyRoyalDecree(queue *InfluenceQueue, otherCard ...*Card) bool {
	if !card.preActionCheck(queue, "Royal Decree") {
		return false
	}

	card.Discard()
	return false
}

// If the card has tokens on it, when the card is revealed, the
// player wins double the number of tokens on the card.
func (card *Card) ApplyConspiracy(queue *InfluenceQueue) bool {
	if !card.preActionCheck(queue, "Conspiracy") {
		return false
	}

	if card.Tokens == 0 {
		return false
	}

	card.Owner.IncreaseTokens(card.Tokens * 2)
	card.Discard()
	return false
}

// The shapeshifter can copy the effect of any adjacent character card in the queue
// when it is revealed.
func (card *Card) ApplyShapeshifter(queue *InfluenceQueue, otherIndex int) bool {
	if !card.preActionCheck(queue, "Shapeshifter") {
		return false
	}

	// The shapeshifter is the only card in the queue,
	// so there are no adjacent cards to copy.
	if queue.NumberOfCards() == 1 {
		return false
	}

	var cardToCopy *Card

	// The shapeshifter is at the front of the queue,
	// so only the card immediately after it
	// can be copied.
	if queue.ResolutionIndex == 0 {
		cardToCopy = queue.Queue[1]
	}

	// The shapeshifter is at the end of the queue,
	// so only the card immediately before it
	// can be copied.
	if queue.ResolutionIndex == queue.NumberOfCards()-1 {
		cardToCopy = queue.Queue[queue.ResolutionIndex-1]
	}

	// The shapeshifter is in the middle of the queue,
	// so the player can choose to copy either
	// the card immediately before or immediately after it.
	if queue.ResolutionIndex > 0 && queue.ResolutionIndex < queue.NumberOfCards()-1 {
		cardToCopy = queue.Queue[otherIndex]
	}

	if cardToCopy.IsCharacter() && cardToCopy.IsRevealed {
		switch cardToCopy.Name {
		case "Archer":
			return cardToCopy.ApplyArcher(queue, true)
		case "Spy":
			return cardToCopy.ApplySpy(queue, true)
		case "Shapeshifter":
			// Copying another shapeshifter brings no
			// additional effect.
			return false
		}
	}

	return false
}

// Gain one token and one token for each card that is adjacent
// to the lord in the queue (revelaled or not).
func (card *Card) ApplyLord(queue *InfluenceQueue) bool {
	if !card.preActionCheck(queue, "Lord") {
		return false
	}

	card.Owner.IncreaseTokens(1)

	if queue.ResolutionIndex == 0 {
		nextCard := queue.Queue[1]

		if nextCard.Color == card.Color {
			card.Owner.IncreaseTokens(1)
		}
	}

	if queue.ResolutionIndex == queue.NumberOfCards()-1 {
		prevCard := queue.Queue[queue.ResolutionIndex-1]

		if prevCard.Color == card.Color {
			card.Owner.IncreaseTokens(1)
		}
	}

	if queue.ResolutionIndex > 0 && queue.ResolutionIndex < queue.NumberOfCards()-1 {
		nextCard := queue.Queue[queue.ResolutionIndex+1]
		prevCard := queue.Queue[queue.ResolutionIndex-1]

		if nextCard.Color == card.Color {
			card.Owner.IncreaseTokens(1)
		}

		if prevCard.Color == card.Color {
			card.Owner.IncreaseTokens(1)
		}
	}

	return false
}

// If there is exactly one heir in the queue, the player wins 2 tokens.
func (card *Card) ApplyHeir(queue *InfluenceQueue) bool {
	if !card.preActionCheck(queue, "Heir") {
		return false
	}

	heirs := 0

	for _, card := range queue.Queue {
		if card.Name == "Heir" {
			heirs++
		}
	}

	if heirs == 1 {
		card.Owner.IncreaseTokens(2)
		return true
	}

	return false
}

func (card *Card) ApplyAmbushReveal(queue *InfluenceQueue) bool {
	if !card.preActionCheck(queue, "Ambush") {
		return false
	}
	card.Owner.IncreaseTokens(1)
	card.Discard()
	return true
}

func (card *Card) ApplyAmbushAttackedl(queue *InfluenceQueue, attackingCard *Card) bool {
	if !card.preActionCheck(queue, "Ambush") {
		return false
	}

	card.Owner.IncreaseTokens(4)
	attackingCard.Discard()
	return true
}
