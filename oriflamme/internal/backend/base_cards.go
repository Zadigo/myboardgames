package backend

import (
	"errors"
	"slices"
)

// Base card colors
var colors = []string{"Red", "Blue", "Green", "Yellow", "Purple"}

// Eliminate a card anywhere in the queue
type AssassinationCard struct {
	BaseCard
}

func (c *AssassinationCard) GetBaseCard() *BaseCard {
	return &c.BaseCard
}

func (c *AssassinationCard) Reveal(queue *InfluenceQueue, choices PlayerChoices) (state bool, err error) {
	if choices.AtIndex < 0 || choices.AtIndex >= queue.NumberOfCards() {
		return false, errors.New("Invalid index for assassination")
	}

	// Eliminate the card at the specified index
	wasCard := queue.RemoveCardAtPosition(choices.AtIndex)

	if wasCard.GetBaseCard().Name == "Ambush" {
		// card.ApplyAmbushAttacked(queue, card)
		return wasCard.GetBaseCard().RevealSideEffect(c, choices)
	}

	c.Owner.IncreaseTokens(1)
	return true, nil
}

// Elimate the first or last card in the queue
type ArcherCard struct {
	BaseCard
}

func (c *ArcherCard) GetBaseCard() *BaseCard {
	return &c.BaseCard
}

func (c *ArcherCard) Reveal(queue *InfluenceQueue, choices PlayerChoices) (state bool, err error) {
	state, err = c.CanReveal(queue)
	if err != nil {
		return state, err
	}

	wasCard := CardInterface(nil)

	if choices.FirstCard {
		wasCard = queue.RemoveCardAtPosition(0)
	} else {
		wasCard = queue.RemoveCardAtPosition(queue.NumberOfCards() - 1)
	}

	c.Owner.IncreaseTokens(1)

	if wasCard.GetBaseCard().Name == "Ambush" {
		// card.ApplyAmbushAttacked(queue, card)
		return wasCard.GetBaseCard().RevealSideEffect(c, choices)
	}

	return true, nil
}

// Eliminates a card adjacent to the soldier in the queue.
// The player can choose to eliminate either the card immediately before
// or immediately after the soldier.
type SoldierCard struct {
	BaseCard
}

func (c *SoldierCard) GetBaseCard() *BaseCard {
	return &c.BaseCard
}

