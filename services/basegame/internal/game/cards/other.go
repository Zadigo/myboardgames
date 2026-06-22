package cards

// SpecialCard represents a special card in the game, inheriting from BaseCard.
type SpecialCard struct {
	BaseCard
}

func (c *SpecialCard) Resolve() {}

// SpecialExtensionCard represents a special extension card in the game, inheriting from BaseCard.
type SpecialExtensionCard struct {
	BaseCard
}

func (c *SpecialExtensionCard) Resolve() {}
