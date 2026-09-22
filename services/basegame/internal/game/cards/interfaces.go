package cards

import "github.com/Zadigo/basegame/internal/models"

// CardFactory is an interface for creating cards.
// It defines a method to create a slice of CardInterface.
type CardFactory interface {
	CreateCards() []models.CardInterface
}