func (c *SoldierCard) Reveal(queue *InfluenceQueue, choices PlayerChoices) (state bool, err error) {
	// The soldier is the only card in the queue,
	// so there are no adjacent cards to eliminate.
	if queue.NumberOfCards() == 1 {
		return false, errors.New("No adjacent cards to eliminate")
	}

	var wasCard CardInterface

	finalize := func() {
		if choices.IsRemote {
			choices.RemoteCard.GetBaseCard().Owner.IncreaseTokens(1)
		} else {
			c.Owner.IncreaseTokens(1)

			if wasCard.GetBaseCard().Name == "Ambush" {
				// card.ApplyAmbushAttacked(queue, wasCard)
				wasCard.GetBaseCard().RevealSideEffect(c, choices)
			}
		}
	}

	if choices.IsRemote {
		if choices.ShapeShifterCardBefore {
			wasCard = queue.RemoveCardAtPosition(choices.TemporaryResolutionIndex - 1)
		} else {
			wasCard = queue.RemoveCardAtPosition(choices.TemporaryResolutionIndex + 1)
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
	if choices.CardBefore {
		wasCard = queue.RemoveCardAtPosition(queue.ResolutionIndex - 1)
	} else {
		wasCard = queue.RemoveCardAtPosition(queue.ResolutionIndex + 1)
	}

	finalize()
	return true, nil
}

// Steal a token from the player's card immediately before or
// after the spy in the queue and add it to the spy.
type SpyCard struct {
	BaseCard
}

func (c SpyCard) GetBaseCard() *BaseCard {
	return &c.BaseCard
}

func (c SpyCard) Reveal(queue *InfluenceQueue, choices PlayerChoices) (state bool, err error) {
	// The spy is the only card in the queue,
	// so there are no adjacent cards to steal tokens from.
	if queue.NumberOfCards() == 1 {
		return false, errors.New("No adjacent cards to steal tokens from")
	}

	if choices.IsRemote {
		var cardToStealFrom CardInterface

		if choices.ShapeShifterCardBefore {
			cardToStealFrom = queue.Queue[choices.TemporaryResolutionIndex-1]
		} else {
			cardToStealFrom = queue.Queue[choices.TemporaryResolutionIndex+1]
		}

		if sameOwner, err := c.IsSameOwnerAs(cardToStealFrom.GetBaseCard()); sameOwner || err != nil {
			return false, errors.Join(err, errors.New("Cannot steal from your own card"))
		}

		c.Owner.IncreaseTokens(1)
		cardToStealFrom.GetBaseCard().Owner.DecreaseTokens(1)
		return true, nil
	}

	// The spy is at the front of the queue,
	// so only the card immediately after it
	// can be stolen from.
	if queue.ResolutionIndex == 0 {
		c.Owner.IncreaseTokens(1)

		nextCard := queue.Queue[1]
		// If the next card is owned by the same player as the spy,
		// the spy cannot steal from his own self.
		if sameOwner, err := c.IsSameOwnerAs(nextCard.GetBaseCard()); sameOwner || err != nil {
			return false, errors.Join(err, errors.New("Cannot steal from your own card"))
		}
		nextCard.GetBaseCard().Owner.DecreaseTokens(1)
		return true, nil
	}

	// The spy is at the end of the queue,
	// so only the card immediately before it
	// can be stolen from.
	if queue.ResolutionIndex == queue.NumberOfCards()-1 {
		prevCard := queue.Queue[queue.ResolutionIndex-1]
		if sameOwner, err := c.IsSameOwnerAs(prevCard.GetBaseCard()); sameOwner || err != nil {
			return false, errors.Join(err, errors.New("Cannot steal from your own card"))
		}

		c.Owner.IncreaseTokens(1)
		prevCard.GetBaseCard().Owner.DecreaseTokens(1)
		return true, nil
	}

	// The spy is in the middle of the queue,
	// so the player can choose to steal from either
	// the card immediately before or immediately after it.
	if choices.CardBefore {
		prevCard := queue.Queue[queue.ResolutionIndex-1]

		if sameOwner, err := c.IsSameOwnerAs(prevCard.GetBaseCard()); sameOwner || err != nil {
			return false, errors.Join(err, errors.New("Cannot steal from your own card"))
		}

		c.Owner.IncreaseTokens(1)
		prevCard.GetBaseCard().Owner.DecreaseTokens(1)
	} else {
		nextCard := queue.Queue[queue.ResolutionIndex+1]

		if sameOwner, err := c.IsSameOwnerAs(nextCard.GetBaseCard()); sameOwner || err != nil {
			return false, errors.Join(err, errors.New("Cannot steal from your own card"))
		}

		c.Owner.IncreaseTokens(1)
		nextCard.GetBaseCard().Owner.DecreaseTokens(1)
	}
	return true, nil
}

// Move a card from its current position in the queue to another position.
type RoyalDecreeCard struct {
	BaseCard
}

func (c *RoyalDecreeCard) GetBaseCard() *BaseCard {
	return &c.BaseCard
}

func (c RoyalDecreeCard) Reveal(queue *InfluenceQueue, choices PlayerChoices) (state bool, err error) {
	return true, nil
}

// If the card has tokens on it, when the card is revealed, the
// player wins double the number of tokens on the card.
type ConspiracyCard struct {
	BaseCard
}

func (c *ConspiracyCard) GetBaseCard() *BaseCard {
	return &c.BaseCard
}

func (c *ConspiracyCard) Reveal(queue *InfluenceQueue, choices PlayerChoices) (state bool, err error) {
	c.Owner.IncreaseTokens(c.Tokens * 2)
	c.Discard()
	return true, nil
}

// The shapeshifter can copy the effect of any adjacent character card in the queue
// when it is revealed.
type ShapeShifterCard struct {
	BaseCard
}

func (c *ShapeShifterCard) GetBaseCard() *BaseCard {
	return &c.BaseCard
}

func (c *ShapeShifterCard) Reveal(queue *InfluenceQueue, choices PlayerChoices) (state bool, err error) {
	// The shapeshifter is the only card in the queue,
	// so there are no adjacent cards to copy.
	if queue.NumberOfCards() == 1 {
		return false, errors.New("No adjacent cards to copy")
	}

	var cardToCopy CardInterface

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
		if choices.CardBefore {
			cardToCopy = queue.Queue[queue.ResolutionIndex-1]
		} else {
			cardToCopy = queue.Queue[queue.ResolutionIndex+1]
		}
	}

	if cardToCopy.GetBaseCard().IsCharacter() && cardToCopy.GetBaseCard().IsRevealed {
		// Temporarily store the index of the card being copied
		choices.TemporaryResolutionIndex = -1

		for i, card := range queue.Queue {
			if card.GetBaseCard().Uuid == cardToCopy.GetBaseCard().Uuid {
				choices.TemporaryResolutionIndex = i
				break
			}
		}

		switch cardToCopy.GetBaseCard().Name {
		case "Archer":
			// return cardToCopy.ApplyArcher(queue, PlayerChoices{
			// 	FirstCard:                choice.ShapeShifterFirstCard,
			// 	TemporaryResolutionIndex: choice.TemporaryResolutionIndex,
			// 	IsRemote:                 true,
			// 	RemoteCard:               card,
			// })
			choices.IsRemote = true
			choices.RemoteCard = c
			return cardToCopy.GetBaseCard().RevealSideEffect(c, choices)
		case "Spy":
			// return cardToCopy.ApplySpy(queue, PlayerChoices{
			// 	CardBefore:               choice.ShapeShifterCardBefore,
			// 	TemporaryResolutionIndex: choice.TemporaryResolutionIndex,
			// 	IsRemote:                 true,
			// 	RemoteCard:               card,
			// })
			choices.IsRemote = true
			choices.RemoteCard = c
			return cardToCopy.GetBaseCard().RevealSideEffect(c, choices)
		case "Shapeshifter":
			// Copying another shapeshifter brings no additional effect
			return false, errors.New("Cannot copy another shapeshifter")
		case "Soldier":
			// return cardToCopy.ApplySoldier(queue, PlayerChoices{
			// 	CardBefore:               choice.ShapeShifterCardBefore,
			// 	TemporaryResolutionIndex: choice.TemporaryResolutionIndex,
			// 	IsRemote:                 true,
			// 	RemoteCard:               card,
			// })
			choices.IsRemote = true
			choices.RemoteCard = c
			return cardToCopy.GetBaseCard().RevealSideEffect(c, choices)
		case "Heir":
			choices.IsRemote = true
			choices.RemoteCard = c
			return cardToCopy.GetBaseCard().RevealSideEffect(c, choices)
			// return cardToCopy.ApplyHeir(queue)
		case "Lord":
			choices.IsRemote = true
			choices.RemoteCard = c
			return cardToCopy.GetBaseCard().RevealSideEffect(c, choices)
			// return cardToCopy.ApplyLord(queue, PlayerChoices{})

		default:
			return false, errors.New("Invalid card to copy")
		}
	}

	return false, nil
}

// Gain one token and one token for each card of the player's family
// that are adjacent to the lord in the queue (revelaled or not).
type LordCard struct {
	BaseCard
}

func (c *LordCard) GetBaseCard() *BaseCard {
	return &c.BaseCard
}

func (c *LordCard) Reveal(queue *InfluenceQueue, choices PlayerChoices) (state bool, err error) {
	return true, nil
}

// If there is exactly one heir in the queue, the player wins 2 tokens.
type HeirCard struct {
	BaseCard
}

func (c *HeirCard) GetBaseCard() *BaseCard {
	return &c.BaseCard
}

func (c *HeirCard) Reveal(queue *InfluenceQueue, choices PlayerChoices) (state bool, err error) {
	heirs := 0

	for _, card := range queue.Queue {
		if card.GetBaseCard().Name == "Heir" {
			heirs++
		}
	}

	if heirs == 1 {
		c.Owner.IncreaseTokens(2)
		return true, nil
	}

	return false, nil
}

// When the Ambush card is revealed during the resolution phase, the owner of the
// Ambush card wins 1 token and the card is discarded. If the Ambush card is attacked by
// another card, the owner of the Ambush card wins 4 tokens and the Ambush card is discarded.
// The attacking card is also discarded.
type AmbushCard struct {
	BaseCard
}

func (c *AmbushCard) GetBaseCard() *BaseCard {
	return &c.BaseCard
}

func (c *AmbushCard) Reveal(queue *InfluenceQueue, choices PlayerChoices) (state bool, err error) {
	c.Owner.IncreaseTokens(1)
	c.Discard()
	return true, nil
}

func (c *AmbushCard) RevealSideEffect(initiator CardInterface, choices PlayerChoices) (state bool, err error) {
	if initiator.GetBaseCard().Name == "Ambush" {
		// The Ambush card cannot attack its own self, so the ApplyAmbushAttacked is necessarily
		// called from another card's effect when the Ambush card is attacked. "initiator" is therefore
		// the attacking card on which the Ambush card's effect is being applied.
		return false, errors.New("Card is not an attacking card")
	} else {
		c.GetBaseCard().Owner.IncreaseTokens(4)
		c.Discard()
		initiator.GetBaseCard().Discard()
		return true, nil
	}
}

// BaseCard represents the common properties of all cards in the game.
func CreateBaseCards() []CardInterface {
	// colors := []string{"Red", "Blue", "Green", "Yellow", "Purple"}
	cards := []CardInterface{}

	for _, color := range colors {
		cards = append(cards, &ArcherCard{BaseCard: NewBaseCard("Archer", "Character", nil, color)})
		cards = append(cards, &SoldierCard{BaseCard: NewBaseCard("Soldier", "Character", nil, color)})
		cards = append(cards, &SpyCard{BaseCard: NewBaseCard("Spy", "Character", nil, color)})
		cards = append(cards, &ShapeShifterCard{BaseCard: NewBaseCard("Shape Shifter", "Character", nil, color)})
		cards = append(cards, &LordCard{BaseCard: NewBaseCard("Lord", "Character", nil, color)})
		cards = append(cards, &HeirCard{BaseCard: NewBaseCard("Heir", "Character", nil, color)})
		cards = append(cards, &RoyalDecreeCard{BaseCard: NewBaseCard("Royal Decree", "Intrigue", nil, color)})
		cards = append(cards, &ConspiracyCard{BaseCard: NewBaseCard("Conspiracy", "Intrigue", nil, color)})
		cards = append(cards, &AmbushCard{BaseCard: NewBaseCard("Ambush", "Intrigue", nil, color)})
	}
	return cards
}

// AttributeBaseCard assigns the given owner to all the base cards of the
// specified color and returns a slice of CardInterface containing all the cards
// of that color.
func AttributeBaseCard(owner *WebsocketClient, color string) ([]CardInterface, error) {
	cards := CreateBaseCards()

	if !slices.Contains(colors, color) {
		return nil, errors.New("Invalid color")
	}

	for i := range cards {
		card := cards[i].GetBaseCard()
		if card.Color == color {
			card.Owner = owner
		}
	}

	return cards, nil
}
