package models

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type PlayerChoices struct {
	// Whether to apply an effect on the card immediately before
	// or immediately after the card being resolved.
	CardBefore bool
	// Whether to apply an effect on the first or last card in the queue.
	FirstCard bool
	// The index of the card in the queue on which to apply an effect.
	// This is used for the Assassination card, which can eliminate any card in the queue.
	AtIndex int
	// The choice to apply to the card that the shapeshifter is copying. This is used when
	// the shapeshifter copies the effect of a soldier, which can eliminate either
	// the card immediately before or immediately after the soldier.
	ShapeShifterCardBefore bool
	// The choice to apply to the card that the shapeshifter is copying. This is used when
	// the shapeshifter copies the effect of an archer, which can eliminate either
	// the first or last card in the queue.
	ShapeShifterFirstCard bool
	// The index of the card in the queue on which to apply an effect. This is used when
	// the shapeshifter copies the effect of the assassination card, which can eliminate any card in the queue.
	ShapeShifterAtIndex int
	// When the Shapesifter copies the effect of a card, we need to resolve the effect based
	// on the position of the card being copied in the queue, not the position of the shapeshifter itself.
	// This field is used to temporarily store the index of the card being copied during the resolution of
	// the shapeshifter's effect.
	TemporaryResolutionIndex int
	// Indicates whether the card is being temporarily controlled by the
	// Shapeshifter's effect.
	IsRemote bool
	// The card that is remotely controlling the effect of
	// another card (e.g. the shapeshifter copying the effect of another card).
	RemoteCard *Card
}

type Card struct {
	// Uuid is a unique identifier for the card,
	// used to track it in the influence queue
	// and other game mechanics.
	Uuid string
	// Indicates the position of the card in the influence queue. Default value
	// is -1, which means that the card is not in the queue. When a card is added to the queue,
	// its position is updated to reflect its index.
	PositionInQueue int
	// Name is the name of the card
	Name string
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
	// Image is the URL of the card's image, used for displaying the card in the frontend.
	Image string
}

func createBaseCards() []Card {
	return []Card{
		{Uuid: uuid.NewString(), PositionInQueue: -1, Name: "Archer", Type: "Character", Tokens: 0},
		{Uuid: uuid.NewString(), PositionInQueue: -1, Name: "Soldier", Type: "Character", Tokens: 0},
		{Uuid: uuid.NewString(), PositionInQueue: -1, Name: "Spy", Type: "Character", Tokens: 0},
		{Uuid: uuid.NewString(), PositionInQueue: -1, Name: "Heir", Type: "Character", Tokens: 0},
		{Uuid: uuid.NewString(), PositionInQueue: -1, Name: "Shapeshifter", Type: "Character", Tokens: 0},
		{Uuid: uuid.NewString(), PositionInQueue: -1, Name: "Lord", Type: "Character", Tokens: 0},

		{Uuid: uuid.NewString(), PositionInQueue: -1, Name: "Assassination", Type: "Intrigue", Tokens: 0},
		{Uuid: uuid.NewString(), PositionInQueue: -1, Name: "Royal Decree", Type: "Intrigue", Tokens: 0},
		{Uuid: uuid.NewString(), PositionInQueue: -1, Name: "Conspiracy", Type: "Intrigue", Tokens: 0},
		{Uuid: uuid.NewString(), PositionInQueue: -1, Name: "Ambush", Type: "Intrigue", Tokens: 0},
	}
}

