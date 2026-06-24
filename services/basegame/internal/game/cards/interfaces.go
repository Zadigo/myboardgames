package cards

type CardInterface interface {
	GetUuid() string
	// Resolution logic for the card's effect, if any
	Resolve(currentScore int, options CardResolveOptions) int
	// Set the player UUID to indicate ownership of the card
	SetPlayer(playerUuid string)
	// Get the player UUID of the player who owns the card
	GetPlayer() string
	// Get the value of the card
	GetValue() int
	// Check if the card is a special card (e.g., freeze, flip3, second_chance)
	IsSpecial() bool
	// Get the operator of the card (e.g., "+", "-", "*", "/", "freeze", "flip3", "second_chance")
	GetOperator() string
	// Check card operation type (e.g., "draw", "flip3", "freeze", "second_chance")
	IsCardOperation(operator string) bool
}

// CardFactory is an interface for creating cards.
// It defines a method to create a slice of CardInterface.
type CardFactory interface {
	CreateCards() []CardInterface
}
