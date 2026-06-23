package cards

type CardInterface interface {
	GetUuid() string
	// Resolution logic for the card's effect, if any
	Resolve(currentScore int, options ResolveOptions) int
	// Set the player UUID to indicate ownership of the card
	SetPlayer(playerUuid string)
	// Get the player UUID of the player who owns the card
	GetPlayer() string
}

// CardFactory is an interface for creating cards.
// It defines a method to create a slice of CardInterface.
type CardFactory interface {
	CreateCards() []CardInterface
}
