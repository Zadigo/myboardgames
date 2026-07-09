package games

import (
	"github.com/Zadigo/basegame/internal/models"
)

const (
	STANDARD  = "standard"
	EXTENSION = "extension"
)

// CreateGame is a factory function that creates a new game instance based on the specified game type.
// It returns an instance of GameInterface, which can be either a StandardGame or an ExtensionGame,
// depending on the provided gameType. If the gameType is not recognized, it returns nil.
func CreateGame(options models.GameAppOptions) GameInterface {
	var game GameInterface

	switch options.GameType {
	case STANDARD:
		game = &StandardGame{}
	case EXTENSION:
		game = &ExtensionGame{}
	default:
		return nil
	}

	game.SetOptions(options)
	return game
}
