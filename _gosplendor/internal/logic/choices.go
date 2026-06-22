package logic

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
	RemoteCard CardInterface `json:"remoteCard,omitempty"`
	// The card that the Royal Decree wants to move. This is used when the player chooses to apply the Royal Decree's effect
	// to move a card from its current position in the queue to another position.
	RoyalDecreeTargetCard CardInterface `json:"royalDecreeTargetCard,omitempty"`
	// The index of the card that the Royal Decree wants to move
	RoyalDecreeFrom int `json:"royalDecreeFrom,omitempty"`
	// The target position in the queue to which the Royal Decree wants to move the card
	RoyalDecreeTo int `json:"royalDecreeTo,omitempty"`
}
