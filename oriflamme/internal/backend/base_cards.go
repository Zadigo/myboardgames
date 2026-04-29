package backend

// Elimate the first or last card in the queue
type ArcherCard struct {
	*BaseCard
}

func (c ArcherCard) GetBaseCard() *BaseCard {
	return c.BaseCard
}

func (c ArcherCard) Reveal(queue *InfluenceQueue) (state bool, err error) {
	return true, nil
}

func (c ArcherCard) Resolve(queue *InfluenceQueue, choice PlayerChoices) (state bool, err error) {
	return true, nil
}

// Eliminates a card adjacent to the soldier in the queue.
// The player can choose to eliminate either the card immediately before
// or immediately after the soldier.
type SoldierCard struct {
	*BaseCard
}

func (c SoldierCard) GetBaseCard() *BaseCard {
	return c.BaseCard
}

func (c SoldierCard) Reveal(queue *InfluenceQueue) (state bool, err error) {
	return true, nil
}

func (c SoldierCard) Resolve(queue *InfluenceQueue, choice PlayerChoices) (state bool, err error) {
	return true, nil
}

// Steal a token from the player's card immediately before or
// after the spy in the queue and add it to the spy.
type SpyCard struct {
	*BaseCard
}

func (c SpyCard) GetBaseCard() *BaseCard {
	return c.BaseCard
}

func (c SpyCard) Reveal(queue *InfluenceQueue) (state bool, err error) {
	return true, nil
}

func (c SpyCard) Resolve(queue *InfluenceQueue, choice PlayerChoices) (state bool, err error) {
	return true, nil
}

// Move a card from its current position in the queue to another position.
type RoyalDecreeCard struct {
	*BaseCard
}

func (c RoyalDecreeCard) GetBaseCard() *BaseCard {
	return c.BaseCard
}

func (c RoyalDecreeCard) Reveal(queue *InfluenceQueue) (state bool, err error) {
	return true, nil
}

func (c RoyalDecreeCard) Resolve(queue *InfluenceQueue, choice PlayerChoices) (state bool, err error) {
	return true, nil
}

// If the card has tokens on it, when the card is revealed, the
// player wins double the number of tokens on the card.
type ConspiracyCard struct {
	*BaseCard
}

func (c ConspiracyCard) GetBaseCard() *BaseCard {
	return c.BaseCard
}

func (c ConspiracyCard) Reveal(queue *InfluenceQueue) (state bool, err error) {
	return true, nil
}

func (c ConspiracyCard) Resolve(queue *InfluenceQueue, choice PlayerChoices) (state bool, err error) {
	return true, nil
}

// The shapeshifter can copy the effect of any adjacent character card in the queue
// when it is revealed.
type ShapeShifterCard struct {
	*BaseCard
}

func (c ShapeShifterCard) GetBaseCard() *BaseCard {
	return c.BaseCard
}

func (c ShapeShifterCard) Reveal(queue *InfluenceQueue) (state bool, err error) {
	return true, nil
}

func (c ShapeShifterCard) Resolve(queue *InfluenceQueue, choice PlayerChoices) (state bool, err error) {
	return true, nil
}

// Gain one token and one token for each card of the player's family
// that are adjacent to the lord in the queue (revelaled or not).
type LordCard struct {
	*BaseCard
}

func (c LordCard) GetBaseCard() *BaseCard {
	return c.BaseCard
}

func (c LordCard) Reveal(queue *InfluenceQueue) (state bool, err error) {
	return true, nil
}

func (c LordCard) Resolve(queue *InfluenceQueue, choice PlayerChoices) (state bool, err error) {
	return true, nil
}

// If there is exactly one heir in the queue, the player wins 2 tokens.
type HeirCard struct {
	*BaseCard
}

func (c HeirCard) GetBaseCard() *BaseCard {
	return c.BaseCard
}

func (c HeirCard) Reveal(queue *InfluenceQueue) (state bool, err error) {
	return true, nil
}

func (c HeirCard) Resolve(queue *InfluenceQueue, choice PlayerChoices) (state bool, err error) {
	return true, nil
}

// When the Ambush card is revealed during the resolution phase, the owner of the
// Ambush card wins 1 token and the card is discarded. If the Ambush card is attacked by
// another card, the owner of the Ambush card wins 4 tokens and the Ambush card is discarded.
// The attacking card is also discarded.
type AmbushCard struct {
	*BaseCard
}

func (c AmbushCard) GetBaseCard() *BaseCard {
	return c.BaseCard
}

func (c AmbushCard) Reveal(queue *InfluenceQueue) (state bool, err error) {
	return true, nil
}

func (c AmbushCard) Resolve(queue *InfluenceQueue, choice PlayerChoices) (state bool, err error) {
	return true, nil
}

// BaseCard represents the common properties of all cards in the game.
func CreateBaseCards() []CardInterface {
	colors := []string{"Red", "Blue", "Green", "Yellow", "Purple"}
	cards := []CardInterface{}

	for _, color := range colors {
		cards = append(cards, ArcherCard{BaseCard: NewBaseCard("Archer", "Character", nil, color)})
		cards = append(cards, SoldierCard{BaseCard: NewBaseCard("Soldier", "Character", nil, color)})
		cards = append(cards, SpyCard{BaseCard: NewBaseCard("Spy", "Character", nil, color)})
		cards = append(cards, ShapeShifterCard{BaseCard: NewBaseCard("Shape Shifter", "Character", nil, color)})
		cards = append(cards, LordCard{BaseCard: NewBaseCard("Lord", "Character", nil, color)})
		cards = append(cards, HeirCard{BaseCard: NewBaseCard("Heir", "Character", nil, color)})
		cards = append(cards, RoyalDecreeCard{BaseCard: NewBaseCard("Royal Decree", "Intrigue", nil, color)})
		cards = append(cards, ConspiracyCard{BaseCard: NewBaseCard("Conspiracy", "Intrigue", nil, color)})
		cards = append(cards, AmbushCard{BaseCard: NewBaseCard("Ambush", "Intrigue", nil, color)})
	}
	return cards
}