func createColorCards(owner *Player, color string) []Card {
	cards := createBaseCards()

	for i := range cards {
		cards[i].Color = color
		cards[i].Owner = owner
		cards[i].Tokens = 0
		cards[i].Image = fmt.Sprintf("/oriflamme/%s/%s.png", color, cards[i].Name)
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
func (card *Card) Reveal(queue *InfluenceQueue) (state bool, err error) {
	switch card.Name {
	case "Ambush":
		return card.ApplyAmbushReveal(queue)
	default:
		if card.IsRemoved || card.IsDiscarded {
			return false, errors.New("Card is removed or discarded and cannot be revealed")
		} else {
			card.IsRevealed = true
			return true, nil
		}
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
func (card *Card) IsSameOwnerAs(b *Card) (bool, error) {
	if b.Owner == nil || card.Owner == nil {
		return false, errors.New("One of the cards does not have an owner")
	} else {
		return card.Owner.Username == b.Owner.Username, nil
	}
}

// Check if the card's name matches the specified name and if it is revealed and
// also if the queue has cards to resolve. This is a common check for the card's special abilities.
func (card *Card) preActionCheck(queue *InfluenceQueue, checkName string) (bool, error) {
	if card.Name != checkName || !card.IsRevealed {
		return false, errors.New("Card is not revealed or does not have the correct name")
	}

	if queue.NumberOfCards() == 0 {
		return false, errors.New("Queue has no cards to resolve")
	}

	return true, nil
}

// Eliminate a card anywhere in the queue
func (card *Card) ApplyAssassination(queue *InfluenceQueue, choice PlayerChoices) (state bool, err error) {
	if ok, _ := card.preActionCheck(queue, "Assassination"); !ok {
		return false, errors.New("Pre-action check failed")
	}

	if choice.AtIndex < 0 || choice.AtIndex >= queue.NumberOfCards() {
		return false, errors.New("Invalid index for assassination")
	}

	// Eliminate the card at the specified index
	wasCard := queue.RemoveCardAtPosition(choice.AtIndex)

	if wasCard.Name == "Ambush" {
		card.ApplyAmbushAttacked(queue, card)
		return true, nil
	}

	card.Owner.IncreaseTokens(1)
	return true, nil
}

// Elimate the first or last card in the queue
func (card *Card) ApplyArcher(queue *InfluenceQueue, choice PlayerChoices) (state bool, err error) {
	if ok, _ := card.preActionCheck(queue, "Archer"); !ok {
		return false, errors.New("Pre-action check failed")
	}

	wasCard := &Card{}

	if choice.FirstCard {
		wasCard = queue.RemoveCardAtPosition(0)
	} else {
		wasCard = queue.RemoveCardAtPosition(queue.NumberOfCards() - 1)
	}

	card.Owner.IncreaseTokens(1)

	if wasCard.Name == "Ambush" {
		card.ApplyAmbushAttacked(queue, card)
	}

	return true, nil
}

// Eliminates a card adjacent to the soldier in the queue.
// The player can choose to eliminate either the card immediately before
// or immediately after the soldier.
func (card *Card) ApplySoldier(queue *InfluenceQueue, choice PlayerChoices) (bool, error) {
	if ok, err := card.preActionCheck(queue, "Soldier"); !ok {
		return false, errors.Join(err, errors.New("Pre-action check failed"))
	}

	// The soldier is the only card in the queue,
	// so there are no adjacent cards to eliminate.
	if queue.NumberOfCards() == 1 {
		return false, errors.New("No adjacent cards to eliminate")
	}

	var wasCard *Card

	finalize := func() {
		if choice.IsRemote {
			choice.RemoteCard.Owner.IncreaseTokens(1)
		} else {
			card.Owner.IncreaseTokens(1)

			if wasCard.Name == "Ambush" {
				card.ApplyAmbushAttacked(queue, wasCard)
			}
		}
	}

	if choice.IsRemote {
		if choice.ShapeShifterCardBefore {
			wasCard = queue.RemoveCardAtPosition(choice.TemporaryResolutionIndex - 1)
		} else {
			wasCard = queue.RemoveCardAtPosition(choice.TemporaryResolutionIndex + 1)
		}
		finalize()
		return true, nil
	}

	// The soldier is at the front of the queue,
	// so only the card immediately after it
	// can be eliminated.
	if queue.ResolutionIndex == 0 {
		wasCard = queue.RemoveCardAtPosition(1)
		finalize()
		return true, nil
	}

	// The soldier is at the end of the queue,
	// so only the card immediately before it
	// can be eliminated.
	if queue.ResolutionIndex == queue.NumberOfCards()-1 {
		wasCard = queue.RemoveCardAtPosition(queue.ResolutionIndex - 1)
		finalize()
		return true, nil
	}

	// The soldier is in the middle of the queue,
	// so the player can choose to eliminate either
	// the card immediately before or immediately after it.
	if choice.CardBefore {
		wasCard = queue.RemoveCardAtPosition(queue.ResolutionIndex - 1)
	} else {
		wasCard = queue.RemoveCardAtPosition(queue.ResolutionIndex + 1)
	}

	finalize()
	return true, nil
}

// Steal a token from the player's card immediately before or
// after the spy in the queue and add it to the spy.
func (card *Card) ApplySpy(queue *InfluenceQueue, choice PlayerChoices) (state bool, err error) {
	if ok, err := card.preActionCheck(queue, "Spy"); !ok {
		return false, errors.Join(err, errors.New("Pre-action check failed"))
	}

	// The spy is the only card in the queue,
	// so there are no adjacent cards to steal tokens from.
	if queue.NumberOfCards() == 1 {
		return false, errors.New("No adjacent cards to steal tokens from")
	}

	if choice.IsRemote {
		var cardToStealFrom *Card

		if choice.ShapeShifterCardBefore {
			cardToStealFrom = queue.Queue[choice.TemporaryResolutionIndex-1]
		} else {
			cardToStealFrom = queue.Queue[choice.TemporaryResolutionIndex+1]
		}

		if sameOwner, err := card.IsSameOwnerAs(cardToStealFrom); sameOwner || err != nil {
			return false, errors.Join(err, errors.New("Cannot steal from your own card"))
		}

		card.Owner.IncreaseTokens(1)
		cardToStealFrom.Owner.DecreaseTokens(1)
		return true, nil
	}

	// The spy is at the front of the queue,
	// so only the card immediately after it
	// can be stolen from.
	if queue.ResolutionIndex == 0 {
		card.Owner.IncreaseTokens(1)

		nextCard := queue.Queue[1]
		// If the next card is owned by the same player as the spy,
		// the spy cannot steal from his own self.
		if sameOwner, err := card.IsSameOwnerAs(nextCard); sameOwner || err != nil {
			return false, errors.Join(err, errors.New("Cannot steal from your own card"))
		}
		nextCard.Owner.DecreaseTokens(1)
		return true, nil
	}

	// The spy is at the end of the queue,
	// so only the card immediately before it
	// can be stolen from.
	if queue.ResolutionIndex == queue.NumberOfCards()-1 {
		prevCard := queue.Queue[queue.ResolutionIndex-1]
		if sameOwner, err := card.IsSameOwnerAs(prevCard); sameOwner || err != nil {
			return false, errors.Join(err, errors.New("Cannot steal from your own card"))
		}

		card.Owner.IncreaseTokens(1)
		prevCard.Owner.DecreaseTokens(1)
		return true, nil
	}

	// The spy is in the middle of the queue,
	// so the player can choose to steal from either
	// the card immediately before or immediately after it.
	if choice.CardBefore {
		prevCard := queue.Queue[queue.ResolutionIndex-1]

		if sameOwner, err := card.IsSameOwnerAs(prevCard); sameOwner || err != nil {
			return false, errors.Join(err, errors.New("Cannot steal from your own card"))
		}

		card.Owner.IncreaseTokens(1)
		prevCard.Owner.DecreaseTokens(1)
	} else {
		nextCard := queue.Queue[queue.ResolutionIndex+1]

		if sameOwner, err := card.IsSameOwnerAs(nextCard); sameOwner || err != nil {
			return false, errors.Join(err, errors.New("Cannot steal from your own card"))
		}

		card.Owner.IncreaseTokens(1)
		nextCard.Owner.DecreaseTokens(1)
	}
	return true, nil
}

// Move a card from its current position in the queue to another position.
func (card *Card) ApplyRoyalDecree(queue *InfluenceQueue, otherCard ...*Card) bool {
	if ok, _ := card.preActionCheck(queue, "Royal Decree"); !ok {
		return false
	}

	card.Discard()
	return false
}

// If the card has tokens on it, when the card is revealed, the
// player wins double the number of tokens on the card.
func (card *Card) ApplyConspiracy(queue *InfluenceQueue) (state bool, err error) {
	if ok, err := card.preActionCheck(queue, "Conspiracy"); !ok {
		return false, errors.Join(err, errors.New("Pre-action check failed"))
	}

	card.Owner.IncreaseTokens(card.Tokens * 2)
	card.Discard()
	return true, nil
}

// The shapeshifter can copy the effect of any adjacent character card in the queue
// when it is revealed.
func (card *Card) ApplyShapeshifter(queue *InfluenceQueue, choice PlayerChoices) (state bool, err error) {
	if ok, _ := card.preActionCheck(queue, "Shapeshifter"); !ok {
		return false, errors.New("Pre-action check failed")
	}

	// The shapeshifter is the only card in the queue,
	// so there are no adjacent cards to copy.
	if queue.NumberOfCards() == 1 {
		return false, errors.New("No adjacent cards to copy")
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
		cardToCopy = queue.Queue[queue.NumberOfCards()-2]
	}

	// The shapeshifter is in the middle of the queue,
	// so the player can choose to copy either
	// the card immediately before or immediately after it.
	if queue.ResolutionIndex > 0 && queue.ResolutionIndex < queue.NumberOfCards()-1 {
		if choice.CardBefore {
			cardToCopy = queue.Queue[queue.ResolutionIndex-1]
		} else {
			cardToCopy = queue.Queue[queue.ResolutionIndex+1]
		}
	}

	if cardToCopy.IsCharacter() && cardToCopy.IsRevealed {
		// Temporarily store the index of the card being copied
		choice.TemporaryResolutionIndex = -1

		for i, card := range queue.Queue {
			if card.Uuid == cardToCopy.Uuid {
				choice.TemporaryResolutionIndex = i
				break
			}
		}

		switch cardToCopy.Name {
		case "Archer":
			return cardToCopy.ApplyArcher(queue, PlayerChoices{
				FirstCard:                choice.ShapeShifterFirstCard,
				TemporaryResolutionIndex: choice.TemporaryResolutionIndex,
				IsRemote:                 true,
				RemoteCard:               card,
			})
		case "Spy":
			return cardToCopy.ApplySpy(queue, PlayerChoices{
				CardBefore:               choice.ShapeShifterCardBefore,
				TemporaryResolutionIndex: choice.TemporaryResolutionIndex,
				IsRemote:                 true,
				RemoteCard:               card,
			})
		case "Shapeshifter":
			// Copying another shapeshifter brings no additional effect
			return false, errors.New("Cannot copy another shapeshifter")
		case "Soldier":
			return cardToCopy.ApplySoldier(queue, PlayerChoices{
				CardBefore:               choice.ShapeShifterCardBefore,
				TemporaryResolutionIndex: choice.TemporaryResolutionIndex,
				IsRemote:                 true,
				RemoteCard:               card,
			})
		case "Heir":
			return cardToCopy.ApplyHeir(queue)
		case "Lord":
			return cardToCopy.ApplyLord(queue, PlayerChoices{})

		default:
			return false, errors.New("Invalid card to copy")
		}
	}

	return false, nil
}

// Gain one token and one token for each card of the player's family
// that are adjacent to the lord in the queue (revelaled or not).
func (card *Card) ApplyLord(queue *InfluenceQueue, choice PlayerChoices) (state bool, err error) {
	if ok, _ := card.preActionCheck(queue, "Lord"); !ok {
		return false, errors.New("Pre-action check failed")
	}

	// TODO: Implement logic for shapeshifter

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
		prevCard := queue.Queue[queue.ResolutionIndex-1]
		nextCard := queue.Queue[queue.ResolutionIndex+1]

		if nextCard.Color == card.Color {
			card.Owner.IncreaseTokens(1)
		}

		if prevCard.Color == card.Color {
			card.Owner.IncreaseTokens(1)
		}
	}

	return false, nil
}

// If there is exactly one heir in the queue, the player wins 2 tokens.
func (card *Card) ApplyHeir(queue *InfluenceQueue) (state bool, err error) {
	if ok, _ := card.preActionCheck(queue, "Heir"); !ok {
		return false, errors.New("Pre-action check failed")
	}

	heirs := 0

	for _, card := range queue.Queue {
		if card.Name == "Heir" {
			heirs++
		}
	}

	if heirs == 1 {
		card.Owner.IncreaseTokens(2)
		return true, nil
	}

	return false, nil
}

// When the Ambush card is revealed during the resolution phase, the owner of the
// Ambush card wins 1 token and the card is discarded. If the Ambush card is attacked by
// another card, the owner of the Ambush card wins 4 tokens and the Ambush card is discarded.
// The attacking card is also discarded.
func (card *Card) ApplyAmbushReveal(queue *InfluenceQueue) (state bool, err error) {
	// if ok, err := card.preActionCheck(queue, "Ambush"); !ok {
	// 	return false, errors.Join(err, errors.New("Pre-action check failed"))
	// }

	card.Owner.IncreaseTokens(1)
	card.Discard()
	return true, nil
}

// When the Ambush card is attacked by another card, the owner of the Ambush card wins 4 tokens and
// the Ambush card is discarded. The attacking card is also discarded.
func (card *Card) ApplyAmbushAttacked(queue *InfluenceQueue, ambushCard *Card) (state bool, err error) {
	if card.Name == "Ambush" {
		// The Ambush card cannot attack its own self, so the ApplyAmbushAttacked is necessarily
		// called from another card's effect when the Ambush card is attacked. "card" is therefore
		// the attacking card on which the Ambush card's effect is being applied.
		return false, errors.New("Card is not an attacking card")
	} else {
		ambushCard.Owner.IncreaseTokens(4)
		ambushCard.Discard()
		card.Discard()
		return true, nil
	}
}
