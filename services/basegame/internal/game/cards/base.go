package cards

import "github.com/Zadigo/basegame/internal/models"

type BaseCard struct {
	models.BaseCardInterface
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
