package game

type BaseCardInterface interface {
	Resolve()
}

type BaseCard struct {
	BaseCardInterface
}

type NormalCard struct {
	BaseCard
}

func (c *NormalCard) Resolve() {}

type SpecialCard struct {
	BaseCard
}

func (c *SpecialCard) Resolve() {}

type SpecialExtensionCard struct {
	BaseCard
}

func (c *SpecialExtensionCard) Resolve() {}

func CreateCards()
