package cards

type BaseCardInterface interface {
	// Resolve defines the behavior of a card when it is played or activated in the game.
	// Each card type will implement its own logic for this method.
	Resolve()
	AttribuRteCard(player any)
}

type BaseCard struct {
	BaseCardInterface
	// Uuid is a unique identifier for the card, used to distinguish it from other cards in the game.
	Uuid string `json:"uuid"`
	// Player is the player associated with the card, if any.
	Player any `json:"player"`
}

func (c *BaseCard) Resolve() {}

// AttributeCard associates a card with a player in Redis, allowing for tracking of which player has which card.
func (c *BaseCard) AttributeCard(player any) {
	c.Player = player
}
