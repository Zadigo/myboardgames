package backend

import (
	"github.com/google/uuid"
)

type PlayerChoices struct {
	// Whether to apply an effect on the card immediately before
	// or immediately after the card being resolved.
	CardBefore bool `json:"cardBefore,omitempty"`
	// Whether to apply an effect on the first or last card in the queue.
	FirstCard bool `json:"firstCard,omitempty"`
	// The index of the card in the queue on which to apply an effect.
	// This is used for the Assassination card, which can eliminate any card in the queue.
	AtIndex int `json:"atIndex,omitempty"`
	// The choice to apply to the card that the shapeshifter is copying. This is used when
	// the shapeshifter copies the effect of a soldier, which can eliminate either
	// the card immediately before or immediately after the soldier.
	ShapeShifterCardBefore bool `json:"shapeShifterCardBefore,omitempty"`
	// The choice to apply to the card that the shapeshifter is copying. This is used when
	// the shapeshifter copies the effect of an archer, which can eliminate either
	// the first or last card in the queue.
	ShapeShifterFirstCard bool `json:"shapeShifterFirstCard,omitempty"`
	// The index of the card in the queue on which to apply an effect. This is used when
	// the shapeshifter copies the effect of the assassination card, which can eliminate any card in the queue.
	ShapeShifterAtIndex int `json:"shapeShifterAtIndex,omitempty"`
	// When the Shapesifter copies the effect of a card, we need to resolve the effect based
	// on the position of the card being copied in the queue, not the position of the shapeshifter itself.
	// This field is used to temporarily store the index of the card being copied during the resolution of
	// the shapeshifter's effect.
	TemporaryResolutionIndex int `json:"temporaryResolutionIndex,omitempty"`
	// Indicates whether the card is being temporarily controlled by the
	// Shapeshifter's effect.
	IsRemote bool `json:"isRemote,omitempty"`
	// The card that is remotely controlling the effect of
	// another card (e.g. the shapeshifter copying the effect of another card).
	RemoteCard *BaseCard `json:"remoteCard,omitempty"`
	// The card that the Royal Decree wants to move. This is used when the player chooses to apply the Royal Decree's effect
	// to move a card from its current position in the queue to another position.
	RoyalDecreeTargetCard *BaseCard `json:"royalDecreeTargetCard,omitempty"`
	// The index of the card that the Royal Decree wants to move
	RoyalDecreeFrom int `json:"royalDecreeFrom,omitempty"`
	// The target position in the queue to which the Royal Decree wants to move the card
	RoyalDecreeTo int `json:"royalDecreeTo,omitempty"`
}

type BaseCard struct {
	// Uuid is a unique identifier for the card,
	// used to track it in the influence queue
	// and other game mechanics.
	Uuid string `json:"uuid"`
	// Indicates the position of the card in the influence queue. Default value
	// is -1, which means that the card is not in the queue. When a card is added to the queue,
	// its position is updated to reflect its index.
	PositionInQueue int `json:"positionInQueue"`
	// Name is the name of the card
	Name string `json:"name"`
	// Type indicates whether the card is a "Character" or an "Intrigue".
	Type string `json:"type"`
	// Stack represents the stack of cards that have been played on top of this card.
	Stack []*BaseCard `json:"stack"`
	// Color represents the color of the card, which can be "Red", "Blue", "Green", "Yellow", or "Purple".
	Color string `json:"color"`
	// Owner is a reference to the player who owns this card.
	Owner *WebsocketClient `json:"owner"`
	// IsSelected indicates whether the card was selected by the player.
	IsSelected bool `json:"isSelected"`
	// IsRemoved indicates whether the card has been removed from the game.
	IsRemoved bool `json:"isRemoved"`
	// IsDiscarded indicates whether the card has been discarded from the queue.
	IsDiscarded bool `json:"isDiscarded"`
	// InQueue indicates whether the card is currently in the influence queue.
	// This is used to determine whether the card's effect can be applied during the
	// resolution phase and for Nuxt to track the card's state in the frontend.
	InQueue bool `json:"inQueue"`
	// A player can choose to stack tokens on his cards which
	// he can win when the card is revealed during the resolution phase.
	Tokens int `json:"tokens"`
	// Indicates whether the card has been revealed during the resolution phase.
	IsRevealed bool `json:"isRevealed"`
	// Image is the URL of the card's image, used for displaying the card in the frontend.
	Image string `json:"image"`
}

type CardInterface interface {
	GetBaseCard() *BaseCard
	Reveal(queue *InfluenceQueue) (state bool, err error)
	Resolve(queue *InfluenceQueue, choice PlayerChoices) (state bool, err error)
}

func NewBaseCard(name string, cardType string, owner *WebsocketClient, color string) *BaseCard {
	return &BaseCard{
		Uuid:            uuid.NewString(),
		Color:           color,
		Name:            name,
		Image:           "",
		PositionInQueue: -1,
		Stack:           []*BaseCard{},
		Type:            cardType,
		Tokens:          0,
		Owner:           owner,
		IsSelected:      false,
		IsRemoved:       false,
		IsDiscarded:     false,
		InQueue:         false,
		IsRevealed:      false,
	}
}

// func createBaseCards() []CardInterface {
// 	cards := []CardInterface{
// 		ArcherCard{BaseCard{}},
// 	}

// 	// cards := []Card{
// 	// 	{Uuid: uuid.NewString(), PositionInQueue: -1, Name: "Archer", Type: "Character", Tokens: 0},
// 	// 	{Uuid: uuid.NewString(), PositionInQueue: -1, Name: "Soldier", Type: "Character", Tokens: 0},
// 	// 	{Uuid: uuid.NewString(), PositionInQueue: -1, Name: "Spy", Type: "Character", Tokens: 0},
// 	// 	{Uuid: uuid.NewString(), PositionInQueue: -1, Name: "Heir", Type: "Character", Tokens: 0},
// 	// 	{Uuid: uuid.NewString(), PositionInQueue: -1, Name: "Shapeshifter", Type: "Character", Tokens: 0},
// 	// 	{Uuid: uuid.NewString(), PositionInQueue: -1, Name: "Lord", Type: "Character", Tokens: 0},

// 	// 	{Uuid: uuid.NewString(), PositionInQueue: -1, Name: "Assassination", Type: "Intrigue", Tokens: 0},
// 	// 	{Uuid: uuid.NewString(), PositionInQueue: -1, Name: "Royal Decree", Type: "Intrigue", Tokens: 0},
// 	// 	{Uuid: uuid.NewString(), PositionInQueue: -1, Name: "Conspiracy", Type: "Intrigue", Tokens: 0},
// 	// 	{Uuid: uuid.NewString(), PositionInQueue: -1, Name: "Ambush", Type: "Intrigue", Tokens: 0},
// 	// }

// 	return cards
// }
