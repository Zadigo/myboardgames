package cards

type BaseCardInterface interface {
	// Resolve defines the behavior of a card when it is played or activated in the game.
	// Each card type will implement its own logic for this method.
	Resolve()
}

type BaseCard struct {
	BaseCardInterface
}

func (c *BaseCard) Resolve() {}
