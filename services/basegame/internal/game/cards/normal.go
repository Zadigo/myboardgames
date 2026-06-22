package cards

// NormalCard represents a standard card in the game, inheriting from BaseCard.
type NormalCard struct {
	BaseCard
}

func (c *NormalCard) Resolve() {}
